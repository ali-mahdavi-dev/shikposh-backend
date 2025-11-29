package admin

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/admin/adapter/phrases"
	"shikposh-backend/internal/admin/adapter/repository"
	"shikposh-backend/internal/admin/entrypoint"
	"shikposh-backend/internal/admin/entrypoint/handler"
	"shikposh-backend/internal/admin/query"
	mw "shikposh-backend/pkg/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config) error {
	// Register admin module error phrases
	phrases.RegisterAdminPhrases()

	// Initialize middleware
	middleware := mw.NewMiddleware(
		mw.MiddlewareConfig{JWTSecret: cfg.JWT.Secret},
		db,
	)

	// Initialize repository
	adminRepo := repository.NewAdminRepository(db)

	// Initialize query handlers
	adminQueryHandler := query.NewAdminQueryHandler(adminRepo)

	// Initialize HTTP handler
	adminHTTPHandler := handler.NewAdminHandler(adminQueryHandler, middleware)

	entrypoint.NewAdminRouter(router, entrypoint.AdminRouter{
		Admin: adminHTTPHandler,
	})

	return nil
}

