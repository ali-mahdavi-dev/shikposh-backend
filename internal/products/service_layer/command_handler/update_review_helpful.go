package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	apperrors "github.com/shikposh/framework/errors"
	"github.com/shikposh/framework/errors/phrases"
)

func (h *ReviewCommandHandler) UpdateReviewHelpfulHandler(ctx context.Context, cmd *commands.UpdateReviewHelpful) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Validate enum type
		if !cmd.Type.IsValid() {
			return apperrors.Validation(phrases.DefaultValidationID, fmt.Sprintf("invalid review helpful type: %s, must be 'helpful' or 'notHelpful'", cmd.Type))
		}

		review, err := h.uow.Review(ctx).FindByID(ctx, cmd.ReviewID)
		if err != nil {
			if errors.Is(err, repository.ErrReviewNotFound) {
				return apperrors.NotFound(phrases.UserNotFound)
			}
			return fmt.Errorf("ReviewCommandHandler.UpdateReviewHelpfulHandler error finding review: %w", err)
		}

		switch cmd.Type {
		case commands.ReviewHelpfulTypeHelpful:
			review.Helpful++
		case commands.ReviewHelpfulTypeNotHelpful:
			review.NotHelpful++
		}

		if err := h.uow.Review(ctx).Modify(ctx, review); err != nil {
			return fmt.Errorf("ReviewCommandHandler.UpdateReviewHelpfulHandler error saving review: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
