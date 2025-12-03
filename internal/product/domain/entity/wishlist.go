package entity

import (
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
)

type WishlistID uint64

type Wishlist struct {
	adapter.BaseEntity
	ID        WishlistID `gorm:"primaryKey"`
	UserID    uint64     `json:"user_id" gorm:"column:user_id;index"`
	ProductID uint64     `json:"product_id" gorm:"column:product_id;index"`
	CreatedAt time.Time
}

func (w *Wishlist) GetID() uint64 {
	return uint64(w.ID)
}

func (Wishlist) TableName() string {
	return "wishlists"
}

func NewWishlist(userID uint64, productID uint64) *Wishlist {
	return &Wishlist{
		UserID:    userID,
		ProductID: productID,
	}
}
