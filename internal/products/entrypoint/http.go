package entrypoint

import (
	"shikposh-backend/internal/products/entrypoint/handler"

	"github.com/gofiber/fiber/v3"
)

type ProductManagementRouter struct {
	Product  *handler.ProductHandler
	Wishlist *handler.WishlistHandler
}

func NewProductsRouter(router fiber.Router, controller ProductManagementRouter) {
	controller.Product.RegisterRoutes(router)
	controller.Wishlist.RegisterRoutes(router)
}
