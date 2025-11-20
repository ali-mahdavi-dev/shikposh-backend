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

func (h *UserHandler) LoginHandler(ctx context.Context, cmd *commands.LoginUser) (string, error) {
	var accessToken string

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		user, err := h.uow.User(ctx).FindByUserName(ctx, cmd.UserName)
		if err != nil {
			if errors.Is(err, repository.ErrUserNotFound) {
				return apperrors.NotFound(accountphrases.UserNotFound)
			}
			return fmt.Errorf("UserHandler.LoginHandler fail get user by username: %w", err)
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
		accessToken, err = jwt.GenerateToken(h.cfg.JWT.AccessTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
		if err != nil {
			return fmt.Errorf("UserHandler.LoginHandler fail generate token: %w", err)
		}

		// Save new token
		err = h.uow.Token(ctx).Save(ctx, entity.NewToken(accessToken, user.ID))
		if err != nil {
			return fmt.Errorf("UserHandler.LoginHandler fail save token to db: %w", err)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
