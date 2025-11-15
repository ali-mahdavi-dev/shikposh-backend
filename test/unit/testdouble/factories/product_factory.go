package factories

import (
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/service_layer/command_handler"
)

func CreateProductCommand(name, brand string, categoryID uint64) *commands.CreateProduct {
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

func CreateUpdateProductCommand(id uint64, name, brand string, categoryID uint64) *commands.UpdateProduct {
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

func CreateCategory(id uint64, name, slug string) *entity.Category {
	return &entity.Category{
		ID:   entity.CategoryID(id),
		Name: name,
		Slug: slug,
	}
}

func CreateProduct(id uint64, name, slug, brand string, categoryID uint64) *productaggregate.Product {
	return &productaggregate.Product{
		ID:         productaggregate.ProductID(id),
		Name:       name,
		Slug:       slug,
		Brand:      brand,
		CategoryID: categoryID,
	}
}
