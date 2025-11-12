package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	config "shikposh-backend/config"
	"shikposh-backend/internal/account"
	"shikposh-backend/internal/products"
	mwF "shikposh-backend/pkg/framework/api/middleware"
	"shikposh-backend/pkg/framework/infrastructure/databases"
	elasticsearchx "shikposh-backend/pkg/framework/infrastructure/elasticsearch"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/infrastructure/tracing"
	mw "shikposh-backend/pkg/middleware"

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
		logging.Warn("Failed to initialize Elasticsearch").
			WithError(err).
			Log()
		// Continue without Elasticsearch - it's optional for now
	}

	// Create Fiber app
	server := createFiberApp(cfg)

	components := &serverComponents{
		db:            db,
		server:        server,
		tracer:        tracer,
		elasticsearch: elasticsearch,
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

func initializeElasticsearch(cfg *config.Config) (elasticsearchx.Connection, error) {
	esCfg := elasticsearchx.Config{
		Host:     cfg.Elasticsearch.Host,
		Port:     cfg.Elasticsearch.Port,
		Username: cfg.Elasticsearch.Username,
		Password: cfg.Elasticsearch.Password,
	}

	conn, err := elasticsearchx.NewElasticsearchConnection(esCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize elasticsearch: %w", err)
	}

	return conn, nil
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
	middlewareF := mwF.NewMiddleware(
		mwF.MiddlewareConfig{},
		components.db,
	)
	middlewareM := mw.NewMiddleware(
		mw.MiddlewareConfig{JWTSecret: cfg.JWT.Secret},
		components.db,
	)

	// Register tracing middleware first (if enabled)
	if components.tracer != nil && cfg.Jaeger.Enabled {
		components.server.Use(middlewareF.TracingMiddleware())
	}

	middlewareF.Register(components.server)
	middlewareM.Register(components.server)
	return nil
}

func setupRoutes(components *serverComponents, cfg *config.Config) error {
	setupHealthRoutes(components.server, cfg)
	setupReadinessRoute(components.server, components.db)
	setupMetricsRoute(components.server)
	registerSwagger(components.server)

	// Bootstrap application routes
	if err := account.Bootstrap(components.server, components.db, cfg); err != nil {
		return fmt.Errorf("failed to bootstrap account module: %w", err)
	}

	if err := products.Bootstrap(components.server, components.db, cfg, components.elasticsearch); err != nil {
		return fmt.Errorf("failed to bootstrap products module: %w", err)
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
