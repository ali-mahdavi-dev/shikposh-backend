package seller

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/seller/adapter/phrases"
	"shikposh-backend/internal/seller/entrypoint"
	"shikposh-backend/internal/seller/entrypoint/handler"
	"shikposh-backend/internal/seller/query"

	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config) error {
	// Register seller module error phrases
	phrases.RegisterSellerPhrases()

	// Create unit of work
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)

	// Initialize query handlers
	sellerQueryHandler := query.NewSellerQueryHandler(uow)

	// Initialize handler
	sellerHTTPHandler := handler.NewSellerHandler(sellerQueryHandler)

	entrypoint.NewSellerRouter(router, entrypoint.SellerManagementRouter{
		Seller: sellerHTTPHandler,
	})

	logging.Info("Seller module bootstrapped successfully").Log()

	return nil
}
