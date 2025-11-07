package entity

import (
	"shikposh-backend/pkg/framework/adapter"
)

type Review struct {
	adapter.BaseEntity
	ProductID  uint64  `json:"product_id" gorm:"product_id"`
	Product    *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	UserID     uint64  `json:"user_id" gorm:"user_id"`
	UserName   string  `json:"user_name" gorm:"user_name"`
	UserAvatar *string `json:"user_avatar,omitempty" gorm:"user_avatar"`
	Rating     int     `json:"rating" gorm:"rating"`
	Comment    string  `json:"comment" gorm:"comment;type:text"`
	Helpful    int     `json:"helpful" gorm:"helpful;default:0"`
	NotHelpful int     `json:"not_helpful" gorm:"not_helpful;default:0"`
	Verified   bool    `json:"verified" gorm:"verified;default:false"`
}

func (r *Review) TableName() string {
	return "reviews"
}

