package product_aggregate

import (
	"time"
)

type ProductImageID uint64

// ProductImage - عکس‌های محصول بر اساس رنگ
// images: { "colorId": ["url1", "url2"] }
type ProductImage struct {
	ID        ProductImageID `gorm:"primaryKey"`
	ProductID ProductID      `gorm:"not null"`
	ColorID   ColorID        `gorm:"not null"`
	URL       string         `gorm:"not null"`
	SortOrder int            `gorm:"default:0"`
	CreatedAt time.Time
}

func (ProductImage) TableName() string {
	return "product_images"
}
