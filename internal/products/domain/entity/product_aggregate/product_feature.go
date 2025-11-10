package product_aggregate

import (
	"shikposh-backend/pkg/framework/adapter"
)

// ProductFeature is an Aggregate Entity within the Product Aggregate.
// It should only be accessed through the Product aggregate root.
type ProductFeature struct {
	adapter.BaseEntity
	ProductID uint64 `json:"product_id" gorm:"product_id"`
	Feature   string `json:"feature" gorm:"feature"`       // e.g., "ضد آب", "قابل شستشو"
	Order     int    `json:"order" gorm:"order;default:0"` // For ordering features
}

func (pf *ProductFeature) TableName() string {
	return "product_features"
}

// NewProductFeature creates a new ProductFeature instance
func NewProductFeature(productID uint64, feature string, order int) ProductFeature {
	return ProductFeature{
		ProductID: productID,
		Feature:   feature,
		Order:     order,
	}
}
