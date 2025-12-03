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

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config) error {
	// Register orders module error phrases
	phrases.RegisterOrdersPhrases()

	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)

	// Initialize ZarinPal payment service
	zarinPalService := payment.NewZarinPalService(&cfg.ZarinPal)

	// Initialize query handlers
	orderQueryHandler := query.NewOrderQueryHandler(uow)

	// Initialize command handlers
	orderCommandHandler := command_handler.NewOrderCommandHandler(uow)

	// Initialize HTTP handler
	orderHTTPHandler := handler.NewOrderHandler(
		orderQueryHandler,
		orderCommandHandler,
		zarinPalService,
	)

	entrypoint.NewOrdersRouter(router, entrypoint.OrderRouter{
		Order: orderHTTPHandler,
	})

	return nil
}
