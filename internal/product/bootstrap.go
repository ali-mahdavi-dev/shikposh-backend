package product

import (
	"context"

	"shikposh-backend/config"
	"shikposh-backend/internal/product/adapter/phrases"
	"shikposh-backend/internal/product/entrypoint"
	"shikposh-backend/internal/product/entrypoint/handler"
	"shikposh-backend/internal/product/query"
	"shikposh-backend/internal/product/service_layer/command_handler"
	"shikposh-backend/internal/product/service_layer/event_handler"
	"shikposh-backend/internal/product/service_layer/outbox"

	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	kafak "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/kafak"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	commandeventhandler "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler"
	commandmiddleware "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler/command_middleware"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	mw "shikposh-backend/pkg/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config, elasticsearch elasticsearchx.Connection, middleware *mw.Middleware) error {
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
	wishlistCommandHandler := command_handler.NewWishlistCommandHandler(uow)

	// Initialize event handlers
	productEventHandler := event_handler.NewProductEventHandler(uow, elasticsearch)

	// Initialize handler
	productHTTPHandler := handler.NewProductHandler(
		productQueryHandler,
		categoryQueryHandler,
		productHandler,
		bus,
		middleware,
	)
	wishlistHandler := handler.NewWishlistHandler(uow, bus)

	entrypoint.NewProductsRouter(router, entrypoint.ProductManagementRouter{
		Product:  productHTTPHandler,
		Wishlist: wishlistHandler,
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
		commandeventhandler.NewCommandHandler(wishlistCommandHandler.AddToWishlistHandler),
		commandeventhandler.NewCommandHandler(wishlistCommandHandler.RemoveFromWishlistHandler),
		commandeventhandler.NewCommandHandler(wishlistCommandHandler.ToggleWishlistHandler),
		commandeventhandler.NewCommandHandler(wishlistCommandHandler.SyncWishlistHandler),
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
