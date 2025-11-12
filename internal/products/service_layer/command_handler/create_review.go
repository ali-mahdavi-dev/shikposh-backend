package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/errors/phrases"
)

func (h *ReviewCommandHandler) CreateReviewHandler(ctx context.Context, cmd *commands.CreateReview) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Verify product exists
		product, err := h.uow.Product(ctx).FindByID(ctx, cmd.ProductID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound(phrases.UserNotFound)
			}
			return fmt.Errorf("ReviewCommandHandler.CreateReviewHandler error finding product: %w", err)
		}

		// Create review
		review := entity.NewReview(cmd)

		if err := h.uow.Review(ctx).Save(ctx, review); err != nil {
			return fmt.Errorf("ReviewCommandHandler.CreateReviewHandler error saving review: %w", err)
		}

		// Update product rating and review count
		product.ReviewCount++
		// Recalculate average rating (simplified - in production, you might want to store this)
		// For now, we'll just increment the count
		if err := h.uow.Product(ctx).Save(ctx, product); err != nil {
			return fmt.Errorf("ReviewCommandHandler.CreateReviewHandler error updating product: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
