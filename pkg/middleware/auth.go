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
			user, err := m.Uow.Token(ctx).FindByUserID(ctx, entity.UserID(userID))
			if err != nil {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token not found"})
			}
			if user.Token != tokenStr {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token mismatch"})
			}

			// Store user_id in Fiber context
			c.Locals("user_id", userID)
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		return c.Next()
	}
}
