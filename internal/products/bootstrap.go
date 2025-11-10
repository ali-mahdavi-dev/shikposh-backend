package products

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/products/entryporint"
	"shikposh-backend/internal/products/entryporint/handler"
	"shikposh-backend/internal/products/query"
	"shikposh-backend/internal/products/service_layer/command_handler"
	"shikposh-backend/internal/products/service_layer/event_handler"

	"shikposh-backend/pkg/framework/adapter"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	commandeventhandler "shikposh-backend/pkg/framework/service_layer/command_event_handler"
	"shikposh-backend/pkg/framework/service_layer/messagebus"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config) error {
	// Create event channel and unit of work for this module
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unit_of_work.New(db, eventCh)
	bus := messagebus.NewMessageBus(uow, eventCh)

	// Initialize query handlers
	productQueryHandler := query.NewProductQueryHandler(uow)
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

	entryporint.NewProductsRouter(router, entryporint.ProductManagementRouter{
		Product: productHTTPHandler,
	})

	// command handlers
	bus.AddHandler(
		commandeventhandler.NewCommandHandlerWithResult(reviewHandler.CreateReviewHandler),
		commandeventhandler.NewCommandHandlerWithResult(reviewHandler.UpdateReviewHelpfulHandler),
		commandeventhandler.NewCommandHandler(productHandler.CreateProductHandler),
		commandeventhandler.NewCommandHandler(productHandler.UpdateProductHandler),
		commandeventhandler.NewCommandHandler(productHandler.DeleteProductHandler),
	)

	// event handlers
	bus.AddHandlerEvent(
		commandeventhandler.NewEventHandler(productEventHandler.ProductCreatedEvent),
	)

	logging.Info("Products module bootstrapped successfully").Log()

	return nil
}
