package middleware

import (
	"net/http"
	"strings"

	"shikposh-backend/internal/account/domain/entity"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

func (m *Middleware) AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get token from Authorization header (frontend manages tokens)
		var tokenStr string

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			if !m.IsAuthenticated {
				return c.Next()
			}
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization required"})
		}

		// Extract token from Authorization header
		// Support both "Bearer <token>" and just "<token>" formats
		prefix := "Bearer "
		if strings.HasPrefix(authHeader, prefix) {
			tokenStr = authHeader[len(prefix):]
		} else {
			tokenStr = authHeader
		}

		if tokenStr == "" {
			if !m.IsAuthenticated {
				return c.Next()
			}
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization required"})
		}

		// Parse JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.Cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Extract claims and validate token from DB
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := cast.ToUint64(claims["user_id"])
			ctx := c.Context()
			tokenEntity, err := m.Uow.Token(ctx).FindByUserID(ctx, entity.UserID(userID))
			if err != nil {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token not found"})
			}
			if tokenEntity.Token != tokenStr {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token mismatch"})
			}

			// Get user from database and store in context to avoid repeated queries
			user, err := m.Uow.User(ctx).FindByID(ctx, uint64(userID))
			if err != nil {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
			}

			// Store user_id and user in Fiber context for later use
			c.Locals("user_id", userID)
			c.Locals("user", user)
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		return c.Next()
	}
}
