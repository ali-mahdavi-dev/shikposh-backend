package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/swagger/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	config "shikposh-backend/config"
	"shikposh-backend/internal/account"
	"shikposh-backend/internal/products"
	"shikposh-backend/internal/seller"
	mw "shikposh-backend/pkg/middleware"

	frameworkmiddleware "github.com/ali-mahdavi-dev/shikposh-framework/api/middleware"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/databases"
	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	redisx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/redisx"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/tracing"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

func runHTTPServerCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "start http server",
		RunE: func(_ *cobra.Command, _ []string) error {
			initializeConfigs()
			log.Println("starting http server")
			return startServer(&cfg)
		},
	}
}

type serverComponents struct {
	db            *gorm.DB
	server        *fiber.App
	tracer        *tracing.Tracer
	elasticsearch elasticsearchx.Connection
	redis         redisx.Connection
}

func startServer(cfg *config.Config) error {
	// Initialize components
	db, err := initializeDatabase(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer closeDatabase(db)

	tracer := initializeTracing(cfg)

	elasticsearch, err := initializeElasticsearch(cfg)
	if err != nil {
		logging.Error("Failed to initialize Elasticsearch").
			WithError(err).
			WithString("host", cfg.Elasticsearch.Host).
			WithString("port", cfg.Elasticsearch.Port).
			Log()
		// Continue without Elasticsearch - it's optional for now
		elasticsearch = nil
	} else {
		logging.Info("Elasticsearch initialized successfully").
			WithString("host", cfg.Elasticsearch.Host).
			WithString("port", cfg.Elasticsearch.Port).
			Log()
	}

	redis, err := initializeRedis(cfg)
	if err != nil {
		logging.Warn("Failed to initialize Redis").
			WithError(err).
			Log()
		return fmt.Errorf("failed to initialize Redis: %w", err)
	}

	// Create Fiber app
	server := createFiberApp(cfg)

	components := &serverComponents{
		db:            db,
		server:        server,
		tracer:        tracer,
		elasticsearch: elasticsearch,
		redis:         redis,
	}

	// Setup routes and middleware
	if err := setupServer(components, cfg); err != nil {
		return fmt.Errorf("failed to setup server: %w", err)
	}

	// Start server and wait for shutdown
	return runServer(components, cfg)
}

func initializeDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := buildDSN(cfg)

	db, err := databases.New(databases.Config{
		DBType:       "postgres",
		DSN:          dsn,
		MaxOpenConns: cfg.Postgres.MaxOpenConns,
		MaxIdleConns: cfg.Postgres.MaxIdleConns,
		MaxLifetime:  int(cfg.Postgres.ConnMaxLifetime.Seconds()),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return db, nil
}

func buildDSN(cfg *config.Config) string {
	if cfg.Postgres.Password != "" {
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.DbName,
			cfg.Postgres.SSLMode,
		)
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)
}

func closeDatabase(db *gorm.DB) {
	if sqlDB, err := db.DB(); err == nil {
		if closeErr := sqlDB.Close(); closeErr != nil {
			logging.Warn("Failed to close database connection").WithError(closeErr).Log()
		}
	}
}

func initializeTracing(cfg *config.Config) *tracing.Tracer {
	if !cfg.Jaeger.Enabled {
		return nil
	}

	serviceName := getServiceName(cfg)
	environment := getEnvironment(cfg)

	tracer, err := tracing.New(tracing.Config{
		ServiceName:  serviceName,
		OTLPEndpoint: cfg.Jaeger.OTLPEndpoint,
		Environment:  environment,
		SamplingRate: cfg.Jaeger.SamplingRate,
		Enabled:      cfg.Jaeger.Enabled,
	})
	if err != nil {
		logging.Warn("Failed to initialize Jaeger tracing").
			WithError(err).
			Log()
		return nil
	}

	return tracer
}

func getServiceName(cfg *config.Config) string {
	if cfg.Jaeger.ServiceName != "" {
		return cfg.Jaeger.ServiceName
	}
	return cfg.Server.Name
}

func getEnvironment(cfg *config.Config) string {
	if cfg.Jaeger.Environment != "" {
		return cfg.Jaeger.Environment
	}
	return cfg.Server.RunMode
}

func initializeRedis(cfg *config.Config) (redisx.Connection, error) {
	ctx := context.Background()

	db, err := strconv.Atoi(cfg.Redis.Db)
	if err != nil {
		db = 0
	}

	options := &redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           db,
		DialTimeout:  cfg.Redis.DialTimeout,
		ReadTimeout:  cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout,
	}

	conn, err := redisx.NewRedisConnection(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	return conn, nil
}

func initializeElasticsearch(cfg *config.Config) (elasticsearchx.Connection, error) {
	logging.Info("Initializing Elasticsearch connection").
		WithString("host", cfg.Elasticsearch.Host).
		WithString("port", cfg.Elasticsearch.Port).
		Log()

	esCfg := elasticsearchx.Config{
		Host:     cfg.Elasticsearch.Host,
		Port:     cfg.Elasticsearch.Port,
		Username: cfg.Elasticsearch.Username,
		Password: cfg.Elasticsearch.Password,
	}

	// Use a timeout context to prevent blocking startup for too long
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	type result struct {
		conn elasticsearchx.Connection
		err  error
	}
	resultChan := make(chan result, 1)

	// Run initialization in a goroutine to make it cancellable
	go func() {
		conn, err := elasticsearchx.NewElasticsearchConnection(esCfg)
		resultChan <- result{conn: conn, err: err}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("elasticsearch initialization timeout after 3 seconds: %w", ctx.Err())
	case r := <-resultChan:
		if r.err != nil {
			return nil, fmt.Errorf("failed to initialize elasticsearch: %w", r.err)
		}
		return r.conn, nil
	}
}

func createFiberApp(cfg *config.Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:      cfg.Server.Name,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})
}

func setupServer(components *serverComponents, cfg *config.Config) error {
	if err := setupMiddleware(components, cfg); err != nil {
		return fmt.Errorf("failed to setup middleware: %w", err)
	}

	if err := setupRoutes(components, cfg); err != nil {
		return fmt.Errorf("failed to setup routes: %w", err)
	}

	return nil
}

func setupMiddleware(components *serverComponents, cfg *config.Config) error {
	// Register CORS middleware first
	// Parse AllowOrigins: if "*", use it directly; otherwise split by comma
	var allowOrigins []string
	if cfg.Cors.AllowOrigins == "*" {
		allowOrigins = []string{"*"}
	} else {
		// Split by comma to support multiple origins
		origins := strings.Split(cfg.Cors.AllowOrigins, ",")
		allowOrigins = make([]string, 0, len(origins))
		for _, origin := range origins {
			trimmed := strings.TrimSpace(origin)
			if trimmed != "" {
				allowOrigins = append(allowOrigins, trimmed)
			}
		}
	}

	// When using wildcard "*", AllowCredentials must be false
	// When using specific origins, AllowCredentials can be true
	allowCredentials := cfg.Cors.AllowOrigins != "*"

	components.server.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowCredentials: allowCredentials,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	}))

	middleware := mw.NewMiddleware(
		mw.MiddlewareConfig{JWTSecret: cfg.JWT.Secret},
		components.db,
	)

	// Register tracing middleware (if enabled)
	if components.tracer != nil && cfg.Jaeger.Enabled {
		components.server.Use(frameworkmiddleware.TracingMiddleware())
	}

	middleware.Register(components.server)
	return nil
}

func setupRoutes(components *serverComponents, cfg *config.Config) error {
	setupHealthRoutes(components.server, cfg)
	setupReadinessRoute(components.server, components.db)
	setupMetricsRoute(components.server)
	registerSwagger(components.server)

	// Bootstrap application routes
	if err := account.Bootstrap(components.server, components.db, cfg, components.redis); err != nil {
		return fmt.Errorf("failed to bootstrap account module: %w", err)
	}

	if err := products.Bootstrap(components.server, components.db, cfg, components.elasticsearch); err != nil {
		return fmt.Errorf("failed to bootstrap products module: %w", err)
	}

	if err := seller.Bootstrap(components.server, components.db, cfg); err != nil {
		return fmt.Errorf("failed to bootstrap seller module: %w", err)
	}

	return nil
}

func setupHealthRoutes(app *fiber.App, cfg *config.Config) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": cfg.Server.Name,
		})
	})
}

func setupReadinessRoute(app *fiber.App, db *gorm.DB) {
	app.Get("/ready", func(c fiber.Ctx) error {
		sqlDB, err := db.DB()
		if err != nil {
			return c.Status(503).JSON(fiber.Map{
				"status": "not ready",
				"error":  "database connection failed",
			})
		}

		if err := sqlDB.Ping(); err != nil {
			return c.Status(503).JSON(fiber.Map{
				"status": "not ready",
				"error":  "database ping failed",
			})
		}

		return c.JSON(fiber.Map{
			"status": "ready",
		})
	})
}

func setupMetricsRoute(app *fiber.App) {
	app.Get("/metrics", func(c fiber.Ctx) error {
		metricsHandler := promhttp.Handler()
		adapter := fasthttpadaptor.NewFastHTTPHandler(metricsHandler)

		reqCtx := c.RequestCtx()
		if reqCtx != nil {
			adapter(reqCtx)
			return nil
		}

		return c.Status(503).SendString("Metrics unavailable")
	})
}

func registerSwagger(app *fiber.App) {
	app.Get("/swagger.json", func(c fiber.Ctx) error {
		return c.SendFile("docs/swagger.json")
	})

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger.json",
	}))
}

func runServer(components *serverComponents, cfg *config.Config) error {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Domain, cfg.Server.InternalPort)
	serverErr := make(chan error, 1)

	startServerAsync(components.server, addr, serverErr)

	// Wait for interrupt signal or server error
	select {
	case err := <-serverErr:
		return err
	case <-quit:
		return gracefulShutdown(shutdownCtx, components)
	}
}

func startServerAsync(server *fiber.App, addr string, serverErr chan<- error) {
	go func() {
		// Log server ready after a short delay
		go func() {
			time.Sleep(100 * time.Millisecond)
			logging.Info("HTTP server ready").
				WithString("address", addr).
				WithString("swagger", fmt.Sprintf("http://%s/swagger/index.html", addr)).
				Log()
		}()

		if err := server.Listen(addr); err != nil {
			serverErr <- fmt.Errorf("server failed: %w", err)
		}
	}()
}

func gracefulShutdown(ctx context.Context, components *serverComponents) error {
	if err := components.server.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	if components.tracer != nil {
		if err := components.tracer.Shutdown(ctx); err != nil {
			logging.Warn("Failed to shutdown Jaeger tracer").WithError(err).Log()
		}
	}

	return nil
}
