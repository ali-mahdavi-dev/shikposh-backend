package acceptance_test

import (
	"context"

	"shikposh-backend/internal/account/service_layer/command_handler"
	apperrors "github.com/shikposh/framework/errors"
	"shikposh-backend/test/acceptance/testdouble/builders"
	"shikposh-backend/test/acceptance/testdouble/factories"
	"shikposh-backend/test/acceptance/testdouble/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Registration Acceptance Scenarios", func() {
	var (
		builder *builders.UserAcceptanceTestBuilder
		factory *factories.UserFactory
		handler *command_handler.UserHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = builders.NewUserAcceptanceTestBuilder()
		factory = factories.NewUserFactory(builder.DB)
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("Complete user registration and login flow", func() {
		It("کاربر جدید می‌تواند ثبت‌نام کند و سپس وارد سیستم شود", func() {
			// Phase 1: Setup (Arrange)
			registerCmd := factory.CreateRegisterCommand("newuser", "newuser@example.com", "password123")

			// Phase 2: Exercise (Act) - Register user
			err := handler.RegisterHandler(ctx, registerCmd)

			// Phase 3: Verify (Assert) - Verify user registration
			Expect(err).NotTo(HaveOccurred())
			user := helpers.FindUserByUsername(builder.DB, "newuser")
			Expect(user.UserName).To(Equal("newuser"))
			Expect(user.Email).To(Equal("newuser@example.com"))
			helpers.VerifyPasswordHashed(user.Password, "password123")

			// Phase 1: Setup (Arrange) - Prepare login command
			loginCmd := factory.CreateLoginCommand("newuser", "password123")

			// Phase 2: Exercise (Act) - Login user
			token, err := handler.LoginHandler(ctx, loginCmd)

			// Phase 3: Verify (Assert) - Verify login and token
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeEmpty())
			savedToken := helpers.FindTokenByUserID(builder.DB, user.ID)
			Expect(savedToken.Token).To(Equal(token))
		})
	})

	Describe("Duplicate username prevention", func() {
		It("کاربر نمی‌تواند با نام کاربری تکراری ثبت‌نام کند", func() {
			// Phase 1: Setup (Arrange)
			cmd1 := factory.CreateRegisterCommand("existinguser", "user1@example.com", "password123")

			// Phase 2: Exercise (Act) - Register first user
			err := handler.RegisterHandler(ctx, cmd1)
			Expect(err).NotTo(HaveOccurred())

			// Phase 1: Setup (Arrange) - Prepare duplicate registration command
			cmd2 := factory.CreateRegisterCommand("existinguser", "user2@example.com", "password456")

			// Phase 2: Exercise (Act) - Try to register with duplicate username
			err = handler.RegisterHandler(ctx, cmd2)

			// Phase 3: Verify (Assert) - Verify conflict error and user count
			Expect(err).To(HaveOccurred())
			Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeConflict))
			helpers.VerifyUserCount(builder.DB, "existinguser", 1)
		})
	})

	Describe("User logout flow", func() {
		It("کاربر می‌تواند از سیستم خارج شود", func() {
			// Phase 1: Setup (Arrange) - Register user
			registerCmd := factory.CreateRegisterCommand("logoutuser", "logout@example.com", "password123")

			// Phase 2: Exercise (Act) - Register user
			err := handler.RegisterHandler(ctx, registerCmd)
			Expect(err).NotTo(HaveOccurred())

			// Phase 1: Setup (Arrange) - Login user
			loginCmd := factory.CreateLoginCommand("logoutuser", "password123")

			// Phase 2: Exercise (Act) - Login user
			_, err = handler.LoginHandler(ctx, loginCmd)
			Expect(err).NotTo(HaveOccurred())

			// Phase 1: Setup (Arrange) - Prepare logout command
			user := helpers.FindUserByUsername(builder.DB, "logoutuser")
			logoutCmd := factory.CreateLogoutCommand(uint64(user.ID))

			// Phase 2: Exercise (Act) - Logout user
			err = handler.LogoutHandler(ctx, logoutCmd)

			// Phase 3: Verify (Assert) - Verify token removal
			Expect(err).NotTo(HaveOccurred())
			helpers.VerifyTokenRemoved(builder.DB, user.ID)
		})
	})
})
