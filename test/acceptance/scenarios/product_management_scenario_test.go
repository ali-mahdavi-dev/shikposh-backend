package acceptance_test

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/service_layer/command_handler"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("Product Management Acceptance Scenarios", func() {
	var (
		builder *ProductAcceptanceTestBuilder
		handler *command_handler.ProductCommandHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = NewProductAcceptanceTestBuilder(nil)
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("Complete product management flow", func() {
		It("مدیر می‌تواند محصول جدید ایجاد کند و سپس آن را به‌روزرسانی و حذف کند", func() {
			// Step 1: Create category
			categoryRepo := repository.NewCategoryRepository(builder.db)
			category := &entity.Category{
				Name: "Clothing",
				Slug: "clothing",
			}
			err := categoryRepo.Save(ctx, category)
			Expect(err).NotTo(HaveOccurred())

			// Step 2: Create product
			desc := "Test product description"
			createCmd := &commands.CreateProduct{
				Name:        "Men's T-Shirt",
				Brand:       "Test Brand",
				Description: &desc,
				CategoryID:  uint64(category.ID),
				Tags:        []string{"men", "summer"},
				Sizes:       []string{"M", "L", "XL"},
				Features: []commands.ProductFeatureInput{
					{Feature: "Waterproof", Order: 1},
					{Feature: "Washable", Order: 2},
				},
				Specs: []commands.ProductSpecInput{
					{Key: "Material", Value: "100% Cotton", Order: 1},
					{Key: "Country", Value: "USA", Order: 2},
				},
				Details: []commands.ProductDetailInput{
					{Price: 100000.0, Stock: 10},
				},
			}

			err = handler.CreateProductHandler(ctx, createCmd)
			Expect(err).NotTo(HaveOccurred())

			// Step 3: Verify product was created with all associations
			productRepo := repository.NewProductRepository(builder.db)
			product, err := productRepo.FindBySlug(ctx, createCmd.Slug)
			Expect(err).NotTo(HaveOccurred())
			Expect(product).NotTo(BeNil())
			Expect(product.Name).To(Equal("Men's T-Shirt"))
			Expect(product.Features).To(HaveLen(2))
			Expect(product.Specs).To(HaveLen(2))
			Expect(product.Details).To(HaveLen(1))

			// Step 4: Update product
			newDesc := "Updated description"
			updateCmd := &commands.UpdateProduct{
				ID:          uint64(product.ID),
				Name:        "Updated Men's T-Shirt",
				Slug:        "updated-mens-tshirt",
				Brand:       "New Brand",
				Description: &newDesc,
				CategoryID:  uint64(category.ID),
				Features: []commands.ProductFeatureInput{
					{Feature: "New Feature", Order: 1},
				},
				Details: []commands.ProductDetailInput{
					{Price: 150000.0, Stock: 20},
				},
			}

			err = handler.UpdateProductHandler(ctx, updateCmd)
			Expect(err).NotTo(HaveOccurred())

			// Step 5: Verify product was updated
			updated, err := productRepo.FindByID(ctx, uint64(product.ID))
			Expect(err).NotTo(HaveOccurred())
			Expect(updated.Name).To(Equal("Updated Men's T-Shirt"))
			Expect(updated.Brand).To(Equal("New Brand"))
			Expect(updated.Features).To(HaveLen(1))
			Expect(updated.Features[0].Feature).To(Equal("New Feature"))

			// Step 6: Soft delete product
			deleteCmd := &commands.DeleteProduct{
				ID:         uint64(product.ID),
				SoftDelete: true,
			}

			err = handler.DeleteProductHandler(ctx, deleteCmd)
			Expect(err).NotTo(HaveOccurred())

			// Step 7: Verify product was soft deleted
			_, err = productRepo.FindByID(ctx, uint64(product.ID))
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(appadapter.ErrEntityNotFound))
		})
	})

	Describe("Duplicate slug prevention", func() {
		It("سیستم از ایجاد محصول با شناسه تکراری جلوگیری می‌کند", func() {
			// Step 1: Create category
			categoryRepo := repository.NewCategoryRepository(builder.db)
			category := &entity.Category{
				Name: "Clothing",
				Slug: "clothing",
			}
			err := categoryRepo.Save(ctx, category)
			Expect(err).NotTo(HaveOccurred())

			// Step 2: Create first product
			desc := "Description"
			cmd1 := &commands.CreateProduct{
				Name:        "First Product",
				Brand:       "Brand",
				Description: &desc,
				CategoryID:  uint64(category.ID),
				Details: []commands.ProductDetailInput{
					{Price: 100000.0},
				},
			}
			err = handler.CreateProductHandler(ctx, cmd1)
			Expect(err).NotTo(HaveOccurred())

			// Step 3: Try to create product with same slug (generated from same name)
			cmd2 := &commands.CreateProduct{
				Name:        "First Product",
				Brand:       "Other Brand",
				Description: &desc,
				CategoryID:  uint64(category.ID),
				Details: []commands.ProductDetailInput{
					{Price: 200000.0},
				},
			}
			err = handler.CreateProductHandler(ctx, cmd2)
			Expect(err).To(HaveOccurred())
			appErr, ok := err.(apperrors.Error)
			Expect(ok).To(BeTrue())
			Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeConflict))
		})
	})

	Describe("Product validation", func() {
		It("سیستم از ایجاد محصول بدون قیمت جلوگیری می‌کند", func() {
			// Step 1: Create category
			categoryRepo := repository.NewCategoryRepository(builder.db)
			category := &entity.Category{
				Name: "Clothing",
				Slug: "clothing",
			}
			err := categoryRepo.Save(ctx, category)
			Expect(err).NotTo(HaveOccurred())

			// Step 2: Try to create product without price details
			desc := "Description"
			cmd := &commands.CreateProduct{
				Name:        "Product Without Price",
				Brand:       "Brand",
				Description: &desc,
				CategoryID:  uint64(category.ID),
				Details:     []commands.ProductDetailInput{}, // No price details
			}

			err = handler.CreateProductHandler(ctx, cmd)
			Expect(err).To(HaveOccurred())
			appErr, ok := err.(apperrors.Error)
			Expect(ok).To(BeTrue())
			Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeValidation))
		})
	})
})

// ProductAcceptanceTestBuilder helps build acceptance test scenarios for products
type ProductAcceptanceTestBuilder struct {
	db  *gorm.DB
	uow unit_of_work.PGUnitOfWork
}

func NewProductAcceptanceTestBuilder(t GinkgoTInterface) *ProductAcceptanceTestBuilder {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	err = db.AutoMigrate(
		&entity.Category{},
		&productaggregate.Product{},
		&productaggregate.ProductFeature{},
		&productaggregate.ProductDetail{},
		&productaggregate.ProductSpec{},
	)
	Expect(err).NotTo(HaveOccurred())

	eventCh := make(chan appadapter.EventWithWaitGroup, 100)
	uow := unit_of_work.New(db, eventCh)

	return &ProductAcceptanceTestBuilder{
		db:  db,
		uow: uow,
	}
}

func (b *ProductAcceptanceTestBuilder) BuildHandler() *command_handler.ProductCommandHandler {
	return command_handler.NewProductCommandHandler(b.uow)
}

func (b *ProductAcceptanceTestBuilder) Cleanup() {
	b.db.Exec("DELETE FROM products")
	b.db.Exec("DELETE FROM categories")
	b.db.Exec("DELETE FROM product_features")
	b.db.Exec("DELETE FROM product_details")
	b.db.Exec("DELETE FROM product_specs")
}
