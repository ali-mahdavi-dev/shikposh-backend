package mocks

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/entity"
	"github.com/shikposh/framework/adapter"

	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository is a mock implementation of CategoryRepository
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) FindByID(ctx context.Context, id uint64) (*entity.Category, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindByField(ctx context.Context, field string, value interface{}) (*entity.Category, error) {
	args := m.Called(ctx, field, value)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) Remove(ctx context.Context, model *entity.Category, softDelete bool) error {
	args := m.Called(ctx, model, softDelete)
	return args.Error(0)
}

func (m *MockCategoryRepository) Modify(ctx context.Context, model *entity.Category) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockCategoryRepository) Save(ctx context.Context, model *entity.Category) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetAll(ctx context.Context) ([]*entity.Category, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindBySlug(ctx context.Context, slug string) (*entity.Category, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) Seen() []adapter.Entity {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]adapter.Entity)
}

func (m *MockCategoryRepository) SetSeen(model adapter.Entity) {
	m.Called(model)
}

var _ repository.CategoryRepository = (*MockCategoryRepository)(nil)

