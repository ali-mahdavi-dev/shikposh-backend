package middleware

import (
	"net/http"

	"shikposh-backend/internal/account/domain/entity"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

// AdminMiddleware checks if user is admin or superuser
// Note: This middleware expects AuthMiddleware to run first and store user in context
func (m *Middleware) AdminMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get user from context (set by AuthMiddleware to avoid repeated DB queries)
		userInterface := c.Locals("user")
		if userInterface == nil {
			// Fallback: if user not in context, try to get from user_id
			userID := c.Locals("user_id")
			if userID == nil {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}

			// Only query DB if user not in context (shouldn't happen if AuthMiddleware ran first)
			ctx := c.Context()
			userIDUint := cast.ToUint64(userID)
			user, err := m.Uow.User(ctx).FindByID(ctx, userIDUint)
			if err != nil {
				return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "User not found"})
			}
			c.Locals("user", user)
			userInterface = user
		}

		// Type assert user
		user, ok := userInterface.(*entity.User)
		if !ok {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user context"})
		}

		// Check if user is admin or superuser
		if !user.IsAdmin && !user.IsSuperuser {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Admin access required"})
		}

		// Store user info in context for later use
		c.Locals("admin_user", user)
		c.Locals("is_superuser", user.IsSuperuser)

		return c.Next()
	}
}

// SuperuserMiddleware checks if user is superuser only
// Note: This middleware expects AuthMiddleware to run first and store user in context
func (m *Middleware) SuperuserMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get user from context (set by AuthMiddleware to avoid repeated DB queries)
		userInterface := c.Locals("user")
		if userInterface == nil {
			// Fallback: if user not in context, try to get from user_id
			userID := c.Locals("user_id")
			if userID == nil {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}

			// Only query DB if user not in context (shouldn't happen if AuthMiddleware ran first)
			ctx := c.Context()
			userIDUint := cast.ToUint64(userID)
			user, err := m.Uow.User(ctx).FindByID(ctx, userIDUint)
			if err != nil {
				return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "User not found"})
			}
			c.Locals("user", user)
			userInterface = user
		}

		// Type assert user
		user, ok := userInterface.(*entity.User)
		if !ok {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user context"})
		}

		// Check if user is superuser
		if !user.IsSuperuser {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Superuser access required"})
		}

		// Store user info in context
		c.Locals("admin_user", user)
		c.Locals("is_superuser", true)

		return c.Next()
	}
}
