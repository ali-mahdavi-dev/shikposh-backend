package order

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/order/adapter/payment"
	"shikposh-backend/internal/order/adapter/phrases"
	"shikposh-backend/internal/order/entrypoint"
	"shikposh-backend/internal/order/entrypoint/handler"
	"shikposh-backend/internal/order/query"
	"shikposh-backend/internal/order/service_layer/command_handler"

	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	commandeventhandler "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler"
	commandmiddleware "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler/command_middleware"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config) error {
	// Register orders module error phrases
	phrases.RegisterOrdersPhrases()

	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)
	bus := messagebus.NewMessageBus(uow, eventCh)

	// Initialize ZarinPal payment service
	zarinPalService := payment.NewZarinPalService(&cfg.ZarinPal)

	// Initialize query handlers
	orderQueryHandler := query.NewOrderQueryHandler(uow)

	// Initialize command handlers
	orderCommandHandler := command_handler.NewOrderCommandHandler(uow)

	// Initialize HTTP handler
	orderHTTPHandler := handler.NewOrderHandler(
		orderQueryHandler,
		zarinPalService,
		bus,
		orderCommandHandler,
	)

	entrypoint.NewOrdersRouter(router, entrypoint.OrderRouter{
		Order: orderHTTPHandler,
	})

	// register command middlewares
	bus.AddCommandMiddleware(
		commandmiddleware.Logging(),
	)

	// register command handlers (only commands without result)
	bus.AddCommandHandler(
		commandeventhandler.NewCommandHandler(orderCommandHandler.CancelOrderHandler),
		commandeventhandler.NewCommandHandler(orderCommandHandler.UpdateOrderStatusHandler),
	)

	return nil
}
