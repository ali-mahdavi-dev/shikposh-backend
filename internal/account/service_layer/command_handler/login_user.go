package command_handler

import (
	"context"
	"errors"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/api/jwt"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

func (h *UserHandler) LoginHandler(ctx context.Context, cmd *commands.LoginUser) (*LoginResult, error) {
	var result *LoginResult

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		user, err := h.uow.User(ctx).FindByPhone(ctx, cmd.Phone)
		if err != nil {
			if errors.Is(err, repository.ErrUserNotFound) {
				return apperrors.NotFound(accountphrases.UserNotFound)
			}
			return fmt.Errorf("UserHandler.LoginHandler fail get user by phone: %w", err)
		}

		// Verify password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cmd.Password))
		if err != nil {
			return apperrors.Unauthorized(accountphrases.UserNotFound)
		}

		// Check if user has existing token and remove it
		token, err := h.uow.Token(ctx).FindByUserID(ctx, user.ID)
		if err != nil && !errors.Is(err, repository.ErrTokenNotFound) {
			return fmt.Errorf("UserHandler.LoginHandler failed to get token by userID: %w", err)
		}

		if token != nil {
			if err := h.uow.Token(ctx).Remove(ctx, token, false); err != nil {
				return fmt.Errorf("UserHandler.LoginHandler failed to remove existing token: %w", err)
			}
		}

		// Generate new access token
		accessToken, err := jwt.GenerateToken(h.cfg.JWT.AccessTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
		if err != nil {
			return fmt.Errorf("UserHandler.LoginHandler fail generate access token: %w", err)
		}

		// Generate new refresh token
		refreshToken, err := jwt.GenerateToken(h.cfg.JWT.RefreshTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
		if err != nil {
			return fmt.Errorf("UserHandler.LoginHandler fail generate refresh token: %w", err)
		}

		// Save new token
		err = h.uow.Token(ctx).Save(ctx, entity.NewToken(accessToken, refreshToken, user.ID))
		if err != nil {
			return fmt.Errorf("UserHandler.LoginHandler fail save token to db: %w", err)
		}

		result = &LoginResult{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
