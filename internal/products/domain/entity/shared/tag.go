package shared

import (
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type TagID uint64

// Tag represents a tag that can be associated with products
type Tag struct {
	adapter.BaseEntity
	ID        TagID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" gorm:"name;uniqueIndex;not null"`
	Slug      string         `json:"slug" gorm:"slug;uniqueIndex;not null"`
}

func (t *Tag) TableName() string {
	return "tags"
}
