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
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	config "shikposh-backend/config"
	"shikposh-backend/internal/account"
	mwF "shikposh-backend/pkg/framework/api/middleware"
	"shikposh-backend/pkg/framework/infrastructure/databases"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/infrastructure/tracing"

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
	db     *gorm.DB
	server *fiber.App
	tracer *tracing.Tracer
}

func startServer(cfg *config.Config) error {
	// Initialize database
	var dsn string
	if cfg.Postgres.Password != "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	} else {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	}

	db, err := databases.New(databases.Config{
		DBType:       "postgres",
		DSN:          dsn,
		MaxOpenConns: cfg.Postgres.MaxOpenConns,
		MaxIdleConns: cfg.Postgres.MaxIdleConns,
		MaxLifetime:  int(cfg.Postgres.ConnMaxLifetime.Seconds()),
	})
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			if closeErr := sqlDB.Close(); closeErr != nil {
				logging.Warn("Failed to close database connection").WithError(closeErr).Log()
			}
		}
	}()

	// Initialize Jaeger tracing via OTLP
	serviceName := cfg.Jaeger.ServiceName
	if serviceName == "" {
		serviceName = cfg.Server.Name
	}

	environment := cfg.Jaeger.Environment
	if environment == "" {
		environment = cfg.Server.RunMode
	}

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
		// Continue without tracing if it fails
		tracer = nil
	}

	// Create Fiber app
	server := fiber.New(fiber.Config{
		AppName:      cfg.Server.Name,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	components := &serverComponents{
		db:     db,
		server: server,
		tracer: tracer,
	}

	// Setup routes and middleware
	if err := setupServer(components, cfg); err != nil {
		return fmt.Errorf("failed to setup server: %w", err)
	}

	// Setup graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Start server in a goroutine
	addr := fmt.Sprintf("%s:%s", cfg.Server.Domain, cfg.Server.InternalPort)
	serverErr := make(chan error, 1)

	go func() {
		// Log server ready after a short delay to ensure it's actually listening
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

	// Wait for interrupt signal or server error
	select {
	case err := <-serverErr:
		return err
	case <-quit:
	}

	// Graceful shutdown
	return gracefulShutdown(shutdownCtx, components)
}

func setupServer(components *serverComponents, cfg *config.Config) error {
	// Middleware
	middlewareF := mwF.NewMiddleware(mwF.MiddlewareConfig{JWTSecret: cfg.JWT.Secret}, components.db)

	// Register tracing middleware first (if enabled)
	if components.tracer != nil && cfg.Jaeger.Enabled {
		components.server.Use(middlewareF.TracingMiddleware())
	}

	middlewareF.Register(components.server)

	// Health check endpoint
	components.server.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": cfg.Server.Name,
		})
	})

	// Readiness check endpoint
	components.server.Get("/ready", func(c fiber.Ctx) error {
		// Check database connection
		sqlDB, err := components.db.DB()
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

	// Metrics (Prometheus)
	components.server.Get("/metrics", func(c fiber.Ctx) error {
		metricsHandler := promhttp.Handler()
		adapter := fasthttpadaptor.NewFastHTTPHandler(metricsHandler)

		if reqCtx, ok := c.Locals("requestCtx").(*fasthttp.RequestCtx); ok && reqCtx != nil {
			adapter(reqCtx)
			return nil
		}

		return c.Status(503).SendString("Metrics unavailable")
	})

	// Swagger Documentation
	registerSwagger(components.server)

	// Bootstrap application routes
	if err := account.Bootstrap(components.server, components.db, cfg); err != nil {
		return fmt.Errorf("failed to bootstrap account module: %w", err)
	}

	return nil
}

func gracefulShutdown(ctx context.Context, components *serverComponents) error {
	// Shutdown HTTP server
	if err := components.server.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	// Shutdown Jaeger tracer
	if components.tracer != nil {
		if err := components.tracer.Shutdown(ctx); err != nil {
			logging.Warn("Failed to shutdown Jaeger tracer").WithError(err).Log()
		}
	}

	return nil
}

func registerSwagger(app *fiber.App) {
	// Serve swagger.json file directly
	app.Get("/swagger.json", func(c fiber.Ctx) error {
		return c.SendFile("docs/swagger.json")
	})

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger.json",
	}))
}
