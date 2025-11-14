package product_aggregate

import (
	"time"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity/shared"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type ProductDetailID uint64

// ProductDetail is an Aggregate Entity within the Product Aggregate.
// It should only be accessed through the Product aggregate root.
type ProductDetail struct {
	adapter.BaseEntity
	ID            ProductDetailID `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt      `gorm:"index"`
	ProductID     ProductID           `json:"product_id" gorm:"product_id"`
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

// NewProductDetail creates a new ProductDetail instance using command root input
func NewProductDetail(productID ProductID, input commands.ProductDetailInput) ProductDetail {
	return ProductDetail{
		ProductID:     productID,
		ColorKey:      input.ColorKey,
		ColorName:     input.ColorName,
		SizeKey:       input.SizeKey,
		Price:         input.Price,
		OriginalPrice: input.OriginalPrice,
		Stock:         input.Stock,
		Discount:      input.Discount,
		Images:        []shared.Attachment{},
	}
}
