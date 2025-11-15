package acceptance_test

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

var _ = Describe("User Registration Acceptance Scenarios", func() {
	var (
		builder *AcceptanceTestBuilder
		handler *command_handler.UserHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = NewAcceptanceTestBuilder(nil)
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("Complete user registration and login flow", func() {
		It("کاربر جدید می‌تواند ثبت‌نام کند و سپس وارد سیستم شود", func() {
			// Step 1: Register new user
			registerCmd := &commands.RegisterUser{
				AvatarIdentifier: "avatar123",
				UserName:         "newuser",
				FirstName:        "John",
				LastName:         "Doe",
				Email:            "newuser@example.com",
				Password:         "password123",
			}

			err := handler.RegisterHandler(ctx, registerCmd)
			Expect(err).NotTo(HaveOccurred())

			// Step 2: Verify user was created in database
			userRepo := repository.NewUserRepository(builder.db)
			user, err := userRepo.FindByUserName(ctx, "newuser")
			Expect(err).NotTo(HaveOccurred())
			Expect(user).NotTo(BeNil())
			Expect(user.UserName).To(Equal("newuser"))
			Expect(user.Email).To(Equal("newuser@example.com"))

			// Step 3: Verify password is hashed
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123"))
			Expect(err).NotTo(HaveOccurred())

			// Step 4: Login with registered credentials
			loginCmd := &commands.LoginUser{
				UserName: "newuser",
				Password: "password123",
			}

			token, err := handler.LoginHandler(ctx, loginCmd)
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeEmpty())

			// Step 5: Verify token was saved
			tokenRepo := repository.NewTokenRepository(builder.db)
			savedToken, err := tokenRepo.FindByUserID(ctx, user.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(savedToken).NotTo(BeNil())
			Expect(savedToken.Token).To(Equal(token))
		})
	})

	Describe("Duplicate username prevention", func() {
		It("کاربر نمی‌تواند با نام کاربری تکراری ثبت‌نام کند", func() {
			// Step 1: Register first user
			cmd1 := &commands.RegisterUser{
				UserName:  "existinguser",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "user1@example.com",
				Password:  "password123",
			}
			err := handler.RegisterHandler(ctx, cmd1)
			Expect(err).NotTo(HaveOccurred())

			// Step 2: Try to register with same username
			cmd2 := &commands.RegisterUser{
				UserName:  "existinguser",
				FirstName: "Jane",
				LastName:  "Smith",
				Email:     "user2@example.com",
				Password:  "password456",
			}
			err = handler.RegisterHandler(ctx, cmd2)
			Expect(err).To(HaveOccurred())
			appErr, ok := err.(apperrors.Error)
			Expect(ok).To(BeTrue())
			Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeConflict))

			// Step 3: Verify only one user exists
			userRepo := repository.NewUserRepository(builder.db)
			users, err := userRepo.FindByField(ctx, "user_name", "existinguser")
			Expect(err).NotTo(HaveOccurred())
			Expect(users).To(HaveLen(1))
		})
	})

	Describe("User logout flow", func() {
		It("کاربر می‌تواند از سیستم خارج شود", func() {
			// Step 1: Register and login
			registerCmd := &commands.RegisterUser{
				UserName:  "logoutuser",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "logout@example.com",
				Password:  "password123",
			}
			err := handler.RegisterHandler(ctx, registerCmd)
			Expect(err).NotTo(HaveOccurred())

			loginCmd := &commands.LoginUser{
				UserName: "logoutuser",
				Password: "password123",
			}
			_, err = handler.LoginHandler(ctx, loginCmd)
			Expect(err).NotTo(HaveOccurred())

			// Step 2: Get user ID
			userRepo := repository.NewUserRepository(builder.db)
			user, err := userRepo.FindByUserName(ctx, "logoutuser")
			Expect(err).NotTo(HaveOccurred())

			// Step 3: Logout
			logoutCmd := &commands.Logout{
				UserID: uint64(user.ID),
			}
			err = handler.LogoutHandler(ctx, logoutCmd)
			Expect(err).NotTo(HaveOccurred())

			// Step 4: Verify token was removed
			tokenRepo := repository.NewTokenRepository(builder.db)
			_, err = tokenRepo.FindByUserID(ctx, user.ID)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(repository.ErrTokenNotFound))
		})
	})
})

// AcceptanceTestBuilder helps build acceptance test scenarios
type AcceptanceTestBuilder struct {
	db  *gorm.DB
	uow unit_of_work.PGUnitOfWork
	cfg *config.Config
}

func NewAcceptanceTestBuilder(t GinkgoTInterface) *AcceptanceTestBuilder {
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
			Secret:                    "test-secret-for-acceptance",
			AccessTokenExpireDuration: 3600,
		},
	}

	return &AcceptanceTestBuilder{
		db:  db,
		uow: uow,
		cfg: cfg,
	}
}

func (b *AcceptanceTestBuilder) BuildHandler() *command_handler.UserHandler {
	return command_handler.NewUserHandler(b.uow, b.cfg)
}

func (b *AcceptanceTestBuilder) Cleanup() {
	b.db.Exec("DELETE FROM users")
	b.db.Exec("DELETE FROM tokens")
	b.db.Exec("DELETE FROM profiles")
}
