package account_test

import (
	"context"
	"errors"

	"shikposh-backend/config"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/account/service_layer/command_handler"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/service_layer/types"
	"shikposh-backend/test/unit/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("UserHandler", func() {
	var (
		builder *TestBuilder
		handler *command_handler.UserHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = NewTestBuilder().
			WithUserRepo().
			WithTokenRepo().
			WithSuccessfulTransaction()
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	Describe("RegisterHandler", func() {
		Context("when registering a new user", func() {
			It("should register successfully", func() {
				cmd := createRegisterCommand("newuser", "newuser@example.com", "password123")

				builder.mockUserRepo.On("FindByUserName", mock.Anything, "newuser").
					Return(nil, repository.ErrUserNotFound).Maybe()
				builder.mockUserRepo.On("Save", mock.Anything, mock.AnythingOfType("*entity.User")).
					Return(nil).Maybe()

				err := handler.RegisterHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when username already exists", func() {
			It("should return conflict error", func() {
				cmd := createRegisterCommand("existinguser", "existing@example.com", "password123")
				existingUser := createUser("existinguser", "existing@example.com", "password123")

				builder.mockUserRepo.On("FindByUserName", mock.Anything, "existinguser").
					Return(existingUser, nil).Maybe()

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
			It("should login successfully and return access token", func() {
				cmd := createLoginCommand("existinguser", "password123")
				user := createUser("existinguser", "user@example.com", "password123")

				builder.mockUserRepo.On("FindByUserName", mock.Anything, "existinguser").
					Return(user, nil).Maybe()
				builder.mockTokenRepo.On("FindByUserID", mock.Anything, entity.UserID(1)).
					Return(nil, repository.ErrTokenNotFound).Maybe()
				builder.mockTokenRepo.On("Save", mock.Anything, mock.AnythingOfType("*entity.Token")).
					Return(nil).Maybe()

				token, err := handler.LoginHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
				Expect(token).NotTo(BeEmpty())
			})
		})

		Context("when username does not exist", func() {
			It("should return not found error", func() {
				cmd := createLoginCommand("nonexistentuser", "password123")

				builder.mockUserRepo.On("FindByUserName", mock.Anything, "nonexistentuser").
					Return(nil, repository.ErrUserNotFound).Maybe()

				token, err := handler.LoginHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				Expect(token).To(BeEmpty())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when password is incorrect", func() {
			It("should return unauthorized error", func() {
				cmd := createLoginCommand("existinguser", "wrongpassword")
				user := createUser("existinguser", "user@example.com", "password123")

				builder.mockUserRepo.On("FindByUserName", mock.Anything, "existinguser").
					Return(user, nil).Maybe()

				token, err := handler.LoginHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				Expect(token).To(BeEmpty())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeUnauthorized))
			})
		})
	})

	Describe("LogoutHandler", func() {
		Context("when user has valid token", func() {
			It("should logout successfully", func() {
				cmd := &commands.Logout{UserID: 1}
				token := &entity.Token{
					ID:     1,
					UserID: 1,
					Token:  "test-token",
				}

				builder.mockTokenRepo.On("FindByUserID", mock.Anything, entity.UserID(1)).
					Return(token, nil).Maybe()
				builder.mockTokenRepo.On("Remove", mock.Anything, token, false).
					Return(nil).Maybe()

				err := handler.LogoutHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when token does not exist", func() {
			It("should return not found error", func() {
				cmd := &commands.Logout{UserID: 999}

				builder.mockTokenRepo.On("FindByUserID", mock.Anything, entity.UserID(999)).
					Return(nil, repository.ErrTokenNotFound).Maybe()

				err := handler.LogoutHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				// Check if error is apperrors.Error or wrapped error
				var appErr apperrors.Error
				if errors.As(err, &appErr) {
					Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
				} else {
					// If not apperrors.Error, check error message contains "not found"
					Expect(err.Error()).To(ContainSubstring("not found"))
				}
			})
		})
	})
})

// TestBuilder helps build test scenarios with mocks
type TestBuilder struct {
	mockUOW       *mocks.MockPGUnitOfWork
	mockUserRepo  *mocks.MockUserRepository
	mockTokenRepo *mocks.MockTokenRepository
	cfg           *config.Config
}

func NewTestBuilder() *TestBuilder {
	return &TestBuilder{
		mockUOW:       new(mocks.MockPGUnitOfWork),
		mockUserRepo:  new(mocks.MockUserRepository),
		mockTokenRepo: new(mocks.MockTokenRepository),
		cfg: &config.Config{
			JWT: config.JWTConfig{
				Secret:                    "test-secret-key-for-jwt-token-generation",
				AccessTokenExpireDuration: 3600,
			},
		},
	}
}

func (b *TestBuilder) BuildHandler() *command_handler.UserHandler {
	return command_handler.NewUserHandler(b.mockUOW, b.cfg)
}

func (b *TestBuilder) WithUserRepo() *TestBuilder {
	b.mockUOW.On("User", mock.Anything).Return(b.mockUserRepo).Maybe()
	return b
}

func (b *TestBuilder) WithTokenRepo() *TestBuilder {
	b.mockUOW.On("Token", mock.Anything).Return(b.mockTokenRepo).Maybe()
	return b
}

func (b *TestBuilder) WithSuccessfulTransaction() *TestBuilder {
	b.mockUOW.On("Do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fc := args.Get(1).(types.UowUseCase)
		fc(args.Get(0).(context.Context))
	}).Maybe()
	return b
}

func createRegisterCommand(username, email, password string) *commands.RegisterUser {
	return &commands.RegisterUser{
		AvatarIdentifier: "avatar123",
		UserName:         username,
		FirstName:        "John",
		LastName:         "Doe",
		Email:            email,
		Password:         password,
	}
}

func createLoginCommand(username, password string) *commands.LoginUser {
	return &commands.LoginUser{
		UserName: username,
		Password: password,
	}
}

func createUser(username, email, password string) *entity.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &entity.User{
		ID:               1,
		AvatarIdentifier: "avatar123",
		UserName:         username,
		FirstName:        "John",
		LastName:         "Doe",
		Email:            email,
		Password:         string(hashedPassword),
	}
}
