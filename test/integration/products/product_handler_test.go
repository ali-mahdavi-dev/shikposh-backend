package integration_test

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

var _ = Describe("ProductCommandHandler Integration", func() {
	var (
		builder *ProductIntegrationTestBuilder
		handler *command_handler.ProductCommandHandler
		ctx     context.Context
	)

	BeforeEach(func() {
		builder = NewProductIntegrationTestBuilder(nil)
		handler = builder.BuildHandler()
		ctx = context.Background()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("CreateProductHandler", func() {
		Context("when creating a new product", func() {
			It("should create product with all associations in database", func() {
				category := createTestCategory(nil, builder.db, "Clothing", "clothing")
				cmd := createTestProductCommand("Men's T-Shirt", "Test Brand", uint64(category.ID))

				cmd.Features = []commands.ProductFeatureInput{
					{Feature: "Waterproof", Order: 1},
					{Feature: "Washable", Order: 2},
				}
				cmd.Specs = []commands.ProductSpecInput{
					{Key: "Material", Value: "100% Cotton", Order: 1},
				}

				err := handler.CreateProductHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())

				// Verify product was persisted with associations
				productRepo := repository.NewProductRepository(builder.db)
				product, err := productRepo.FindBySlug(ctx, cmd.Slug)
				Expect(err).NotTo(HaveOccurred())
				Expect(product).NotTo(BeNil())
				Expect(product.Name).To(Equal("Men's T-Shirt"))
				Expect(product.Features).To(HaveLen(2))
				Expect(product.Details).To(HaveLen(1))
				Expect(product.Specs).To(HaveLen(1))
			})
		})

		Context("when slug already exists", func() {
			It("should enforce unique slug constraint", func() {
				category := createTestCategory(nil, builder.db, "Clothing", "clothing")

				// Create first product
				cmd1 := createTestProductCommand("First Product", "Brand", uint64(category.ID))
				err := handler.CreateProductHandler(ctx, cmd1)
				Expect(err).NotTo(HaveOccurred())

				// Try to create product with same slug
				cmd2 := createTestProductCommand("First Product", "Other Brand", uint64(category.ID))
				err = handler.CreateProductHandler(ctx, cmd2)
				Expect(err).To(HaveOccurred())
				appErr, ok := err.(apperrors.Error)
				Expect(ok).To(BeTrue())
				Expect(appErr.Type()).To(Equal(apperrors.ErrorTypeConflict))
			})
		})
	})

	Describe("UpdateProductHandler", func() {
		Context("when updating an existing product", func() {
			It("should update product and replace associations", func() {
				category := createTestCategory(nil, builder.db, "Clothing", "clothing")

				// Create product
				createCmd := createTestProductCommand("Old Product", "Old Brand", uint64(category.ID))
				createCmd.Features = []commands.ProductFeatureInput{
					{Feature: "Old Feature", Order: 1},
				}
				err := handler.CreateProductHandler(ctx, createCmd)
				Expect(err).NotTo(HaveOccurred())

				productRepo := repository.NewProductRepository(builder.db)
				product, err := productRepo.FindBySlug(ctx, createCmd.Slug)
				Expect(err).NotTo(HaveOccurred())

				// Update product
				desc := "New description"
				updateCmd := &commands.UpdateProduct{
					ID:          uint64(product.ID),
					Name:        "New Product",
					Slug:        "new-product",
					Brand:       "New Brand",
					Description: &desc,
					CategoryID:  uint64(category.ID),
					Features: []commands.ProductFeatureInput{
						{Feature: "New Feature", Order: 1},
					},
					Details: []commands.ProductDetailInput{
						{Price: 150000.0},
					},
				}

				err = handler.UpdateProductHandler(ctx, updateCmd)
				Expect(err).NotTo(HaveOccurred())

				// Verify update
				updated, err := productRepo.FindByID(ctx, uint64(product.ID))
				Expect(err).NotTo(HaveOccurred())
				Expect(updated.Name).To(Equal("New Product"))
				Expect(updated.Brand).To(Equal("New Brand"))
				Expect(updated.Features).To(HaveLen(1))
				Expect(updated.Features[0].Feature).To(Equal("New Feature"))
			})
		})
	})

	Describe("DeleteProductHandler", func() {
		Context("when soft deleting a product", func() {
			It("should soft delete product and keep associations", func() {
				category := createTestCategory(nil, builder.db, "Clothing", "clothing")
				cmd := createTestProductCommand("Product To Delete", "Brand", uint64(category.ID))
				err := handler.CreateProductHandler(ctx, cmd)
				Expect(err).NotTo(HaveOccurred())

				productRepo := repository.NewProductRepository(builder.db)
				product, err := productRepo.FindBySlug(ctx, cmd.Slug)
				Expect(err).NotTo(HaveOccurred())

				deleteCmd := &commands.DeleteProduct{
					ID:         uint64(product.ID),
					SoftDelete: true,
				}

				err = handler.DeleteProductHandler(ctx, deleteCmd)
				Expect(err).NotTo(HaveOccurred())

				// Verify soft delete (product should not be found)
				_, err = productRepo.FindByID(ctx, uint64(product.ID))
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(appadapter.ErrEntityNotFound))
			})
		})
	})
})

// ProductIntegrationTestBuilder helps build integration test scenarios
type ProductIntegrationTestBuilder struct {
	db  *gorm.DB
	uow unit_of_work.PGUnitOfWork
}

func NewProductIntegrationTestBuilder(t GinkgoTInterface) *ProductIntegrationTestBuilder {
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

	return &ProductIntegrationTestBuilder{
		db:  db,
		uow: uow,
	}
}

func (b *ProductIntegrationTestBuilder) BuildHandler() *command_handler.ProductCommandHandler {
	return command_handler.NewProductCommandHandler(b.uow)
}

func (b *ProductIntegrationTestBuilder) Cleanup() {
	b.db.Exec("DELETE FROM products")
	b.db.Exec("DELETE FROM categories")
	b.db.Exec("DELETE FROM product_features")
	b.db.Exec("DELETE FROM product_details")
	b.db.Exec("DELETE FROM product_specs")
}

// Helper functions
func createTestCategory(t GinkgoTInterface, db *gorm.DB, name, slug string) *entity.Category {
	category := &entity.Category{
		Name: name,
		Slug: slug,
	}
	categoryRepo := repository.NewCategoryRepository(db)
	err := categoryRepo.Save(context.Background(), category)
	Expect(err).NotTo(HaveOccurred())
	return category
}

func createTestProductCommand(name, brand string, categoryID uint64) *commands.CreateProduct {
	desc := "Product description"
	return &commands.CreateProduct{
		Name:        name,
		Slug:        "",
		Brand:       brand,
		Description: &desc,
		CategoryID:  categoryID,
		Tags:        []string{"tag1"},
		Sizes:       []string{"M", "L"},
		Details: []commands.ProductDetailInput{
			{Price: 100000.0, Stock: 10},
		},
	}
}
