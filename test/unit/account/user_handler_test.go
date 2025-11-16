package account_test

import (
	"context"
	"errors"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/account/service_layer/command_handler"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/test/unit/testdouble/builders"
	"shikposh-backend/test/unit/testdouble/factories"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("UserHandler", func() {
	var (
		builder *builders.UserTestBuilder
		handler *command_handler.UserHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = builders.NewUserTestBuilder().
			WithUserRepo().
			WithTokenRepo().
			WithSuccessfulTransaction()
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	Describe("RegisterHandler", func() {
		Context("when registering a new user", func() {
			It("should register successfully", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateRegisterCommand("newuser", "newuser@example.com", "password123")
				builder.MockUserRepo.On("FindByUserName", mock.Anything, "newuser").
					Return(nil, repository.ErrUserNotFound).Maybe()
				builder.MockUserRepo.On("Save", mock.Anything, mock.AnythingOfType("*entity.User")).
					Return(nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.RegisterHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when username already exists", func() {
			It("should return conflict error", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateRegisterCommand("existinguser", "existing@example.com", "password123")
				existingUser := factories.CreateUser("existinguser", "existing@example.com", "password123")
				builder.MockUserRepo.On("FindByUserName", mock.Anything, "existinguser").
					Return(existingUser, nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.RegisterHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
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
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateLoginCommand("existinguser", "password123")
				user := factories.CreateUser("existinguser", "user@example.com", "password123")
				builder.MockUserRepo.On("FindByUserName", mock.Anything, "existinguser").
					Return(user, nil).Maybe()
				builder.MockTokenRepo.On("FindByUserID", mock.Anything, entity.UserID(1)).
					Return(nil, repository.ErrTokenNotFound).Maybe()
				builder.MockTokenRepo.On("Save", mock.Anything, mock.AnythingOfType("*entity.Token")).
					Return(nil).Maybe()

				// Phase 2: Exercise (Act)
				token, err := handler.LoginHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				Expect(token).NotTo(BeEmpty())
			})
		})

		Context("when username does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateLoginCommand("nonexistentuser", "password123")
				builder.MockUserRepo.On("FindByUserName", mock.Anything, "nonexistentuser").
					Return(nil, repository.ErrUserNotFound).Maybe()

				// Phase 2: Exercise (Act)
				token, err := handler.LoginHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(token).To(BeEmpty())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when password is incorrect", func() {
			It("should return unauthorized error", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateLoginCommand("existinguser", "wrongpassword")
				user := factories.CreateUser("existinguser", "user@example.com", "password123")
				builder.MockUserRepo.On("FindByUserName", mock.Anything, "existinguser").
					Return(user, nil).Maybe()

				// Phase 2: Exercise (Act)
				token, err := handler.LoginHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
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
				// Phase 1: Setup (Arrange)
				cmd := &commands.Logout{UserID: 1}
				token := &entity.Token{
					ID:     1,
					UserID: 1,
					Token:  "test-token",
				}
				builder.MockTokenRepo.On("FindByUserID", mock.Anything, entity.UserID(1)).
					Return(token, nil).Maybe()
				builder.MockTokenRepo.On("Remove", mock.Anything, token, false).
					Return(nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.LogoutHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when token does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				cmd := &commands.Logout{UserID: 999}
				builder.MockTokenRepo.On("FindByUserID", mock.Anything, entity.UserID(999)).
					Return(nil, repository.ErrTokenNotFound).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.LogoutHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				var appErr apperrors.Error
				if errors.As(err, &appErr) {
					Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
				} else {
					Expect(err.Error()).To(ContainSubstring("not found"))
				}
			})
		})
	})
})
