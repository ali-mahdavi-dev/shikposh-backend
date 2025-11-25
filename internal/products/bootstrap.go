package products

import (
	"context"
	"fmt"

	"shikposh-backend/config"
	"shikposh-backend/internal/products/adapter/phrases"
	"shikposh-backend/internal/products/entrypoint"
	"shikposh-backend/internal/products/entrypoint/handler"
	"shikposh-backend/internal/products/query"
	"shikposh-backend/internal/products/service_layer/command_handler"
	"shikposh-backend/internal/products/service_layer/event_handler"
	"shikposh-backend/internal/products/service_layer/outbox"

	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	kafak "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/kafak"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	commandeventhandler "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler"
	commandmiddleware "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler/command_middleware"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config, elasticsearch elasticsearchx.Connection) error {
	// Register products module error phrases
	phrases.RegisterProductsPhrases()
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)
	bus := messagebus.NewMessageBus(uow, eventCh)

	// Initialize query handlers
	productQueryHandler := query.NewProductQueryHandler(uow, elasticsearch)
	categoryQueryHandler := query.NewCategoryQueryHandler(uow)

	// Initialize command handlers
	productHandler := command_handler.NewProductCommandHandler(uow)

	// Initialize event handlers
	productEventHandler := event_handler.NewProductEventHandler(uow)

	// Initialize handler
	productHTTPHandler := handler.NewProductHandler(
		productQueryHandler,
		categoryQueryHandler,
		productHandler,
		bus,
	)

	entrypoint.NewProductsRouter(router, entrypoint.ProductManagementRouter{
		Product: productHTTPHandler,
	})

	// register command middlewares
	bus.AddCommandMiddleware(
		commandmiddleware.Logging(),
	)

	// command handlers
	bus.AddCommandHandler(
		commandeventhandler.NewCommandHandler(productHandler.CreateProductHandler),
		commandeventhandler.NewCommandHandler(productHandler.UpdateProductHandler),
		commandeventhandler.NewCommandHandler(productHandler.DeleteProductHandler),
	)

	// event handlers
	bus.AddEventHandler(
		commandeventhandler.NewEventHandler(productEventHandler.ProductCreatedEvent),
		commandeventhandler.NewEventHandler(productEventHandler.ProductUpdatedEvent),
		commandeventhandler.NewEventHandler(productEventHandler.ProductDeletedEvent),
	)

	// Initialize outbox processor (reads from outbox and sends to Kafka)
	kafkaService := kafak.Service
	outboxProcessor := outbox.NewProcessor(uow, kafkaService)
	ctx := context.Background()
	outboxProcessor.Start(ctx)

	// Initialize Kafka consumer (consumes from Kafka and indexes in Elasticsearch)
	if elasticsearch != nil {
		// Ensure products index exists
		if err := ensureProductsIndex(ctx, elasticsearch); err != nil {
			logging.Warn("Failed to ensure products index exists").
				WithError(err).
				Log()
		}

		outboxConsumer := outbox.NewConsumer(uow, elasticsearch, kafkaService)
		if outboxConsumer != nil {
			go func() {
				if err := outboxConsumer.Start(ctx); err != nil {
					logging.Error("Failed to start outbox consumer").
						WithError(err).
						Log()
				}
			}()
		}
	} else {
		logging.Warn("Elasticsearch not available, outbox consumer will not start").Log()
	}

	logging.Info("Products module bootstrapped successfully").Log()

	return nil
}

// ensureProductsIndex creates the products index if it doesn't exist
func ensureProductsIndex(ctx context.Context, elasticsearch elasticsearchx.Connection) error {
	indexName := "products"

	// Check if index exists
	exists, err := elasticsearch.IndexExists(ctx, indexName)
	if err != nil {
		return fmt.Errorf("failed to check if index exists: %w", err)
	}

	if exists {
		logging.Debug("Products index already exists").WithString("index", indexName).Log()
		return nil
	}

	// Create index with mapping
	mapping := map[string]interface{}{
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

	if err := elasticsearch.CreateIndex(ctx, indexName, mapping); err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}

	logging.Info("Products index created successfully").WithString("index", indexName).Log()
	return nil
}
