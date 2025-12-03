package command

import (
	"context"
	"fmt"
	config "shikposh-backend/config"
	"strconv"
	"strings"
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/databases"
	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	redisx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/redisx"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/tracing"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// initializeElasticsearch is a helper function for other commands
// It reuses NewElasticsearchService to avoid code duplication
func initializeElasticsearch(cfg *config.Config) (elasticsearchx.Connection, error) {
	service, err := NewElasticsearchService(cfg)
	if err != nil {
		return nil, err
	}
	return service.Connection(), nil
}

// closeDatabase is a helper function for other commands
func closeDatabase(db *gorm.DB) {
	if sqlDB, err := db.DB(); err == nil {
		if closeErr := sqlDB.Close(); closeErr != nil {
			logging.Warn("Failed to close database connection").WithError(closeErr).Log()
		}
	}
}

// initializeDatabase is a helper function for other commands
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

// ============================================================================
// Database Service
// ============================================================================

type DatabaseService struct {
	db *gorm.DB
}

func NewDatabaseService(cfg *config.Config) (*DatabaseService, error) {
	dsn := buildDSN(cfg)

	db, err := databases.New(databases.Config{
		DBType:       "postgres",
		DSN:          dsn,
		MaxOpenConns: cfg.Postgres.MaxOpenConns,
		MaxIdleConns: cfg.Postgres.MaxIdleConns,
		MaxLifetime:  int(cfg.Postgres.ConnMaxLifetime.Seconds()),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize: %w", err)
	}

	return &DatabaseService{db: db}, nil
}

func (s *DatabaseService) DB() *gorm.DB {
	return s.db
}

// Service interface implementation
func (s *DatabaseService) Name() string {
	return "database"
}

func (s *DatabaseService) Start() error {
	// Database is already started when initialized
	return nil
}

func (s *DatabaseService) Shutdown(ctx context.Context) error {
	if sqlDB, err := s.db.DB(); err == nil {
		if closeErr := sqlDB.Close(); closeErr != nil {
			logging.Warn("Failed to close database connection").WithError(closeErr).Log()
			return closeErr
		}
	}
	return nil
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

// ============================================================================
// Redis Service
// ============================================================================

type RedisService struct {
	conn redisx.Connection
}

func NewRedisService(cfg *config.Config) (*RedisService, error) {
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
		return nil, fmt.Errorf("failed to initialize: %w", err)
	}

	return &RedisService{conn: conn}, nil
}

func (s *RedisService) Connection() redisx.Connection {
	return s.conn
}

// ============================================================================
// Elasticsearch Service
// ============================================================================

type ElasticsearchService struct {
	conn elasticsearchx.Connection
}

func NewElasticsearchService(cfg *config.Config) (*ElasticsearchService, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	type result struct {
		conn elasticsearchx.Connection
		err  error
	}
	resultChan := make(chan result, 1)

	go func() {
		conn, err := elasticsearchx.NewElasticsearchConnection(esCfg)
		resultChan <- result{conn: conn, err: err}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("initialization timeout: %w", ctx.Err())
	case r := <-resultChan:
		if r.err != nil {
			return nil, fmt.Errorf("failed to initialize: %w", r.err)
		}
		logging.Info("Elasticsearch initialized successfully").
			WithString("host", cfg.Elasticsearch.Host).
			WithString("port", cfg.Elasticsearch.Port).
			Log()
		return &ElasticsearchService{conn: r.conn}, nil
	}
}

func (s *ElasticsearchService) Connection() elasticsearchx.Connection {
	return s.conn
}

// ============================================================================
// Tracing Service
// ============================================================================

type TracingService struct {
	tracer *tracing.Tracer
	cfg    *config.Config
}

func NewTracingService(cfg *config.Config) *TracingService {
	if !cfg.Jaeger.Enabled {
		return &TracingService{cfg: cfg}
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
		return &TracingService{cfg: cfg}
	}

	return &TracingService{tracer: tracer, cfg: cfg}
}

func (s *TracingService) Tracer() *tracing.Tracer {
	return s.tracer
}

// Service interface implementation
func (s *TracingService) Name() string {
	return "tracing"
}

func (s *TracingService) Start() error {
	// Tracing is already started when initialized
	return nil
}

func (s *TracingService) Shutdown(ctx context.Context) error {
	if s.tracer == nil {
		return nil
	}
	return s.tracer.Shutdown(ctx)
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

// ============================================================================
// Helpers
// ============================================================================

func parseCORSOrigins(origins string) []string {
	if origins == "*" {
		return []string{"*"}
	}

	split := strings.Split(origins, ",")
	result := make([]string, 0, len(split))
	for _, origin := range split {
		trimmed := strings.TrimSpace(origin)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

// ============================================================================
// Products Index Mapping
// ============================================================================

const productsIndexName = "products"

// GetProductsIndexMapping returns the Elasticsearch mapping for products index
func GetProductsIndexMapping() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type": "keyword",
			},
			"seller_id": map[string]interface{}{
				"type": "integer",
			},
			"brand": map[string]interface{}{
				"type": "text",
			},
			"title": map[string]interface{}{
				"type": "text",
			},
			"slug": map[string]interface{}{
				"type": "keyword",
			},
			"description": map[string]interface{}{
				"type": "text",
			},
			"thumbnail": map[string]interface{}{
				"type": "keyword",
			},
			"discount": map[string]interface{}{
				"type": "integer",
			},
			"stock": map[string]interface{}{
				"type": "integer",
			},
			"price": map[string]interface{}{
				"type": "integer",
			},
			"rating": map[string]interface{}{
				"type": "float",
			},
			"is_featured": map[string]interface{}{
				"type": "boolean",
			},
			"is_new": map[string]interface{}{
				"type": "boolean",
			},
			"created_at": map[string]interface{}{
				"type": "date",
			},
			"category_id": map[string]interface{}{
				"type": "integer",
			},
			"categories": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"name": map[string]interface{}{
						"type": "text",
					},
					"slug": map[string]interface{}{
						"type": "keyword",
					},
				},
			},
			"colors": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"name": map[string]interface{}{
						"type": "text",
					},
					"hex": map[string]interface{}{
						"type": "keyword",
					},
				},
			},
			"sizes": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"name": map[string]interface{}{
						"type": "keyword",
					},
				},
			},
			"tags": map[string]interface{}{
				"type": "keyword",
			},
			"features": map[string]interface{}{
				"type": "text",
			},
			"specs": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"key": map[string]interface{}{
						"type": "keyword",
					},
					"value": map[string]interface{}{
						"type": "text",
					},
				},
			},
			"images": map[string]interface{}{
				"type": "object",
			},
			"variant": map[string]interface{}{
				"type": "object",
			},
		},
	}
}

// EnsureProductsIndex creates the products index if it doesn't exist
func EnsureProductsIndex(ctx context.Context, elasticsearch elasticsearchx.Connection) error {
	// Check if index exists
	exists, err := elasticsearch.IndexExists(ctx, productsIndexName)
	if err != nil {
		return fmt.Errorf("failed to check if index exists: %w", err)
	}

	if exists {
		logging.Debug("Products index already exists").WithString("index", productsIndexName).Log()
		return nil
	}

	// Create index with mapping
	mapping := GetProductsIndexMapping()
	if err := elasticsearch.CreateIndex(ctx, productsIndexName, mapping); err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}

	logging.Info("Products index created successfully").WithString("index", productsIndexName).Log()
	return nil
}
