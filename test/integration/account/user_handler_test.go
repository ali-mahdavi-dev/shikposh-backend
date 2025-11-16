package integration_test

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/service_layer/command_handler"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/test/integration/testdouble/builders"
	"shikposh-backend/test/integration/testdouble/factories"
	"shikposh-backend/test/integration/testdouble/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserHandler Integration", func() {
	var (
		builder *builders.UserIntegrationTestBuilder
		handler *command_handler.UserHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		var err error
		builder, err = builders.NewUserIntegrationTestBuilder()
		Expect(err).NotTo(HaveOccurred())
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("RegisterHandler", func() {
		Context("when registering a new user", func() {
			It("should register and persist user to database", func() {
				// Phase 1: Setup (Arrange)
				registerCmd := factories.CreateRegisterCommand("newuser", "newuser@example.com", "password123")

				// Phase 2: Exercise (Act)
				err := handler.RegisterHandler(ctx, registerCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				user := helpers.FindUserByUsername(builder.DB, "newuser")
				Expect(user.UserName).To(Equal("newuser"))
				Expect(user.Email).To(Equal("newuser@example.com"))
				Expect(helpers.IsPasswordHashed(user.Password, "password123")).To(BeTrue())
			})
		})

		Context("when username already exists", func() {
			It("should return conflict error", func() {
				// Phase 1: Setup (Arrange)
				factories.CreateUser(builder.DB, "existinguser", "existing@example.com", "password123")
				registerCmd := factories.CreateRegisterCommand("existinguser", "test@example.com", "password123")

				// Phase 2: Exercise (Act)
				err := handler.RegisterHandler(ctx, registerCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeConflict))
			})
		})
	})

	Describe("LoginHandler", func() {
		Context("when credentials are valid", func() {
			It("should login and create token in database", func() {
				// Phase 1: Setup (Arrange)
				user := factories.CreateUser(builder.DB, "testuser", "test@example.com", "password123")
				loginCmd := factories.CreateLoginCommand("testuser", "password123")

				// Phase 2: Exercise (Act)
				token, err := handler.LoginHandler(ctx, loginCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				Expect(token).NotTo(BeEmpty())
				savedToken := helpers.FindTokenByUserID(builder.DB, user.ID)
				Expect(savedToken.Token).To(Equal(token))
			})
		})

		Context("when user logs in again", func() {
			It("should replace existing token", func() {
				// Phase 1: Setup (Arrange)
				user := factories.CreateUser(builder.DB, "testuser", "test@example.com", "password123")
				factories.CreateToken(builder.DB, user.ID, "old-token")
				loginCmd := factories.CreateLoginCommand("testuser", "password123")

				// Phase 2: Exercise (Act)
				newToken, err := handler.LoginHandler(ctx, loginCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				Expect(newToken).NotTo(Equal("old-token"))
				savedToken := helpers.FindTokenByUserID(builder.DB, user.ID)
				Expect(savedToken.Token).To(Equal(newToken))
			})
		})
	})

	Describe("LogoutHandler", func() {
		Context("when user has valid token", func() {
			It("should remove token from database", func() {
				// Phase 1: Setup (Arrange)
				user := factories.CreateUser(builder.DB, "testuser", "test@example.com", "password123")
				factories.CreateToken(builder.DB, user.ID, "test-token")
				logoutCmd := &commands.Logout{UserID: uint64(user.ID)}

				// Phase 2: Exercise (Act)
				err := handler.LogoutHandler(ctx, logoutCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				_, err = helpers.FindTokenByUserIDWithError(builder.DB, user.ID)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(repository.ErrTokenNotFound))
			})
		})
	})
})
