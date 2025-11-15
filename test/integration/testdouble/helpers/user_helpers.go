package helpers

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/entity"

	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindUserByUsername(db *gorm.DB, username string) *entity.User {
	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.FindByUserName(context.Background(), username)
	Expect(err).NotTo(HaveOccurred())
	return user
}

func FindTokenByUserID(db *gorm.DB, userID entity.UserID) *entity.Token {
	tokenRepo := repository.NewTokenRepository(db)
	token, err := tokenRepo.FindByUserID(context.Background(), userID)
	Expect(err).NotTo(HaveOccurred())
	return token
}

func FindTokenByUserIDWithError(db *gorm.DB, userID entity.UserID) (*entity.Token, error) {
	tokenRepo := repository.NewTokenRepository(db)
	return tokenRepo.FindByUserID(context.Background(), userID)
}

func IsPasswordHashed(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}


