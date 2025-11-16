package products_test

import (
	"context"
	"errors"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/service_layer/command_handler"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/test/unit/testdouble/builders"
	"shikposh-backend/test/unit/testdouble/factories"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("ProductCommandHandler", func() {
	var (
		builder *builders.ProductTestBuilder
		handler *command_handler.ProductCommandHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = builders.NewProductTestBuilder().
			WithProductRepo().
			WithCategoryRepo().
			WithSuccessfulTransaction()
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	Describe("CreateProductHandler", func() {
		Context("when creating a new product", func() {
			It("should create product successfully", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateProductCommand("Men's T-Shirt", "Test Brand", 1)
				category := factories.CreateCategory(1, "Clothing", "clothing")
				builder.MockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.MockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(nil, repository.ErrProductNotFound).Maybe()
				builder.MockProductRepo.On("Save", mock.Anything, mock.AnythingOfType("*product_aggregate.Product")).
					Return(nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when category does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateProductCommand("Test Product", "Test Brand", 999)
				builder.MockCategoryRepo.On("FindByID", mock.Anything, uint64(999)).
					Return(nil, appadapter.ErrEntityNotFound).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when product slug already exists", func() {
			It("should return conflict error", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateProductCommand("Test Product", "Test Brand", 1)
				category := factories.CreateCategory(1, "Clothing", "clothing")
				existingProduct := factories.CreateProduct(1, "Existing Product", "duplicate-slug", "Brand", 1)
				builder.MockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.MockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(existingProduct, nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeConflict))
			})
		})

		Context("when product has no price details", func() {
			It("should return validation error", func() {
				// Phase 1: Setup (Arrange)
				desc := "Product description"
				cmd := &commands.CreateProduct{
					Name:        "Product Without Price",
					Brand:       "Brand",
					Description: &desc,
					CategoryID:  1,
				}
				category := factories.CreateCategory(1, "Clothing", "clothing")
				builder.MockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.MockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(nil, repository.ErrProductNotFound).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
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
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateUpdateProductCommand(1, "Updated Product", "New Brand", 1)
				product := factories.CreateProduct(1, "Old Product", "old-product", "Old Brand", 1)
				product.Details = []productaggregate.ProductDetail{
					{ProductID: product.ID, Price: 100000.0},
				}
				category := factories.CreateCategory(1, "Clothing", "clothing")
				builder.MockProductRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(product, nil).Maybe()
				builder.MockCategoryRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(category, nil).Maybe()
				builder.MockProductRepo.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).
					Return(nil, repository.ErrProductNotFound).Maybe()
				builder.MockProductRepo.On("ClearDetails", mock.Anything, product).
					Return(nil).Maybe()
				builder.MockProductRepo.On("Modify", mock.Anything, mock.AnythingOfType("*product_aggregate.Product")).
					Return(nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.UpdateProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				cmd := factories.CreateUpdateProductCommand(999, "Nonexistent Product", "Brand", 1)
				builder.MockProductRepo.On("FindByID", mock.Anything, uint64(999)).
					Return(nil, appadapter.ErrEntityNotFound).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.UpdateProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
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
				// Phase 1: Setup (Arrange)
				cmd := &commands.DeleteProduct{
					ID:         1,
					SoftDelete: true,
				}
				product := factories.CreateProduct(1, "Product To Delete", "product-to-delete", "Brand", 1)
				builder.MockProductRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(product, nil).Maybe()
				builder.MockProductRepo.On("ClearAllAssociations", mock.Anything, product).
					Return(nil).Maybe()
				builder.MockProductRepo.On("Remove", mock.Anything, product, true).
					Return(nil).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.DeleteProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				cmd := &commands.DeleteProduct{
					ID:         999,
					SoftDelete: false,
				}
				builder.MockProductRepo.On("FindByID", mock.Anything, uint64(999)).
					Return(nil, repository.ErrProductNotFound).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.DeleteProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when database error occurs during deletion", func() {
			It("should return error", func() {
				// Phase 1: Setup (Arrange)
				cmd := &commands.DeleteProduct{
					ID:         1,
					SoftDelete: false,
				}
				product := factories.CreateProduct(1, "Product", "product", "Brand", 1)
				builder.MockProductRepo.On("FindByID", mock.Anything, uint64(1)).
					Return(product, nil).Maybe()
				builder.MockProductRepo.On("ClearAllAssociations", mock.Anything, product).
					Return(nil).Maybe()
				builder.MockProductRepo.On("Remove", mock.Anything, product, false).
					Return(errors.New("database deletion error")).Maybe()

				// Phase 2: Exercise (Act)
				err := handler.DeleteProductHandler(ctx, cmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
