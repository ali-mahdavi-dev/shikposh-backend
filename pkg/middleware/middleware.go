package middleware

import (
	"shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/framework/adapter"
	frameworkmiddleware "github.com/ali-mahdavi-dev/framework/api/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type MiddlewareConfig struct {
	JWTSecret string
}

type Middleware struct {
	Cfg MiddlewareConfig
	Uow unitofwork.PGUnitOfWork
}

func NewMiddleware(cfg MiddlewareConfig, db *gorm.DB) *Middleware {
	// Create uow for middleware
	eventCh := make(chan adapter.EventWithWaitGroup, 1)
	uow := unitofwork.New(db, eventCh)

	return &Middleware{
		Cfg: cfg,
		Uow: uow,
	}
}

func (m *Middleware) Register(app *fiber.App) {
	// Request ID middleware should be registered first
	// so it's available for all subsequent middleware and handlers
	app.Use(frameworkmiddleware.RequestIDMiddleware())
	app.Use(frameworkmiddleware.DefaultStructuredLogger())
	app.Use(m.AuthMiddleware())
}
