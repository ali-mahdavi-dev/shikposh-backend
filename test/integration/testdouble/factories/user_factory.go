package factories

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, email, password string) *entity.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	Expect(err).NotTo(HaveOccurred())

	user := &entity.User{
		UserName:  username,
		FirstName: "Test",
		LastName:  "User",
		Email:     email,
		Password:  string(hashedPassword),
	}

	userRepo := repository.NewUserRepository(db)
	err = userRepo.Save(context.Background(), user)
	Expect(err).NotTo(HaveOccurred())
	return user
}

func CreateRegisterCommand(username, email, password string) *commands.RegisterUser {
	return &commands.RegisterUser{
		AvatarIdentifier: "avatar123",
		UserName:         username,
		FirstName:        "John",
		LastName:         "Doe",
		Email:            email,
		Password:         password,
	}
}

func CreateLoginCommand(username, password string) *commands.LoginUser {
	return &commands.LoginUser{
		UserName: username,
		Password: password,
	}
}

func CreateToken(db *gorm.DB, userID entity.UserID, tokenValue string) {
	tokenRepo := repository.NewTokenRepository(db)
	token := entity.NewToken(tokenValue, userID)
	err := tokenRepo.Save(context.Background(), token)
	Expect(err).NotTo(HaveOccurred())
}
