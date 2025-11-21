package shared

import (
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type SizeID uint64

// Size represents a size that can be associated with products
type Size struct {
	adapter.BaseEntity
	ID        SizeID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" gorm:"name;uniqueIndex;not null"`
	Slug      string         `json:"slug" gorm:"slug;uniqueIndex;not null"`
}

func (s *Size) TableName() string {
	return "sizes"
}
