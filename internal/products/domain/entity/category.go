package entity

import (
	"time"

	"shikposh-backend/internal/products/domain/types"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type Category struct {
	adapter.BaseEntity
	ID           types.CategoryID `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string          `json:"name" gorm:"name"`
	Slug         string          `json:"slug" gorm:"slug;uniqueIndex"`
	Description  *string         `json:"description,omitempty" gorm:"description;type:text"`
	Image        *string         `json:"image,omitempty" gorm:"image"`
	ParentID     *types.CategoryID `json:"parent_id,omitempty" gorm:"parent_id"`
	Parent       *Category       `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	ProductCount int       `json:"product_count,omitempty" gorm:"-"`
}

func (c *Category) TableName() string {
	return "categories"
}
