package entrypoint

import (
	"shikposh-backend/internal/seller/entrypoint/handler"

	"github.com/gofiber/fiber/v3"
)

type SellerManagementRouter struct {
	Seller *handler.SellerHandler
}

func NewSellerRouter(router fiber.Router, controller SellerManagementRouter) {
	controller.Seller.RegisterRoutes(router)
}
