package product_aggregate

import (
	"time"

	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type ProductFeatureID uint64

// ProductFeature is an Aggregate Entity within the Product Aggregate.
// It should only be accessed through the Product aggregate root.
type ProductFeature struct {
	adapter.BaseEntity
	ID        ProductFeatureID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ProductID ProductID      `json:"product_id" gorm:"product_id"`
	Feature   string         `json:"feature" gorm:"feature"`       // e.g., "ضد آب", "قابل شستشو"
	Order     int            `json:"order" gorm:"order;default:0"` // For ordering features
}

func (pf *ProductFeature) TableName() string {
	return "product_features"
}

// NewProductFeature creates a new ProductFeature instance
func NewProductFeature(productID ProductID, feature string, order int) ProductFeature {
	return ProductFeature{
		ProductID: productID,
		Feature:   feature,
		Order:     order,
	}
}
