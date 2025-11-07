package product_aggregate

import (
	"shikposh-backend/internal/products/domain/entity/shared"
	"shikposh-backend/pkg/framework/adapter"
)

// ProductDetail is an Aggregate Entity within the Product Aggregate.
// It should only be accessed through the Product aggregate root.
type ProductDetail struct {
	adapter.BaseEntity
	ProductID     uint64              `json:"product_id" gorm:"product_id"`
	ColorKey      *string             `json:"color_key,omitempty" gorm:"color_key"`           // e.g., "red", "blue" - nullable for size-only details
	ColorName     *string             `json:"color_name,omitempty" gorm:"color_name"`         // e.g., "قرمز", "آبی" - nullable
	SizeKey       *string             `json:"size_key,omitempty" gorm:"size_key"`             // e.g., "M", "L" - nullable for color-only details
	Price         float64             `json:"price" gorm:"price"`                             // Price for this variant
	OriginalPrice *float64            `json:"original_price,omitempty" gorm:"original_price"` // Original price before discount
	Stock         int                 `json:"stock" gorm:"stock;default:0"`
	Discount      int                 `json:"discount" gorm:"discount;default:0"`
	Images        []shared.Attachment `json:"-" gorm:"polymorphic:Attachable;polymorphicValue:ProductDetail"` // Polymorphic relationship
}

func (pd *ProductDetail) TableName() string {
	return "product_details"
}
