package entity

import (
	"time"

	accountentity "shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type ReviewID uint64

type Review struct {
	adapter.BaseEntity
	ID         ReviewID `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt              `gorm:"index"`
	ProductID  product_aggregate.ProductID `json:"product_id" gorm:"product_id"`
	Product    *product_aggregate.Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	UserID     accountentity.UserID        `json:"user_id" gorm:"user_id"`
	UserName   string                      `json:"user_name" gorm:"user_name"`
	UserAvatar *string                     `json:"user_avatar,omitempty" gorm:"user_avatar"`
	Rating     int                         `json:"rating" gorm:"rating"`
	Comment    string                      `json:"comment" gorm:"comment;type:text"`
	Helpful    int                         `json:"helpful" gorm:"helpful;default:0"`
	NotHelpful int                         `json:"not_helpful" gorm:"not_helpful;default:0"`
	Verified   bool                        `json:"verified" gorm:"verified;default:false"`
}

func (r *Review) TableName() string {
	return "reviews"
}

// NewReview creates a new Review instance using a command
func NewReview(cmd *commands.CreateReview) *Review {
	return &Review{
		ProductID:  product_aggregate.ProductID(cmd.ProductID),
		UserID:     accountentity.UserID(cmd.UserID),
		UserName:   cmd.UserName,
		Rating:     cmd.Rating,
		Comment:    cmd.Comment,
		Helpful:    0,
		NotHelpful: 0,
		Verified:   false,
	}
}
