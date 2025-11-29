package entrypoint

import (
	"shikposh-backend/internal/admin/entrypoint/handler"

	"github.com/gofiber/fiber/v3"
)

type AdminRouter struct {
	Admin *handler.AdminHandler
}

func NewAdminRouter(router fiber.Router, adminRouter AdminRouter) {
	adminRouter.Admin.RegisterRoutes(router)
}
