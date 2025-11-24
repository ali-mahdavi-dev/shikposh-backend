package product_aggregate

import (
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type ColorID uint64

// Color represents a color that can be associated with products
type Color struct {
	adapter.BaseEntity
	ID        ColorID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" gorm:"name;uniqueIndex;not null"`
	Slug      string         `json:"slug" gorm:"slug;uniqueIndex;not null"`
	Hex       string         `json:"hex" gorm:"hex;not null"` // Hex color code (e.g., "#FF0000")
}

func (c *Color) TableName() string {
	return "colors"
}
