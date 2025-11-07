package entity

import (
	"shikposh-backend/pkg/framework/adapter"
)

type Category struct {
	adapter.BaseEntity
	Name        string  `json:"name" gorm:"name"`
	Slug        string  `json:"slug" gorm:"slug;uniqueIndex"`
	Description *string `json:"description,omitempty" gorm:"description;type:text"`
	Image       *string `json:"image,omitempty" gorm:"image"`
	ParentID    *uint64 `json:"parent_id,omitempty" gorm:"parent_id"`
	Parent      *Category `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	ProductCount int     `json:"product_count,omitempty" gorm:"-"`
}

func (c *Category) TableName() string {
	return "categories"
}

