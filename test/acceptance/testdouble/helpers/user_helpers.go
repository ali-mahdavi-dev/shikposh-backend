package helpers

import (
	"context"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/entity"

	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// FindUserByUsername finds a user by username
func FindUserByUsername(db *gorm.DB, username string) *entity.User {
	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.FindByUserName(context.Background(), username)
	Expect(err).NotTo(HaveOccurred())
	return user
}

// FindTokenByUserID finds a token by user ID
func FindTokenByUserID(db *gorm.DB, userID entity.UserID) *entity.Token {
	tokenRepo := repository.NewTokenRepository(db)
	token, err := tokenRepo.FindByUserID(context.Background(), userID)
	Expect(err).NotTo(HaveOccurred())
	return token
}

// FindTokenByUserIDWithError finds a token by user ID and returns error if not found
func FindTokenByUserIDWithError(db *gorm.DB, userID entity.UserID) (*entity.Token, error) {
	tokenRepo := repository.NewTokenRepository(db)
	return tokenRepo.FindByUserID(context.Background(), userID)
}

// VerifyPasswordHashed verifies that password is hashed correctly
func VerifyPasswordHashed(hashedPassword, plainPassword string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	Expect(err).NotTo(HaveOccurred())
}

// VerifyTokenRemoved verifies that token was removed
func VerifyTokenRemoved(db *gorm.DB, userID entity.UserID) {
	_, err := FindTokenByUserIDWithError(db, userID)
	Expect(err).To(HaveOccurred())
}

// VerifyUserCount verifies the count of users with given username
func VerifyUserCount(db *gorm.DB, username string, expectedCount int) {
	userRepo := repository.NewUserRepository(db)
	users, err := userRepo.FindByField(context.Background(), "user_name", username)
	Expect(err).NotTo(HaveOccurred())
	Expect(users).To(HaveLen(expectedCount))
}
