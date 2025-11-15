package products_test

import (
	"context"
	"errors"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/service_layer/command_handler"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/service_layer/types"
	"shikposh-backend/test/unit/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("ProductCommandHandler", func() {
	var (
		builder *ProductTestBuilder
		handler *command_handler.ProductCommandHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = NewProductTestBuilder().
			WithProductRepo().
			WithCategoryRepo().
			WithSuccessfulTransaction()
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	Describe("CreateProductHandler", func() {
		Context("when creating a new product", func() {
			It("should create product successfully", func() {
				cmd := createProductCommand("Men's T-Shirt", "Test Brand", 1)
				category := createCategory(1, "Clothing", "clothing")

				builder.mockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.mockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(nil, repository.ErrProductNotFound).Maybe()
				builder.mockProductRepo.On("Save", mock.Anything, mock.AnythingOfType("*product_aggregate.Product")).
					Return(nil).Maybe()

				err := handler.CreateProductHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when category does not exist", func() {
			It("should return not found error", func() {
				cmd := createProductCommand("Test Product", "Test Brand", 999)

				builder.mockCategoryRepo.On("FindByID", mock.Anything, uint64(999)).
					Return(nil, appadapter.ErrEntityNotFound).Maybe()

				err := handler.CreateProductHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when product slug already exists", func() {
			It("should return conflict error", func() {
				cmd := createProductCommand("Test Product", "Test Brand", 1)
				category := createCategory(1, "Clothing", "clothing")
				existingProduct := createProduct(1, "Existing Product", "duplicate-slug", "Brand", 1)

				builder.mockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.mockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(existingProduct, nil).Maybe()

				err := handler.CreateProductHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeConflict))
			})
		})

		Context("when product has no price details", func() {
			It("should return validation error", func() {
				desc := "Product description"
				cmd := &commands.CreateProduct{
					Name:        "Product Without Price",
					Brand:       "Brand",
					Description: &desc,
					CategoryID:  1,
				}
				category := createCategory(1, "Clothing", "clothing")

				builder.mockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.mockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(nil, repository.ErrProductNotFound).Maybe()

				err := handler.CreateProductHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				var appErr apperrors.Error
				ok := errors.As(err, &appErr)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeValidation))
			})
		})
	})

	Describe("UpdateProductHandler", func() {
		Context("when updating an existing product", func() {
			It("should update product successfully", func() {
				cmd := createUpdateProductCommand(1, "Updated Product", "New Brand", 1)
				product := createProduct(1, "Old Product", "old-product", "Old Brand", 1)
				product.Details = []productaggregate.ProductDetail{
					{ProductID: product.ID, Price: 100000.0},
				}
				category := createCategory(1, "Clothing", "clothing")

				builder.mockProductRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(product, nil).Maybe()
				builder.mockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.mockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(nil, repository.ErrProductNotFound).Maybe()
				builder.mockProductRepo.On("ClearDetails", mock.Anything, product).
					Return(nil).Maybe()
				builder.mockProductRepo.On("Modify", mock.Anything, mock.AnythingOfType("*product_aggregate.Product")).
					Return(nil).Maybe()

				err := handler.UpdateProductHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				cmd := createUpdateProductCommand(999, "Nonexistent Product", "Brand", 1)

				builder.mockProductRepo.On("FindByID", mock.Anything, uint64(999)).
					Return(nil, appadapter.ErrEntityNotFound).Maybe()

				err := handler.UpdateProductHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})
	})

	Describe("DeleteProductHandler", func() {
		Context("when soft deleting a product", func() {
			It("should soft delete product successfully", func() {
				cmd := &commands.DeleteProduct{
					ID:         1,
					SoftDelete: true,
				}
				product := createProduct(1, "Product To Delete", "product-to-delete", "Brand", 1)

				builder.mockProductRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(product, nil).Maybe()
				builder.mockProductRepo.On("ClearAllAssociations", mock.Anything, product).
					Return(nil).Maybe()
				builder.mockProductRepo.On("Remove", mock.Anything, product, true).
					Return(nil).Maybe()

				err := handler.DeleteProductHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				cmd := &commands.DeleteProduct{
					ID:         999,
					SoftDelete: false,
				}

				builder.mockProductRepo.On("FindByID", mock.Anything, uint64(999)).
					Return(nil, repository.ErrProductNotFound).Maybe()

				err := handler.DeleteProductHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when database error occurs during deletion", func() {
			It("should return error", func() {
				cmd := &commands.DeleteProduct{
					ID:         1,
					SoftDelete: false,
				}
				product := createProduct(1, "Product", "product", "Brand", 1)

				builder.mockProductRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(product, nil).Maybe()
				builder.mockProductRepo.On("ClearAllAssociations", mock.Anything, product).
					Return(nil).Maybe()
				builder.mockProductRepo.On("Remove", mock.Anything, product, false).
					Return(errors.New("database deletion error")).Maybe()

				err := handler.DeleteProductHandler(ctx, cmd)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

// ProductTestBuilder helps build test scenarios for product handlers
type ProductTestBuilder struct {
	mockUOW          *mocks.MockPGUnitOfWork
	mockProductRepo  *mocks.MockProductRepository
	mockCategoryRepo *mocks.MockCategoryRepository
}

func NewProductTestBuilder() *ProductTestBuilder {
	return &ProductTestBuilder{
		mockUOW:          new(mocks.MockPGUnitOfWork),
		mockProductRepo:  new(mocks.MockProductRepository),
		mockCategoryRepo: new(mocks.MockCategoryRepository),
	}
}

func (b *ProductTestBuilder) BuildHandler() *command_handler.ProductCommandHandler {
	return command_handler.NewProductCommandHandler(b.mockUOW)
}

func (b *ProductTestBuilder) WithProductRepo() *ProductTestBuilder {
	b.mockUOW.On("Product", mock.Anything).Return(b.mockProductRepo).Maybe()
	return b
}

func (b *ProductTestBuilder) WithCategoryRepo() *ProductTestBuilder {
	b.mockUOW.On("Category", mock.Anything).Return(b.mockCategoryRepo).Maybe()
	return b
}

func (b *ProductTestBuilder) WithSuccessfulTransaction() *ProductTestBuilder {
	b.mockUOW.On("Do", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fc := args.Get(1).(types.UowUseCase)
		fc(args.Get(0).(context.Context))
	}).Maybe()
	return b
}

func createProductCommand(name, brand string, categoryID uint64) *commands.CreateProduct {
	desc := "Test product description"
	return &commands.CreateProduct{
		Name:        name,
		Slug:        command_handler.GenerateSlug(name),
		Brand:       brand,
		Description: &desc,
		CategoryID:  categoryID,
		Tags:        []string{"tag1", "tag2"},
		Sizes:       []string{"M", "L"},
		Image:       "image.jpg",
		IsNew:       true,
		IsFeatured:  false,
		Details: []commands.ProductDetailInput{
			{Price: 100000.0, Stock: 10},
		},
	}
}

func createUpdateProductCommand(id uint64, name, brand string, categoryID uint64) *commands.UpdateProduct {
	desc := "Updated description"
	slug := "slug-" + name
	return &commands.UpdateProduct{
		ID:          id,
		Name:        name,
		Slug:        slug,
		Brand:       brand,
		Description: &desc,
		CategoryID:  categoryID,
		Details: []commands.ProductDetailInput{
			{Price: 150000.0},
		},
	}
}

func createCategory(id uint64, name, slug string) *entity.Category {
	return &entity.Category{
		ID:   entity.CategoryID(id),
		Name: name,
		Slug: slug,
	}
}

func createProduct(id uint64, name, slug, brand string, categoryID uint64) *productaggregate.Product {
	return &productaggregate.Product{
		ID:         productaggregate.ProductID(id),
		Name:       name,
		Slug:       slug,
		Brand:      brand,
		CategoryID: categoryID,
	}
}
