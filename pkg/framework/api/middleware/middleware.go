package middleware

import (
	"errors"

	"shikposh-backend/pkg/framework/adapter"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var errFailGetTokenFromDB = errors.New("fail to get token from DB")
var errTokenDoesNotExist = errors.New("token does not exist")

// MiddlewareConfig holds configuration for middleware
type MiddlewareConfig struct {
	JWTSecret string
}

type Middleware struct {
	Cfg MiddlewareConfig
	Uow unit_of_work.PGUnitOfWork
}

func NewMiddleware(cfg MiddlewareConfig, db *gorm.DB) *Middleware {
	// Create uow for middleware
	eventCh := make(chan adapter.EventWithWaitGroup, 2)
	uow := unit_of_work.New(db, eventCh)

	return &Middleware{
		Cfg: cfg,
		Uow: uow,
	}
}

func (m *Middleware) Register(app *fiber.App) {
	// Request ID middleware should be registered first
	// so it's available for all subsequent middleware and handlers
	app.Use(m.RequestIDMiddleware())
	app.Use(m.DefaultStructuredLogger())
}
