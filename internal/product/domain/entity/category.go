package entity

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type CategoryID uint64

// MarshalJSON converts CategoryID to string in JSON
func (id CategoryID) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(id), 10))
}

// UnmarshalJSON converts string to CategoryID from JSON
func (id *CategoryID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		// Try to unmarshal as number
		var n uint64
		if err2 := json.Unmarshal(data, &n); err2 != nil {
			return err
		}
		*id = CategoryID(n)
		return nil
	}
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*id = CategoryID(val)
	return nil
}

type Category struct {
	adapter.BaseEntity
	ID           CategoryID     `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `json:"name" gorm:"name"`
	Slug         string         `json:"slug" gorm:"slug;uniqueIndex"`
	Description  *string        `json:"description,omitempty" gorm:"description;type:text"`
	Image        *string        `json:"image,omitempty" gorm:"image"`
	ParentID     *CategoryID    `json:"parent_id,omitempty" gorm:"parent_id"`
	Parent       *Category      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	ProductCount int            `json:"product_count,omitempty" gorm:"-"`
}

func (c *Category) GetID() uint64 {
	return uint64(c.ID)
}

func (c *Category) TableName() string {
	return "categories"
}
