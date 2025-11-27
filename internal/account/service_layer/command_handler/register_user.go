package command_handler

import (
	"context"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/api/jwt"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"

	"golang.org/x/crypto/bcrypt"
)

// RegisterResult contains the result of registration
type RegisterResult struct {
	Token        string
	RefreshToken string
	User         *entity.User
}

func (h *UserHandler) RegisterHandler(ctx context.Context, cmd *commands.RegisterUser) (*RegisterResult, error) {
	var registeredUser *entity.User

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Check if phone already exists
		_, err := h.uow.User(ctx).FindByPhone(ctx, cmd.Phone)
		if err != nil {
			if err != repository.ErrUserNotFound {
				return fmt.Errorf("UserHandler.Register error checking existing phone: %w", err)
			}
		} else {
			return apperrors.Conflict(accountphrases.PhoneAlreadyExists)
		}

		// Hash password if provided, otherwise use empty string
		var hashedPassword string
		if cmd.Password != "" {
			hashed, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
			if err != nil {
				return fmt.Errorf("UserHandler.Register error hashing password: %w", err)
			}
			hashedPassword = string(hashed)
		}

		// Use phone as avatar identifier if not provided
		avatarIdentifier := cmd.AvatarIdentifier
		if avatarIdentifier == "" {
			avatarIdentifier = cmd.Phone
		}

		user := entity.NewUser(
			avatarIdentifier,
			cmd.FirstName,
			cmd.LastName,
			cmd.Email,
			cmd.Phone,
			hashedPassword,
		)

		err = h.uow.User(ctx).Save(ctx, user)
		if err != nil {
			return fmt.Errorf("UserHandler.Register error saving user: %w", err)
		}

		registeredUser = user
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Generate tokens for the new user
	accessToken, err := jwt.GenerateToken(h.cfg.JWT.AccessTokenExpireDuration, h.cfg.JWT.Secret, uint64(registeredUser.ID))
	if err != nil {
		return nil, fmt.Errorf("UserHandler.Register error generating access token: %w", err)
	}

	refreshToken, err := jwt.GenerateToken(h.cfg.JWT.RefreshTokenExpireDuration, h.cfg.JWT.Secret, uint64(registeredUser.ID))
	if err != nil {
		return nil, fmt.Errorf("UserHandler.Register error generating refresh token: %w", err)
	}

	// Save token
	err = h.uow.Token(ctx).Save(ctx, entity.NewToken(accessToken, refreshToken, registeredUser.ID))
	if err != nil {
		return nil, fmt.Errorf("UserHandler.Register error saving token: %w", err)
	}

	return &RegisterResult{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         registeredUser,
	}, nil
}
