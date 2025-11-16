package mocks

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/entity"
	"github.com/ali-mahdavi-dev/framework/adapter"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindByField(ctx context.Context, field string, value interface{}) (*entity.User, error) {
	args := m.Called(ctx, field, value)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Remove(ctx context.Context, model *entity.User, softDelete bool) error {
	args := m.Called(ctx, model, softDelete)
	return args.Error(0)
}

func (m *MockUserRepository) Modify(ctx context.Context, model *entity.User) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockUserRepository) Save(ctx context.Context, model *entity.User) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockUserRepository) FindByUserName(ctx context.Context, username string) (*entity.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsernameExcludingID(ctx context.Context, username string, id uint) (*entity.User, error) {
	args := m.Called(ctx, username, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Seen() []adapter.Entity {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]adapter.Entity)
}

func (m *MockUserRepository) SetSeen(model adapter.Entity) {
	m.Called(model)
}

var _ repository.UserRepository = (*MockUserRepository)(nil)

