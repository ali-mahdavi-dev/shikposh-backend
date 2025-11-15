package factories

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/service_layer/command_handler"

	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name, slug string) *entity.Category {
	category := &entity.Category{
		Name: name,
		Slug: slug,
	}
	categoryRepo := repository.NewCategoryRepository(db)
	err := categoryRepo.Save(context.Background(), category)
	Expect(err).NotTo(HaveOccurred())
	return category
}

func CreateProductCommand(name, brand string, categoryID uint64) *commands.CreateProduct {
	description := "Product description"
	return &commands.CreateProduct{
		Name:        name,
		Slug:        "",
		Brand:       brand,
		Description: &description,
		CategoryID:  categoryID,
		Tags:        []string{"tag1"},
		Sizes:       []string{"M", "L"},
		Details: []commands.ProductDetailInput{
			{Price: 100000.0, Stock: 10},
		},
	}
}

func CreateProductCommandWithEmptyName(categoryID uint64) *commands.CreateProduct {
	description := "Description"
	return &commands.CreateProduct{
		Name:        "",
		Slug:        "",
		Brand:       "Brand",
		Description: &description,
		CategoryID:  categoryID,
		Tags:        []string{"tag1"},
		Sizes:       []string{"M"},
		Details: []commands.ProductDetailInput{
			{Price: 100000.0, Stock: 10},
		},
	}
}

func CreateProductCommandWithoutDetails(name, brand string, categoryID uint64) *commands.CreateProduct {
	description := "Description"
	return &commands.CreateProduct{
		Name:        name,
		Slug:        "",
		Brand:       brand,
		Description: &description,
		CategoryID:  categoryID,
		Tags:        []string{"tag1"},
		Sizes:       []string{"M"},
		Details:     []commands.ProductDetailInput{},
	}
}

func CreateUpdateCommand(productID uint64, name, brand string, categoryID uint64) *commands.UpdateProduct {
	description := "New description"
	return &commands.UpdateProduct{
		ID:          productID,
		Name:        name,
		Slug:        command_handler.GenerateSlug(name),
		Brand:       brand,
		Description: &description,
		CategoryID:  categoryID,
		Features: []commands.ProductFeatureInput{
			{Feature: "New Feature", Order: 1},
		},
		Details: []commands.ProductDetailInput{
			{Price: 150000.0},
		},
	}
}

func CreateUpdateCommandWithDuplicateSlug(productID uint64, duplicateSlug string, categoryID uint64) *commands.UpdateProduct {
	description := "Description"
	return &commands.UpdateProduct{
		ID:          productID,
		Name:        "Product Two",
		Slug:        duplicateSlug,
		Brand:       "Brand2",
		Description: &description,
		CategoryID:  categoryID,
	}
}
