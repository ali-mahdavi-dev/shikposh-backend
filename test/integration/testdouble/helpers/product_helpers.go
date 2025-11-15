package helpers

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func FindProductBySlug(db *gorm.DB, slug string) *productaggregate.Product {
	productRepo := repository.NewProductRepository(db)
	product, err := productRepo.FindBySlug(context.Background(), slug)
	Expect(err).NotTo(HaveOccurred())
	return product
}

func FindProductByID(db *gorm.DB, productID uint64) *productaggregate.Product {
	productRepo := repository.NewProductRepository(db)
	product, err := productRepo.FindByID(context.Background(), productID)
	Expect(err).NotTo(HaveOccurred())
	return product
}

func FindProductByIDWithError(db *gorm.DB, productID uint64) (*productaggregate.Product, error) {
	productRepo := repository.NewProductRepository(db)
	return productRepo.FindByID(context.Background(), productID)
}
