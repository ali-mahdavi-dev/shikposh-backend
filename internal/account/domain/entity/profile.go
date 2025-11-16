package entity

import (
	"time"

	"github.com/shikposh/framework/adapter"

	"gorm.io/gorm"
)

type ProfileID uint64

type Profile struct {
	adapter.BaseEntity
	ID        ProfileID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    UserID         `json:"user_id" gorm:"user_id;uniqueIndex"`
	Bio       string         `json:"bio" gorm:"bio"`
	Phone     string         `json:"phone" gorm:"phone"`
	Address   string         `json:"address" gorm:"address"`
}

func NewProfile(userID UserID) *Profile {
	return &Profile{
		UserID: userID,
	}
}
