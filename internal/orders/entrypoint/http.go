package entrypoint

import (
	"shikposh-backend/internal/orders/entrypoint/handler"

	"github.com/gofiber/fiber/v3"
)

type OrderRouter struct {
	Order *handler.OrderHandler
}

func NewOrdersRouter(router fiber.Router, controller OrderRouter) {
	controller.Order.RegisterRoutes(router)
}

