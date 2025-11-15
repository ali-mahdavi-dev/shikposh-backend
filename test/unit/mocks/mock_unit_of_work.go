package mocks

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	productrepository "shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/pkg/framework/service_layer/types"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockPGUnitOfWork is a mock implementation of PGUnitOfWork
type MockPGUnitOfWork struct {
	mock.Mock
}

func (m *MockPGUnitOfWork) Do(ctx context.Context, fc types.UowUseCase) error {
	args := m.Called(ctx, fc)
	// If no error is set, execute the function
	if args.Error(0) == nil {
		return fc(ctx)
	}
	return args.Error(0)
}

func (m *MockPGUnitOfWork) GetSession(ctx context.Context) *gorm.DB {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}

func (m *MockPGUnitOfWork) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPGUnitOfWork) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPGUnitOfWork) User(ctx context.Context) repository.UserRepository {
	args := m.Called(ctx)
	return args.Get(0).(repository.UserRepository)
}

func (m *MockPGUnitOfWork) Token(ctx context.Context) repository.TokenRepository {
	args := m.Called(ctx)
	return args.Get(0).(repository.TokenRepository)
}

func (m *MockPGUnitOfWork) Profile(ctx context.Context) repository.ProfileRepository {
	args := m.Called(ctx)
	return args.Get(0).(repository.ProfileRepository)
}

func (m *MockPGUnitOfWork) Product(ctx context.Context) productrepository.ProductRepository {
	args := m.Called(ctx)
	return args.Get(0).(productrepository.ProductRepository)
}

func (m *MockPGUnitOfWork) Category(ctx context.Context) productrepository.CategoryRepository {
	args := m.Called(ctx)
	return args.Get(0).(productrepository.CategoryRepository)
}

func (m *MockPGUnitOfWork) Review(ctx context.Context) productrepository.ReviewRepository {
	args := m.Called(ctx)
	return args.Get(0).(productrepository.ReviewRepository)
}

func (m *MockPGUnitOfWork) Outbox(ctx context.Context) productrepository.OutboxRepository {
	args := m.Called(ctx)
	return args.Get(0).(productrepository.OutboxRepository)
}

var _ unit_of_work.PGUnitOfWork = (*MockPGUnitOfWork)(nil)

