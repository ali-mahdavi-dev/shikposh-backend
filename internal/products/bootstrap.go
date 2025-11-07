package products

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/products/entryporint/handler"
	"shikposh-backend/internal/products/query"
	"shikposh-backend/internal/products/service_layer/command_handler"
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

	// Initialize handler
	productHandler := handler.NewProductHandler(
		productQueryHandler,
		categoryQueryHandler,
		reviewQueryHandler,
		reviewHandler,
		bus,
	)
	productHandler.RegisterRoutes(router)

	// Register command handlers
	bus.AddHandler(
		commandeventhandler.NewCommandHandlerWithResult(reviewHandler.CreateReviewHandler),
		commandeventhandler.NewCommandHandlerWithResult(reviewHandler.UpdateReviewHelpfulHandler),
	)

	logging.Info("Products module bootstrapped successfully").Log()

	return nil
}
