package entity

import (
	"time"

	"github.com/shikposh/framework/adapter"

	"gorm.io/gorm"
)

type TokenID uint64

type Token struct {
	adapter.BaseEntity
	ID        TokenID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Token     string         `json:"token" gorm:"token"`
	UserID    UserID         `json:"user_id" gorm:"user_id"`
}

func NewToken(token string, userID UserID) *Token {
	return &Token{
		Token:  token,
		UserID: userID,
	}
}
