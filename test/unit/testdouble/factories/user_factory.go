package factories

import (
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	"golang.org/x/crypto/bcrypt"
)

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

func CreateUser(username, email, password string) *entity.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &entity.User{
		ID:               1,
		AvatarIdentifier: "avatar123",
		UserName:         username,
		FirstName:        "John",
		LastName:         "Doe",
		Email:            email,
		Password:         string(hashedPassword),
	}
}
