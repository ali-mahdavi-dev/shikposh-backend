package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/errors/phrases"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type ReviewCommandHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewReviewCommandHandler(uow unit_of_work.PGUnitOfWork) *ReviewCommandHandler {
	return &ReviewCommandHandler{uow: uow}
}

func (h *ReviewCommandHandler) CreateReviewHandler(ctx context.Context, cmd *commands.CreateReview) (*entity.Review, error) {
	var review *entity.Review

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Verify product exists
		product, err := h.uow.Product(ctx).FindByID(ctx, cmd.ProductID)
		if err != nil {
			if errors.Is(err, repository.ErrProductNotFound) {
				return apperrors.NotFound(phrases.UserNotFound)
			}
			return fmt.Errorf("ReviewCommandHandler.CreateReviewHandler error finding product: %w", err)
		}

		// Create review
		review = &entity.Review{
			ProductID:  cmd.ProductID,
			UserID:     cmd.UserID,
			UserName:   cmd.UserName,
			Rating:     cmd.Rating,
			Comment:    cmd.Comment,
			Helpful:    0,
			NotHelpful: 0,
			Verified:   false,
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
		return nil, err
	}

	return review, nil
}

func (h *ReviewCommandHandler) UpdateReviewHelpfulHandler(ctx context.Context, cmd *commands.UpdateReviewHelpful) (*entity.Review, error) {
	var review *entity.Review

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		review, err = h.uow.Review(ctx).FindByID(ctx, cmd.ReviewID)
		if err != nil {
			if errors.Is(err, repository.ErrReviewNotFound) {
				return apperrors.NotFound(phrases.UserNotFound)
			}
			return fmt.Errorf("ReviewCommandHandler.UpdateReviewHelpfulHandler error finding review: %w", err)
		}

		if cmd.Type == "helpful" {
			review.Helpful++
		} else if cmd.Type == "notHelpful" {
			review.NotHelpful++
		}

		if err := h.uow.Review(ctx).Save(ctx, review); err != nil {
			return fmt.Errorf("ReviewCommandHandler.UpdateReviewHelpfulHandler error saving review: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return review, nil
}
