package mocks

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/entity"
	"github.com/shikposh/framework/adapter"

	"github.com/stretchr/testify/mock"
)

// MockTokenRepository is a mock implementation of TokenRepository
type MockTokenRepository struct {
	mock.Mock
}

func (m *MockTokenRepository) FindByID(ctx context.Context, id uint64) (*entity.Token, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Token), args.Error(1)
}

func (m *MockTokenRepository) FindByField(ctx context.Context, field string, value interface{}) (*entity.Token, error) {
	args := m.Called(ctx, field, value)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Token), args.Error(1)
}

func (m *MockTokenRepository) Remove(ctx context.Context, model *entity.Token, softDelete bool) error {
	args := m.Called(ctx, model, softDelete)
	return args.Error(0)
}

func (m *MockTokenRepository) Modify(ctx context.Context, model *entity.Token) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockTokenRepository) Save(ctx context.Context, model *entity.Token) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockTokenRepository) FindByUserID(ctx context.Context, userID entity.UserID) (*entity.Token, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Token), args.Error(1)
}

func (m *MockTokenRepository) Seen() []adapter.Entity {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]adapter.Entity)
}

func (m *MockTokenRepository) SetSeen(model adapter.Entity) {
	m.Called(model)
}

var _ repository.TokenRepository = (*MockTokenRepository)(nil)

