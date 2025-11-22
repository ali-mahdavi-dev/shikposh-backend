package factories

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

// UserFactory provides factory methods for user acceptance tests
type UserFactory struct {
	db *gorm.DB
}

func NewUserFactory(db *gorm.DB) *UserFactory {
	return &UserFactory{db: db}
}

// CreateRegisterCommand creates a register user command
func (f *UserFactory) CreateRegisterCommand(phone, email, password string) *commands.RegisterUser {
	return &commands.RegisterUser{
		AvatarIdentifier: "avatar123",
		Phone:            phone,
		FirstName:        "John",
		LastName:         "Doe",
		Email:            email,
		Password:         password,
	}
}

// CreateLoginCommand creates a login user command
func (f *UserFactory) CreateLoginCommand(phone, password string) *commands.LoginUser {
	return &commands.LoginUser{
		Phone:    phone,
		Password: password,
	}
}

// CreateLogoutCommand creates a logout command
func (f *UserFactory) CreateLogoutCommand(userID uint64) *commands.Logout {
	return &commands.Logout{
		UserID: userID,
	}
}

// CreateUser creates a user in database
func (f *UserFactory) CreateUser(phone, email, password string) *entity.User {
	userRepo := repository.NewUserRepository(f.db)
	user := &entity.User{
		Phone:     phone,
		FirstName: "John",
		LastName:  "Doe",
		Email:     email,
		Password:  password, // Should be hashed, but for fixture we'll let handler do it
	}
	err := userRepo.Save(context.Background(), user)
	Expect(err).NotTo(HaveOccurred())
	return user
}
