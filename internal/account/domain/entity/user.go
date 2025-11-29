package entity

import (
	"time"

	"shikposh-backend/internal/account/domain/events"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type UserID uint64

type User struct {
	adapter.BaseEntity
	ID               UserID `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	AvatarIdentifier string         `json:"avatar_identifier" gorm:"column:avatar_identifier"`
	FirstName        string         `json:"first_name" gorm:"column:first_name"`
	LastName         string         `json:"last_name" gorm:"column:last_name"`
	Email            string         `json:"email" gorm:"column:email"`
	Phone            string         `json:"phone" gorm:"column:phone;uniqueIndex"`
	Password         string         `json:"password" gorm:"column:password"`
	IsSuperuser      bool           `json:"is_superuser" gorm:"column:is_superuser;default:false"`
	IsAdmin          bool           `json:"is_admin" gorm:"column:is_admin;default:false"`
}

func NewUser(
	avatarIdentifier string,
	firstName string,
	lastName string,
	email string,
	phone string,
	password string,
) *User {
	user := &User{
		AvatarIdentifier: avatarIdentifier,
		FirstName:        firstName,
		LastName:         lastName,
		Email:            email,
		Phone:            phone,
		Password:         password,
	}

	// Add register event with pointer to user.ID so it updates when ID is set
	userID := uint64(user.ID)
	user.AddEvent(&events.RegisterUserEvent{
		UserID:           &userID,
		AvatarIdentifier: user.AvatarIdentifier,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
	})

	return user
}
