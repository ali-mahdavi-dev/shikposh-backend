package e2e_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"shikposh-backend/config"
	account "shikposh-backend/internal/account"
	"shikposh-backend/internal/account/domain/commands"

	"github.com/gofiber/fiber/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("User API E2E", func() {
	var (
		builder *E2ETestBuilder
	)

	BeforeEach(func() {
		builder = NewE2ETestBuilder()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("POST /api/v1/public/register", func() {
		Context("when registering a new user", func() {
			It("should register user via HTTP API", func() {
				cmd := commands.RegisterUser{
					AvatarIdentifier: "avatar123",
					UserName:         "newuser",
					FirstName:        "John",
					LastName:         "Doe",
					Email:            "newuser@example.com",
					Password:         "password123",
				}

				body, err := json.Marshal(cmd)
				Expect(err).NotTo(HaveOccurred())

				req := httptest.NewRequest(http.MethodPost, "/api/v1/public/register", bytes.NewBuffer(body))
				req.Header.Set("Content-Type", "application/json")

				resp, err := builder.app.Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			})
		})

		Context("when username already exists", func() {
			It("should return conflict status", func() {
				// Register first user
				cmd1 := commands.RegisterUser{
					UserName:  "existinguser",
					FirstName: "Test",
					LastName:  "User",
					Email:     "test1@example.com",
					Password:  "password123",
				}
				body1, _ := json.Marshal(cmd1)
				req1 := httptest.NewRequest(http.MethodPost, "/api/v1/public/register", bytes.NewBuffer(body1))
				req1.Header.Set("Content-Type", "application/json")
				builder.app.Test(req1)

				// Try to register with same username
				cmd2 := commands.RegisterUser{
					UserName:  "existinguser",
					FirstName: "Another",
					LastName:  "User",
					Email:     "test2@example.com",
					Password:  "password123",
				}
				body2, _ := json.Marshal(cmd2)
				req2 := httptest.NewRequest(http.MethodPost, "/api/v1/public/register", bytes.NewBuffer(body2))
				req2.Header.Set("Content-Type", "application/json")

				resp, err := builder.app.Test(req2)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			})
		})
	})

	Describe("POST /api/v1/public/login", func() {
		Context("when credentials are valid", func() {
			It("should login and return access token", func() {
				// First register a user
				registerCmd := commands.RegisterUser{
					UserName:  "testuser",
					FirstName: "Test",
					LastName:  "User",
					Email:     "test@example.com",
					Password:  "password123",
				}
				registerBody, _ := json.Marshal(registerCmd)
				registerReq := httptest.NewRequest(http.MethodPost, "/api/v1/public/register", bytes.NewBuffer(registerBody))
				registerReq.Header.Set("Content-Type", "application/json")
				builder.app.Test(registerReq)

				// Now login
				loginCmd := commands.LoginUser{
					UserName: "testuser",
					Password: "password123",
				}
				loginBody, _ := json.Marshal(loginCmd)
				loginReq := httptest.NewRequest(http.MethodPost, "/api/v1/public/login", bytes.NewBuffer(loginBody))
				loginReq.Header.Set("Content-Type", "application/json")

				resp, err := builder.app.Test(loginReq)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				// Verify response contains access token
				var result map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&result)
				Expect(err).NotTo(HaveOccurred())
				Expect(result["data"]).NotTo(BeNil())
				if data, ok := result["data"].(map[string]interface{}); ok {
					Expect(data["access"]).NotTo(BeEmpty())
				}
			})
		})
	})
})

// E2ETestBuilder helps build E2E test scenarios with HTTP server
type E2ETestBuilder struct {
	app *fiber.App
	db  *gorm.DB
}

func NewE2ETestBuilder() *E2ETestBuilder {
	// Create in-memory database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	// Auto-migrate
	err = db.AutoMigrate(
	// Account tables will be migrated by bootstrap
	)
	Expect(err).NotTo(HaveOccurred())

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Bootstrap account module
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:                    "test-secret-for-e2e",
			AccessTokenExpireDuration: 3600,
		},
	}

	// Bootstrap account routes
	err = account.Bootstrap(app, db, cfg)
	Expect(err).NotTo(HaveOccurred())

	return &E2ETestBuilder{
		app: app,
		db:  db,
	}
}

func (b *E2ETestBuilder) Cleanup() {
	b.db.Exec("DELETE FROM users")
	b.db.Exec("DELETE FROM tokens")
	b.db.Exec("DELETE FROM profiles")
}
