package entrypoint

import (
	"shikposh-backend/internal/account/entrypoint/handler"

	"github.com/gofiber/fiber/v3"
)

type UserManagementRouter struct {
	User *handler.UserController
}

func NewAccountRouter(router fiber.Router, controller UserManagementRouter) {
	controller.User.RegisterRoutes(router)
}
