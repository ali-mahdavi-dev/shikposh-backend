package entryporint

import (
	"shikposh-backend/internal/products/entryporint/handler"

	"github.com/gofiber/fiber/v3"
)

type ProductManagementRouter struct {
	Product *handler.ProductHandler
}

func NewProductsRouter(router fiber.Router, controller ProductManagementRouter) {
	controller.Product.RegisterRoutes(router)
}
