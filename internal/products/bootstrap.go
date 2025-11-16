package products

import (
	"context"

	"shikposh-backend/config"
	"shikposh-backend/internal/products/entrypoint"
	"shikposh-backend/internal/products/entrypoint/handler"
	"shikposh-backend/internal/products/query"
	"shikposh-backend/internal/products/service_layer/command_handler"
	"shikposh-backend/internal/products/service_layer/event_handler"
	"shikposh-backend/internal/products/service_layer/outbox"

	"github.com/ali-mahdavi-dev/framework/adapter"
	elasticsearchx "github.com/ali-mahdavi-dev/framework/infrastructure/elasticsearch"
	kafak "github.com/ali-mahdavi-dev/framework/infrastructure/kafak"
	"github.com/ali-mahdavi-dev/framework/infrastructure/logging"
	commandeventhandler "github.com/ali-mahdavi-dev/framework/service_layer/command_event_handler"
	commandmiddleware "github.com/ali-mahdavi-dev/framework/service_layer/command_event_handler/command_middleware"
	"github.com/ali-mahdavi-dev/framework/service_layer/messagebus"
	"shikposh-backend/internal/unit_of_work"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config, elasticsearch elasticsearchx.Connection) error {
	// Create event channel and unit of work for this module
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)
	bus := messagebus.NewMessageBus(uow, eventCh)

	// Initialize query handlers
	productQueryHandler := query.NewProductQueryHandler(uow, elasticsearch)
	categoryQueryHandler := query.NewCategoryQueryHandler(uow)
	reviewQueryHandler := query.NewReviewQueryHandler(uow)

	// Initialize command handlers
	reviewHandler := command_handler.NewReviewCommandHandler(uow)
	productHandler := command_handler.NewProductCommandHandler(uow)

	// Initialize event handlers
	productEventHandler := event_handler.NewProductEventHandler(uow)

	// Initialize handler
	productHTTPHandler := handler.NewProductHandler(
		productQueryHandler,
		categoryQueryHandler,
		reviewQueryHandler,
		reviewHandler,
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
		commandeventhandler.NewCommandHandler(reviewHandler.CreateReviewHandler),
		commandeventhandler.NewCommandHandler(reviewHandler.UpdateReviewHelpfulHandler),
		commandeventhandler.NewCommandHandler(productHandler.CreateProductHandler),
		commandeventhandler.NewCommandHandler(productHandler.UpdateProductHandler),
		commandeventhandler.NewCommandHandler(productHandler.DeleteProductHandler),
	)

	// event handlers
	bus.AddEventHandler(
		commandeventhandler.NewEventHandler(productEventHandler.ProductCreatedEvent),
	)

	// Initialize outbox processor (reads from outbox and sends to Kafka)
	kafkaService := kafak.Service
	outboxProcessor := outbox.NewProcessor(uow, kafkaService)
	ctx := context.Background()
	outboxProcessor.Start(ctx)

	// Initialize Kafka consumer (consumes from Kafka and indexes in Elasticsearch)
	if elasticsearch != nil {
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
