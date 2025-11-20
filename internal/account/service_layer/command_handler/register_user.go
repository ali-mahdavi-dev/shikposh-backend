package command_handler

import (
	"context"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"

	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) RegisterHandler(ctx context.Context, cmd *commands.RegisterUser) error {
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

		// Check if username already exists
		_, err = h.uow.User(ctx).FindByUserName(ctx, cmd.UserName)
		if err != nil {
			if err != repository.ErrUserNotFound {
				return fmt.Errorf("UserHandler.Register error checking existing username: %w", err)
			}
		} else {
			return apperrors.Conflict(accountphrases.UserAlreadyExists)
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

		// Use phone as username if username is not provided
		userName := cmd.UserName
		if userName == "" {
			userName = cmd.Phone
		}

		// Use phone as avatar identifier if not provided
		avatarIdentifier := cmd.AvatarIdentifier
		if avatarIdentifier == "" {
			avatarIdentifier = cmd.Phone
		}

		user := entity.NewUser(
			avatarIdentifier,
			userName,
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

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
