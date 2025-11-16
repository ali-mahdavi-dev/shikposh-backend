package helpers

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	appadapter "github.com/ali-mahdavi-dev/framework/adapter"

	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

// FindProductBySlug finds a product by slug
func FindProductBySlug(db *gorm.DB, slug string) *productaggregate.Product {
	productRepo := repository.NewProductRepository(db)
	product, err := productRepo.FindBySlug(context.Background(), slug)
	Expect(err).NotTo(HaveOccurred())
	return product
}

// FindProductByID finds a product by ID
func FindProductByID(db *gorm.DB, productID uint64) *productaggregate.Product {
	productRepo := repository.NewProductRepository(db)
	product, err := productRepo.FindByID(context.Background(), productID)
	Expect(err).NotTo(HaveOccurred())
	return product
}

// FindProductByIDWithError finds a product by ID and returns error if not found
func FindProductByIDWithError(db *gorm.DB, productID uint64) (*productaggregate.Product, error) {
	productRepo := repository.NewProductRepository(db)
	return productRepo.FindByID(context.Background(), productID)
}


// VerifyProductNotFound verifies that product was not found
func VerifyProductNotFound(db *gorm.DB, productID uint64) {
	_, err := FindProductByIDWithError(db, productID)
	Expect(err).To(HaveOccurred())
	Expect(err).To(Equal(appadapter.ErrEntityNotFound))
}

