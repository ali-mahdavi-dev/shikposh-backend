package command_handler

import (
	"context"
	"errors"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *UserHandler) LogoutHandler(ctx context.Context, cmd *commands.Logout) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		token, err := h.uow.Token(ctx).FindByUserID(ctx, entity.UserID(cmd.UserID))
		if err != nil {
			if errors.Is(err, repository.ErrTokenNotFound) {
				return apperrors.NotFound(accountphrases.UserNotFound)
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
