package product_aggregate

import (
	"time"

	"gorm.io/gorm"
)

type ProductVariantID uint64

// ProductVariant - ترکیب رنگ + سایز با موجودی
// variant: { "colorId": { "sizeId": { "stock": N } } }
type ProductVariant struct {
	ID        ProductVariantID `gorm:"primaryKey"`
	ProductID ProductID        `gorm:"not null"`
	ColorID   ColorID          `gorm:"not null"`
	SizeID    SizeID           `gorm:"not null"`
	Stock     int              `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductVariant) TableName() string {
	return "product_variants"
}
