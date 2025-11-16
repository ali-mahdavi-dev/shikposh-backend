package acceptance_test

import (
	"context"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/service_layer/command_handler"
	apperrors "github.com/ali-mahdavi-dev/framework/errors"
	"shikposh-backend/test/acceptance/testdouble/builders"
	"shikposh-backend/test/acceptance/testdouble/factories"
	"shikposh-backend/test/acceptance/testdouble/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Management Acceptance Scenarios", func() {
	var (
		builder *builders.ProductAcceptanceTestBuilder
		factory *factories.ProductFactory
		handler *command_handler.ProductCommandHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = builders.NewProductAcceptanceTestBuilder()
		factory = factories.NewProductFactory(builder.DB)
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("Complete product management flow", func() {
		It("مدیر می‌تواند محصول جدید ایجاد کند و سپس آن را به‌روزرسانی و حذف کند", func() {
			// Phase 1: Setup (Arrange)
			category := factory.CreateCategory("Clothing", "clothing")
			createCmd := factory.CreateProductCommand("Men's T-Shirt", "Test Brand", uint64(category.ID))

			// Phase 2: Exercise (Act) - Create product
			err := handler.CreateProductHandler(ctx, createCmd)

			// Phase 3: Verify (Assert) - Verify product creation
			Expect(err).NotTo(HaveOccurred())
			product := helpers.FindProductBySlug(builder.DB, command_handler.GenerateSlug(createCmd.Name))
			Expect(product.Name).To(Equal("Men's T-Shirt"))
			Expect(product.Features).To(HaveLen(2))
			Expect(product.Specs).To(HaveLen(2))
			Expect(product.Details).To(HaveLen(1))

			// Phase 2: Exercise (Act) - Update product
			updateCmd := factory.CreateUpdateCommand(uint64(product.ID), "Updated Men's T-Shirt", "New Brand", uint64(category.ID))
			err = handler.UpdateProductHandler(ctx, updateCmd)

			// Phase 3: Verify (Assert) - Verify product update
			Expect(err).NotTo(HaveOccurred())
			updated := helpers.FindProductByID(builder.DB, uint64(product.ID))
			Expect(updated.Name).To(Equal("Updated Men's T-Shirt"))
			Expect(updated.Brand).To(Equal("New Brand"))
			Expect(updated.Features).To(HaveLen(1))
			Expect(updated.Features[0].Feature).To(Equal("New Feature"))

			// Phase 2: Exercise (Act) - Delete product
			deleteCmd := factory.CreateDeleteCommand(uint64(product.ID), true)
			err = handler.DeleteProductHandler(ctx, deleteCmd)

			// Phase 3: Verify (Assert) - Verify product deletion
			Expect(err).NotTo(HaveOccurred())
			helpers.VerifyProductNotFound(builder.DB, uint64(product.ID))
		})
	})

	Describe("Duplicate slug prevention", func() {
		It("سیستم از ایجاد محصول با شناسه تکراری جلوگیری می‌کند", func() {
			// Phase 1: Setup (Arrange)
			category := factory.CreateCategory("Clothing", "clothing")
			cmd1 := factory.CreateProductCommand("First Product", "Brand", uint64(category.ID))

			// Phase 2: Exercise (Act) - Create first product
			err := handler.CreateProductHandler(ctx, cmd1)
			Expect(err).NotTo(HaveOccurred())

			// Phase 1: Setup (Arrange) - Prepare duplicate command
			cmd2 := factory.CreateProductCommand("First Product", "Other Brand", uint64(category.ID))

			// Phase 2: Exercise (Act) - Try to create duplicate product
			err = handler.CreateProductHandler(ctx, cmd2)

			// Phase 3: Verify (Assert) - Verify conflict error
			Expect(err).To(HaveOccurred())
			Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeConflict))
		})
	})

	Describe("Product validation", func() {
		It("سیستم از ایجاد محصول بدون قیمت جلوگیری می‌کند", func() {
			// Phase 1: Setup (Arrange)
			category := factory.CreateCategory("Clothing", "clothing")
			desc := "Description"
			cmd := &commands.CreateProduct{
				Name:        "Product Without Price",
				Brand:       "Brand",
				Description: &desc,
				CategoryID:  uint64(category.ID),
				Details:     []commands.ProductDetailInput{},
			}

			// Phase 2: Exercise (Act)
			err := handler.CreateProductHandler(ctx, cmd)

			// Phase 3: Verify (Assert)
			Expect(err).To(HaveOccurred())
			Expect(helpers.GetErrorType(err)).To(Equal(apperrors.ErrorTypeValidation))
		})
	})
})
