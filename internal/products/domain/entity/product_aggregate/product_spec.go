package product_aggregate

import (
	"time"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type ProductSpecID uint64

// ProductSpec is an Aggregate Entity within the Product Aggregate.
// It should only be accessed through the Product aggregate root.
type ProductSpec struct {
	adapter.BaseEntity
	ID        ProductSpecID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ProductID ProductID      `json:"product_id" gorm:"product_id"`
	Key       string         `json:"key" gorm:"key"`               // e.g., "جنس", "کشور سازنده"
	Value     string         `json:"value" gorm:"value"`           // e.g., "پنبه 100%", "ایران"
	Order     int            `json:"order" gorm:"order;default:0"` // For ordering specs
}

func (ps *ProductSpec) TableName() string {
	return "product_specs"
}

// NewProductSpec creates a new ProductSpec instance using command root input
func NewProductSpec(productID ProductID, input commands.ProductSpecInput) ProductSpec {
	return ProductSpec{
		ProductID: productID,
		Key:       input.Key,
		Value:     input.Value,
		Order:     input.Order,
	}
}
