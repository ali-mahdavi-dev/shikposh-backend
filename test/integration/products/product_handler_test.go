package integration_test

import (
	"context"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/service_layer/command_handler"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/test/integration/testdouble/builders"
	"shikposh-backend/test/integration/testdouble/factories"
	"shikposh-backend/test/integration/testdouble/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductCommandHandler Integration", func() {
	var (
		builder *builders.ProductIntegrationTestBuilder
		handler *command_handler.ProductCommandHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		var err error
		builder, err = builders.NewProductIntegrationTestBuilder()
		Expect(err).NotTo(HaveOccurred())
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("CreateProductHandler", func() {
		Context("when creating a new product", func() {
			It("should create product with all associations in database", func() {
				// Phase 1: Setup (Arrange)
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				productCmd := factories.CreateProductCommand("Men's T-Shirt", "Test Brand", uint64(category.ID))
				productCmd.Features = []commands.ProductFeatureInput{
					{Feature: "Waterproof", Order: 1},
					{Feature: "Washable", Order: 2},
				}
				productCmd.Specs = []commands.ProductSpecInput{
					{Key: "Material", Value: "100% Cotton", Order: 1},
				}

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, productCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				product := helpers.FindProductBySlug(builder.DB, command_handler.GenerateSlug(productCmd.Name))
				Expect(product.Name).To(Equal("Men's T-Shirt"))
				Expect(product.Features).To(HaveLen(2))
				Expect(product.Details).To(HaveLen(1))
				Expect(product.Specs).To(HaveLen(1))
				Expect(product.Tags).To(Equal([]string{"tag1"}))
				Expect(product.Sizes).To(Equal([]string{"M", "L"}))
			})
		})

		Context("when slug already exists", func() {
			It("should return conflict error", func() {
				// Phase 1: Setup (Arrange)
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				firstProduct := factories.CreateProductCommand("First Product", "Brand", uint64(category.ID))

				// Phase 2: Exercise (Act) - Create first product
				err := handler.CreateProductHandler(ctx, firstProduct)
				Expect(err).NotTo(HaveOccurred())

				// Phase 1: Setup (Arrange) - Prepare duplicate product
				duplicateProduct := factories.CreateProductCommand("First Product", "Other Brand", uint64(category.ID))

				// Phase 2: Exercise (Act) - Try to create duplicate
				err = handler.CreateProductHandler(ctx, duplicateProduct)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeConflict))
			})
		})

		Context("when category does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				nonExistentCategoryID := uint64(99999)
				productCmd := factories.CreateProductCommand("Product", "Brand", nonExistentCategoryID)

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, productCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when product validation fails", func() {
			It("should return validation error when name is empty", func() {
				// Phase 1: Setup (Arrange)
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				productCmd := factories.CreateProductCommandWithEmptyName(uint64(category.ID))

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, productCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeValidation))
			})

			It("should return validation error when no details provided", func() {
				// Phase 1: Setup (Arrange)
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				productCmd := factories.CreateProductCommandWithoutDetails("Product", "Brand", uint64(category.ID))

				// Phase 2: Exercise (Act)
				err := handler.CreateProductHandler(ctx, productCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeValidation))
			})
		})
	})

	Describe("UpdateProductHandler", func() {
		Context("when updating an existing product", func() {
			It("should update product and replace associations", func() {
				// Phase 1: Setup (Arrange) - Create product
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				existingProduct := factories.CreateProductCommand("Old Product", "Old Brand", uint64(category.ID))
				existingProduct.Features = []commands.ProductFeatureInput{
					{Feature: "Old Feature", Order: 1},
				}
				err := handler.CreateProductHandler(ctx, existingProduct)
				Expect(err).NotTo(HaveOccurred())

				product := helpers.FindProductBySlug(builder.DB, command_handler.GenerateSlug(existingProduct.Name))
				updateCmd := factories.CreateUpdateCommand(uint64(product.ID), "New Product", "New Brand", uint64(category.ID))

				// Phase 2: Exercise (Act)
				err = handler.UpdateProductHandler(ctx, updateCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				updatedProduct := helpers.FindProductByID(builder.DB, uint64(product.ID))
				Expect(updatedProduct.Name).To(Equal("New Product"))
				Expect(updatedProduct.Brand).To(Equal("New Brand"))
				Expect(updatedProduct.Features).To(HaveLen(1))
				Expect(updatedProduct.Features[0].Feature).To(Equal("New Feature"))
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				nonExistentProductID := uint64(99999)
				updateCmd := factories.CreateUpdateCommand(nonExistentProductID, "Product", "Brand", uint64(category.ID))

				// Phase 2: Exercise (Act)
				err := handler.UpdateProductHandler(ctx, updateCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})

		Context("when updating with duplicate slug", func() {
			It("should return conflict error", func() {
				// Phase 1: Setup (Arrange) - Create products
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				firstProduct := factories.CreateProductCommand("Product One", "Brand1", uint64(category.ID))
				err := handler.CreateProductHandler(ctx, firstProduct)
				Expect(err).NotTo(HaveOccurred())

				secondProduct := factories.CreateProductCommand("Product Two", "Brand2", uint64(category.ID))
				err = handler.CreateProductHandler(ctx, secondProduct)
				Expect(err).NotTo(HaveOccurred())

				product2 := helpers.FindProductBySlug(builder.DB, command_handler.GenerateSlug(secondProduct.Name))
				updateCmd := factories.CreateUpdateCommandWithDuplicateSlug(uint64(product2.ID), command_handler.GenerateSlug(firstProduct.Name), uint64(category.ID))

				// Phase 2: Exercise (Act)
				err = handler.UpdateProductHandler(ctx, updateCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeConflict))
			})
		})
	})

	Describe("DeleteProductHandler", func() {
		Context("when soft deleting a product", func() {
			It("should soft delete product and keep associations", func() {
				// Phase 1: Setup (Arrange) - Create product
				category := factories.CreateCategory(builder.DB, "Clothing", "clothing")
				productCmd := factories.CreateProductCommand("Product To Delete", "Brand", uint64(category.ID))
				err := handler.CreateProductHandler(ctx, productCmd)
				Expect(err).NotTo(HaveOccurred())

				product := helpers.FindProductBySlug(builder.DB, command_handler.GenerateSlug(productCmd.Name))
				deleteCmd := &commands.DeleteProduct{
					ID:         uint64(product.ID),
					SoftDelete: true,
				}

				// Phase 2: Exercise (Act)
				err = handler.DeleteProductHandler(ctx, deleteCmd)

				// Phase 3: Verify (Assert)
				Expect(err).NotTo(HaveOccurred())
				_, err = helpers.FindProductByIDWithError(builder.DB, uint64(product.ID))
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(appadapter.ErrEntityNotFound))
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				// Phase 1: Setup (Arrange)
				nonExistentProductID := uint64(99999)
				deleteCmd := &commands.DeleteProduct{
					ID:         nonExistentProductID,
					SoftDelete: true,
				}

				// Phase 2: Exercise (Act)
				err := handler.DeleteProductHandler(ctx, deleteCmd)

				// Phase 3: Verify (Assert)
				Expect(err).To(HaveOccurred())
				Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeNotFound))
			})
		})
	})
})
