package mocks

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/pkg/framework/adapter"

	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of ProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByID(ctx context.Context, id uint64) (*productaggregate.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) FindByField(ctx context.Context, field string, value interface{}) (*productaggregate.Product, error) {
	args := m.Called(ctx, field, value)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) Remove(ctx context.Context, model *productaggregate.Product, softDelete bool) error {
	args := m.Called(ctx, model, softDelete)
	return args.Error(0)
}

func (m *MockProductRepository) Modify(ctx context.Context, model *productaggregate.Product) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockProductRepository) Save(ctx context.Context, model *productaggregate.Product) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockProductRepository) GetAll(ctx context.Context) ([]*productaggregate.Product, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) FindBySlug(ctx context.Context, slug string) (*productaggregate.Product, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) FindByCategoryID(ctx context.Context, categoryID entity.CategoryID) ([]*productaggregate.Product, error) {
	args := m.Called(ctx, categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) FindByCategorySlug(ctx context.Context, categorySlug string) ([]*productaggregate.Product, error) {
	args := m.Called(ctx, categorySlug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) FindFeatured(ctx context.Context) ([]*productaggregate.Product, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) Search(ctx context.Context, query string) ([]*productaggregate.Product, error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) Filter(ctx context.Context, filters repository.ProductFilters) ([]*productaggregate.Product, error) {
	args := m.Called(ctx, filters)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*productaggregate.Product), args.Error(1)
}

func (m *MockProductRepository) ClearFeatures(ctx context.Context, product *productaggregate.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) ClearDetails(ctx context.Context, product *productaggregate.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) ClearSpecs(ctx context.Context, product *productaggregate.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) ClearAllAssociations(ctx context.Context, product *productaggregate.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Seen() []adapter.Entity {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]adapter.Entity)
}

func (m *MockProductRepository) SetSeen(model adapter.Entity) {
	m.Called(model)
}

var _ repository.ProductRepository = (*MockProductRepository)(nil)

