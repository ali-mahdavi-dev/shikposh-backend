package entity

import (
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type TokenID uint64

type Token struct {
	adapter.BaseEntity
	ID           TokenID `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Token        string         `json:"token" gorm:"token"`
	RefreshToken string         `json:"refresh_token" gorm:"refresh_token"`
	UserID       UserID         `json:"user_id" gorm:"user_id"`
}

func (t *Token) GetID() uint64 {
	return uint64(t.ID)
}

func NewToken(token string, refreshToken string, userID UserID) *Token {
	return &Token{
		Token:        token,
		RefreshToken: refreshToken,
		UserID:       userID,
	}
}
