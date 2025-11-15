package integration_test

import (
	"context"

	"shikposh-backend/config"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/account/service_layer/command_handler"
	"shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("UserHandler Integration", func() {
	var (
		builder *IntegrationTestBuilder
		handler *command_handler.UserHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = NewIntegrationTestBuilder(nil)
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("RegisterHandler", func() {
		Context("when registering a new user", func() {
			It("should register and persist to database", func() {
				cmd := &commands.RegisterUser{
					AvatarIdentifier: "avatar123",
					UserName:         "newuser",
					FirstName:        "John",
					LastName:         "Doe",
					Email:            "newuser@example.com",
					Password:         "password123",
				}

				err := handler.RegisterHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())

				// Verify user was persisted
				userRepo := repository.NewUserRepository(builder.db)
				user, err := userRepo.FindByUserName(ctx, "newuser")
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(user.UserName).To(Equal("newuser"))
				Expect(user.Email).To(Equal("newuser@example.com"))

				// Verify password is hashed
				err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123"))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when username already exists", func() {
			It("should return conflict error", func() {
				// Create existing user
				createTestUser(nil, builder.db, "existinguser", "existing@example.com", "password123")

				cmd := &commands.RegisterUser{
					UserName:  "existinguser",
					FirstName: "Test",
					LastName:  "User",
					Email:     "test@example.com",
					Password:  "password123",
				}

				err := handler.RegisterHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeConflict))
			})
		})
	})

	Describe("LoginHandler", func() {
		Context("when credentials are valid", func() {
			It("should login and create token in database", func() {
				// Create user
				user := createTestUser(nil, builder.db, "testuser", "test@example.com", "password123")

				cmd := &commands.LoginUser{
					UserName: "testuser",
					Password: "password123",
				}

				token, err := handler.LoginHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
				Expect(token).NotTo(BeEmpty())

				// Verify token was saved
				tokenRepo := repository.NewTokenRepository(builder.db)
				savedToken, err := tokenRepo.FindByUserID(ctx, user.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(savedToken).NotTo(BeNil())
				Expect(savedToken.Token).To(Equal(token))
			})
		})

		Context("when user logs in again", func() {
			It("should replace existing token", func() {
				user := createTestUser(nil, builder.db, "testuser", "test@example.com", "password123")

				// Create old token
				tokenRepo := repository.NewTokenRepository(builder.db)
				oldToken := entity.NewToken("old-token", user.ID)
				err := tokenRepo.Save(ctx, oldToken)
				Expect(err).NotTo(HaveOccurred())

				cmd := &commands.LoginUser{
					UserName: "testuser",
					Password: "password123",
				}

				newToken, err := handler.LoginHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
				Expect(newToken).NotTo(BeEmpty())
				Expect(newToken).NotTo(Equal("old-token"))

				// Verify old token was replaced
				savedToken, err := tokenRepo.FindByUserID(ctx, user.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(savedToken.Token).To(Equal(newToken))
			})
		})
	})

	Describe("LogoutHandler", func() {
		Context("when user has valid token", func() {
			It("should remove token from database", func() {
				user := createTestUser(nil, builder.db, "testuser", "test@example.com", "password123")

				// Create token
				tokenRepo := repository.NewTokenRepository(builder.db)
				token := entity.NewToken("test-token", user.ID)
				err := tokenRepo.Save(ctx, token)
				Expect(err).NotTo(HaveOccurred())

				cmd := &commands.Logout{
					UserID: uint64(user.ID),
				}

				err = handler.LogoutHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())

				// Verify token was removed
				_, err = tokenRepo.FindByUserID(ctx, user.ID)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(repository.ErrTokenNotFound))
			})
		})
	})
})

// IntegrationTestBuilder helps build integration test scenarios with real database
type IntegrationTestBuilder struct {
	db  *gorm.DB
	uow unit_of_work.PGUnitOfWork
	cfg *config.Config
}

func NewIntegrationTestBuilder(t GinkgoTInterface) *IntegrationTestBuilder {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Token{},
		&entity.Profile{},
	)
	Expect(err).NotTo(HaveOccurred())

	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unit_of_work.New(db, eventCh)

	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:                    "test-secret-key-for-integration-tests",
			AccessTokenExpireDuration: 3600,
		},
	}

	return &IntegrationTestBuilder{
		db:  db,
		uow: uow,
		cfg: cfg,
	}
}

func (b *IntegrationTestBuilder) BuildHandler() *command_handler.UserHandler {
	return command_handler.NewUserHandler(b.uow, b.cfg)
}

func (b *IntegrationTestBuilder) Cleanup() {
	b.db.Exec("DELETE FROM users")
	b.db.Exec("DELETE FROM tokens")
	b.db.Exec("DELETE FROM profiles")
}

// Helper functions
func createTestUser(t GinkgoTInterface, db *gorm.DB, username, email, password string) *entity.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	Expect(err).NotTo(HaveOccurred())

	user := &entity.User{
		UserName:  username,
		FirstName: "Test",
		LastName:  "User",
		Email:     email,
		Password:  string(hashedPassword),
	}

	userRepo := repository.NewUserRepository(db)
	err = userRepo.Save(context.Background(), user)
	Expect(err).NotTo(HaveOccurred())

	return user
}
