package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/domain/specification"
	appadapter "github.com/ali-mahdavi-dev/framework/adapter"
	apperrors "github.com/ali-mahdavi-dev/framework/errors"
	"github.com/ali-mahdavi-dev/framework/errors/phrases"
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

		// Validate review using specification pattern
		canBePublishedSpec := specification.NewReviewCanBePublishedSpecification()
		if !canBePublishedSpec.IsSatisfiedBy(review) {
			return apperrors.Validation("", "Review must have a valid rating (1-5), a comment, and a valid user")
		}

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
