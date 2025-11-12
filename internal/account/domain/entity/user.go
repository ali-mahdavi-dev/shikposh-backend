package entity

import (
	"time"

	"shikposh-backend/internal/account/domain/events"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type UserID uint64

type User struct {
	adapter.BaseEntity
	ID               UserID `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	AvatarIdentifier string         `json:"avatar_identifier" gorm:"avatar_identifier"`
	UserName         string         `json:"user_name" gorm:"user_name"`
	FirstName        string         `json:"first_name" gorm:"first_name"`
	LastName         string         `json:"last_name" gorm:"last_name"`
	Email            string         `json:"email" gorm:"email"`
	Password         string         `json:"password" gorm:"password"`
}

func NewUser(
	avatarIdentifier string,
	userName string,
	firstName string,
	lastName string,
	email string,
	password string,
) *User {
	user := &User{
		AvatarIdentifier: avatarIdentifier,
		UserName:         userName,
		FirstName:        firstName,
		LastName:         lastName,
		Email:            email,
		Password:         password,
	}

	// Add register event with pointer to user.ID so it updates when ID is set
	userID := uint64(user.ID)
	user.AddEvent(&events.RegisterUserEvent{
		UserID:           &userID,
		AvatarIdentifier: user.AvatarIdentifier,
		UserName:         user.UserName,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
	})

	return user
}
