package command_handler

import (
	"context"
	"fmt"

	"shikposh-backend/config"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"
	"github.com/ali-mahdavi-dev/framework/api/jwt"
	apperrors "github.com/ali-mahdavi-dev/framework/errors"
	"github.com/ali-mahdavi-dev/framework/errors/phrases"
	"shikposh-backend/internal/unit_of_work"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	uow unitofwork.PGUnitOfWork
	cfg *config.Config
}

type RegisterResult struct {
	UserID uint64 `json:"user_id"`
}

type LoginResult struct {
	Access string `json:"access"`
}

func NewUserHandler(uow unitofwork.PGUnitOfWork, cfg *config.Config) *UserHandler {
	return &UserHandler{uow: uow, cfg: cfg}
}

func (h *UserHandler) RegisterHandler(ctx context.Context, cmd *commands.RegisterUser) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Check if username already exists
		_, err := h.uow.User(ctx).FindByUserName(ctx, cmd.UserName)
		if err != nil {
			if err != repository.ErrUserNotFound {
				return fmt.Errorf("UserHandler.Register error checking existing username: %w", err)
			}
		} else {
			return apperrors.Conflict(phrases.UserAlreadyExists)
		}

		// Hash password before saving
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("UserHandler.Register error hashing password: %w", err)
		}

		user := entity.NewUser(
			cmd.AvatarIdentifier,
			cmd.UserName,
			cmd.FirstName,
			cmd.LastName,
			cmd.Email,
			string(hashedPassword),
		)

		err = h.uow.User(ctx).Save(ctx, user)
		if err != nil {
			return fmt.Errorf("UserHandler.Register error saving user: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) LogoutHandler(ctx context.Context, cmd *commands.Logout) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		token, err := h.uow.Token(ctx).FindByUserID(ctx, entity.UserID(cmd.UserID))
		if err != nil {
			if errors.Is(err, repository.ErrTokenNotFound) {
				return apperrors.NotFound(phrases.UserNotFound)
			}

			return fmt.Errorf("UserHandler.LogoutHandler failed to get token by userID: %w", err)
		}

		if err := h.uow.Token(ctx).Remove(ctx, token, false); err != nil {
			return fmt.Errorf("UserHandler.LogoutHandler failed to remove existing token: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("UserHandler.LogoutHandler fail transaction: %w", err)
	}

	return nil
}

func (h *UserHandler) LoginHandler(ctx context.Context, cmd *commands.LoginUser) (string, error) {
	var accessToken string

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		user, err := h.uow.User(ctx).FindByUserName(ctx, cmd.UserName)
		if err != nil {
			if errors.Is(err, repository.ErrUserNotFound) {
				return apperrors.NotFound(phrases.UserNotFound)
			}
			return fmt.Errorf("UserHandler.LoginHandler fail get user by username: %w", err)
		}

		// Verify password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cmd.Password))
		if err != nil {
			return apperrors.Unauthorized(phrases.UserNotFound)
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
