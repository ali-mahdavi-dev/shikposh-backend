package middleware

import (
	"shikposh-backend/pkg/framework/adapter"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type MiddlewareConfig struct {
	JWTSecret string
}

type Middleware struct {
	Cfg MiddlewareConfig
	Uow unit_of_work.PGUnitOfWork
}

func NewMiddleware(cfg MiddlewareConfig, db *gorm.DB) *Middleware {
	// Create uow for middleware
	eventCh := make(chan adapter.EventWithWaitGroup, 1)
	uow := unit_of_work.New(db, eventCh)

	return &Middleware{
		Cfg: cfg,
		Uow: uow,
	}
}

func (m *Middleware) Register(app *fiber.App) {
	app.Use(m.AuthMiddleware())
}
