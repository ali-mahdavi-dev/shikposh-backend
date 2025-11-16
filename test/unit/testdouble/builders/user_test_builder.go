package builders

import (
	"context"

	"shikposh-backend/config"
	"shikposh-backend/internal/account/service_layer/command_handler"
	"github.com/shikposh/framework/service_layer/types"
	"shikposh-backend/test/unit/testdouble/mocks"

	"github.com/stretchr/testify/mock"
)

// UserTestBuilder helps build test scenarios with mocks
type UserTestBuilder struct {
	MockUOW       *mocks.MockPGUnitOfWork
	MockUserRepo  *mocks.MockUserRepository
	MockTokenRepo *mocks.MockTokenRepository
	cfg           *config.Config
}

func NewUserTestBuilder() *UserTestBuilder {
	return &UserTestBuilder{
		MockUOW:       new(mocks.MockPGUnitOfWork),
		MockUserRepo:  new(mocks.MockUserRepository),
		MockTokenRepo: new(mocks.MockTokenRepository),
		cfg: &config.Config{
			JWT: config.JWTConfig{
				Secret:                    "test-secret-key-for-jwt-token-generation",
				AccessTokenExpireDuration: 3600,
			},
		},
	}
}

func (b *UserTestBuilder) BuildHandler() *command_handler.UserHandler {
	return command_handler.NewUserHandler(b.MockUOW, b.cfg)
}

func (b *UserTestBuilder) WithUserRepo() *UserTestBuilder {
	b.MockUOW.On("User", mock.Anything).Return(b.MockUserRepo).Maybe()
	return b
}

func (b *UserTestBuilder) WithTokenRepo() *UserTestBuilder {
	b.MockUOW.On("Token", mock.Anything).Return(b.MockTokenRepo).Maybe()
	return b
}

func (b *UserTestBuilder) WithSuccessfulTransaction() *UserTestBuilder {
	b.MockUOW.On("Do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fc := args.Get(1).(types.UowUseCase)
		fc(args.Get(0).(context.Context))
	}).Maybe()
	return b
}
