package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/swagger/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	config "shikposh-backend/config"
	"shikposh-backend/internal/account"
	"shikposh-backend/internal/admin"
	"shikposh-backend/internal/order"
	"shikposh-backend/internal/product"
	"shikposh-backend/internal/seller"
	mw "shikposh-backend/pkg/middleware"

	frameworkmiddleware "github.com/ali-mahdavi-dev/shikposh-framework/api/middleware"
	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	redisx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/redisx"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/tracing"
	servicehost "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/service_host"

	"gorm.io/gorm"
)

// ============================================================================
// Command Entry Point
// ============================================================================

func runHTTPServerCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "start http server",
		RunE: func(_ *cobra.Command, _ []string) error {
			initializeConfigs()
			log.Println("starting http server")

			return setupAndStartServices(&cfg)
		},
	}
}

// ============================================================================
// Service Setup
// ============================================================================

func setupAndStartServices(cfg *config.Config) error {
	// Initialize infrastructure services
	database, err := NewDatabaseService(cfg)
	if err != nil {
		return fmt.Errorf("database: %w", err)
	}

	redis, err := NewRedisService(cfg)
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	elasticsearch, err := NewElasticsearchService(cfg)
	if err != nil {
		logging.Warn("Elasticsearch initialization failed, continuing without it").
			WithError(err).
			Log()
		elasticsearch = nil
	} else {
		// Ensure products index exists
		ctx := context.Background()
		if err := EnsureProductsIndex(ctx, elasticsearch.Connection()); err != nil {
			logging.Warn("Failed to ensure products index exists").
				WithError(err).
				Log()
		}
	}

	tracing := NewTracingService(cfg)

	// Create HTTP service
	var elasticsearchConn elasticsearchx.Connection
	if elasticsearch != nil {
		elasticsearchConn = elasticsearch.Connection()
	}
	httpService, err := NewHTTPService(HTTPServiceConfig{
		Config:        cfg,
		Database:      database.DB(),
		Redis:         redis.Connection(),
		Elasticsearch: elasticsearchConn,
		Tracer:        tracing.Tracer(),
	})
	if err != nil {
		return fmt.Errorf("HTTP service: %w", err)
	}

	// Create service host and add services
	host := servicehost.NewServiceHost(30 * time.Second)
	host.AddService(httpService)
	host.AddService(tracing)
	host.AddService(database)

	// Start all services
	return host.Start()
}

// ============================================================================
// HTTP Service
// ============================================================================

type HTTPServiceConfig struct {
	Config        *config.Config
	Database      *gorm.DB
	Redis         redisx.Connection
	Elasticsearch elasticsearchx.Connection
	Tracer        *tracing.Tracer
}

type HTTPService struct {
	app        *fiber.App
	middleware *mw.Middleware
	cfg        *config.Config
	addr       string
}

func NewHTTPService(config HTTPServiceConfig) (*HTTPService, error) {
	app := fiber.New(fiber.Config{
		AppName:      config.Config.Server.Name,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	middleware := mw.NewMiddleware(
		mw.MiddlewareConfig{
			JWTSecret: config.Config.JWT.Secret,
			DB:        config.Database,
		},
	)

	addr := fmt.Sprintf("%s:%s", config.Config.Server.Domain, config.Config.Server.InternalPort)

	service := &HTTPService{
		app:        app,
		middleware: middleware,
		cfg:        config.Config,
		addr:       addr,
	}

	// Setup middleware
	if err := service.setupMiddleware(config); err != nil {
		return nil, fmt.Errorf("middleware setup: %w", err)
	}

	// Setup routes
	if err := service.setupRoutes(config); err != nil {
		return nil, fmt.Errorf("routes setup: %w", err)
	}

	return service, nil
}

// Service interface implementation
func (s *HTTPService) Name() string {
	return "http-server"
}

func (s *HTTPService) Start() error {
	// Log server ready after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		logging.Info("HTTP server ready").
			WithString("address", s.addr).
			WithString("swagger", fmt.Sprintf("http://%s/swagger/index.html", s.addr)).
			Log()
	}()

	return s.app.Listen(s.addr)
}

func (s *HTTPService) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func (s *HTTPService) setupMiddleware(config HTTPServiceConfig) error {
	// CORS
	allowOrigins := parseCORSOrigins(s.cfg.Cors.AllowOrigins)
	allowCredentials := s.cfg.Cors.AllowOrigins != "*"

	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowCredentials: allowCredentials,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	}))

	// Tracing middleware
	if config.Tracer != nil && s.cfg.Jaeger.Enabled {
		s.app.Use(frameworkmiddleware.TracingMiddleware())
	}

	// Application middleware
	s.middleware.Register(s.app)

	return nil
}

func (s *HTTPService) setupRoutes(config HTTPServiceConfig) error {
	// System routes
	s.setupSystemRoutes(config.Database)

	// Application routes
	if err := s.bootstrapModules(config); err != nil {
		return err
	}

	return nil
}

func (s *HTTPService) setupSystemRoutes(db *gorm.DB) {
	// Health check
	s.app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": s.cfg.Server.Name,
		})
	})

	// Readiness check
	s.app.Get("/ready", func(c fiber.Ctx) error {
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

	// Metrics
	s.app.Get("/metrics", func(c fiber.Ctx) error {
		metricsHandler := promhttp.Handler()
		adapter := fasthttpadaptor.NewFastHTTPHandler(metricsHandler)

		reqCtx := c.RequestCtx()
		if reqCtx != nil {
			adapter(reqCtx)
			return nil
		}

		return c.Status(503).SendString("Metrics unavailable")
	})

	// Swagger
	s.app.Get("/swagger.json", func(c fiber.Ctx) error {
		return c.SendFile("docs/swagger.json")
	})

	s.app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger.json",
	}))
}

func (s *HTTPService) bootstrapModules(config HTTPServiceConfig) error {
	modules := []struct {
		name string
		boot func() error
	}{
		{
			name: "account",
			boot: func() error {
				return account.Bootstrap(s.app, config.Database, s.cfg, config.Redis)
			},
		},
		{
			name: "order",
			boot: func() error {
				return order.Bootstrap(s.app, config.Database, s.cfg)
			},
		},
		{
			name: "product",
			boot: func() error {
				return product.Bootstrap(s.app, config.Database, s.cfg, config.Elasticsearch, s.middleware)
			},
		},
		{
			name: "seller",
			boot: func() error {
				return seller.Bootstrap(s.app, config.Database, s.cfg)
			},
		},
		{
			name: "admin",
			boot: func() error {
				return admin.Bootstrap(s.app, config.Database, s.cfg, s.middleware)
			},
		},
	}

	for _, module := range modules {
		if err := module.boot(); err != nil {
			return fmt.Errorf("failed to bootstrap %s module: %w", module.name, err)
		}
	}

	return nil
}
