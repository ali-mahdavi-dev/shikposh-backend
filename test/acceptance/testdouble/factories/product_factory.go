package factories

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/service_layer/command_handler"

	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

// ProductFactory provides factory methods for product acceptance tests
type ProductFactory struct {
	db *gorm.DB
}

func NewProductFactory(db *gorm.DB) *ProductFactory {
	return &ProductFactory{db: db}
}

// CreateCategory creates a category
func (f *ProductFactory) CreateCategory(name, slug string) *entity.Category {
	categoryRepo := repository.NewCategoryRepository(f.db)
	category := &entity.Category{
		Name: name,
		Slug: slug,
	}
	err := categoryRepo.Save(context.Background(), category)
	Expect(err).NotTo(HaveOccurred())
	return category
}

// CreateProduct creates a product with all associations
func (f *ProductFactory) CreateProduct(name, brand string, categoryID uint64) *productaggregate.Product {
	desc := "Product description"
	cmd := &commands.CreateProduct{
		Name:        name,
		Brand:       brand,
		Description: &desc,
		CategoryID:  categoryID,
		Tags:        []string{"tag1"},
		Sizes:       []string{"M", "L"},
		Details: []commands.ProductDetailInput{
			{Price: 100000.0, Stock: 10},
		},
	}

	productRepo := repository.NewProductRepository(f.db)
	product := &productaggregate.Product{
		Name:        cmd.Name,
		Slug:        command_handler.GenerateSlug(cmd.Name),
		Brand:       cmd.Brand,
		Description: cmd.Description,
		CategoryID:  categoryID,
		Tags:        cmd.Tags,
		Sizes:       cmd.Sizes,
	}
	err := productRepo.Save(context.Background(), product)
	Expect(err).NotTo(HaveOccurred())

	// Add details
	for _, detailInput := range cmd.Details {
		detail := &productaggregate.ProductDetail{
			ProductID: product.ID,
			Price:    detailInput.Price,
			Stock:    detailInput.Stock,
		}
		f.db.Create(detail)
	}

	return product
}

// CreateProductCommand creates a create product command
func (f *ProductFactory) CreateProductCommand(name, brand string, categoryID uint64) *commands.CreateProduct {
	desc := "Product description"
	return &commands.CreateProduct{
		Name:        name,
		Brand:       brand,
		Description: &desc,
		CategoryID:  categoryID,
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
}

// CreateUpdateCommand creates an update product command
func (f *ProductFactory) CreateUpdateCommand(productID uint64, name, brand string, categoryID uint64) *commands.UpdateProduct {
	desc := "Updated description"
	return &commands.UpdateProduct{
		ID:          productID,
		Name:        name,
		Slug:        command_handler.GenerateSlug(name),
		Brand:       brand,
		Description: &desc,
		CategoryID:  categoryID,
		Features: []commands.ProductFeatureInput{
			{Feature: "New Feature", Order: 1},
		},
		Details: []commands.ProductDetailInput{
			{Price: 150000.0, Stock: 20},
		},
	}
}

// CreateDeleteCommand creates a delete product command
func (f *ProductFactory) CreateDeleteCommand(productID uint64, softDelete bool) *commands.DeleteProduct {
	return &commands.DeleteProduct{
		ID:         productID,
		SoftDelete: softDelete,
	}
}

