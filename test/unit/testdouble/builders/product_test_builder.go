package builders

import (
	"context"

	"shikposh-backend/internal/products/service_layer/command_handler"
	"github.com/shikposh/framework/service_layer/types"
	"shikposh-backend/test/unit/testdouble/mocks"

	"github.com/stretchr/testify/mock"
)

// ProductTestBuilder helps build test scenarios for product handlers
type ProductTestBuilder struct {
	MockUOW          *mocks.MockPGUnitOfWork
	MockProductRepo  *mocks.MockProductRepository
	MockCategoryRepo *mocks.MockCategoryRepository
}

func NewProductTestBuilder() *ProductTestBuilder {
	return &ProductTestBuilder{
		MockUOW:          new(mocks.MockPGUnitOfWork),
		MockProductRepo:  new(mocks.MockProductRepository),
		MockCategoryRepo: new(mocks.MockCategoryRepository),
	}
}

func (b *ProductTestBuilder) BuildHandler() *command_handler.ProductCommandHandler {
	return command_handler.NewProductCommandHandler(b.MockUOW)
}

func (b *ProductTestBuilder) WithProductRepo() *ProductTestBuilder {
	b.MockUOW.On("Product", mock.Anything).Return(b.MockProductRepo).Maybe()
	return b
}

func (b *ProductTestBuilder) WithCategoryRepo() *ProductTestBuilder {
	b.MockUOW.On("Category", mock.Anything).Return(b.MockCategoryRepo).Maybe()
	return b
}

func (b *ProductTestBuilder) WithSuccessfulTransaction() *ProductTestBuilder {
	b.MockUOW.On("Do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fc := args.Get(1).(types.UowUseCase)
		fc(args.Get(0).(context.Context))
	}).Maybe()
	return b
}

