package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"

	appadapter "github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *WishlistCommandHandler) ToggleWishlistHandler(ctx context.Context, cmd *commands.ToggleWishlist) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Check if item exists
		_, err := h.uow.Wishlist(ctx).FindByUserAndProduct(ctx, cmd.UserID, cmd.ProductID)
		if err == nil {
			// Item exists, remove it
			if err := h.uow.Wishlist(ctx).DeleteByUserAndProduct(ctx, cmd.UserID, cmd.ProductID); err != nil {
				return fmt.Errorf("ToggleWishlistHandler: error removing from wishlist: %w", err)
			}
			return nil
		}

		if !errors.Is(err, repository.ErrWishlistNotFound) {
			return fmt.Errorf("ToggleWishlistHandler: error checking wishlist: %w", err)
		}

		// Item doesn't exist, add it
		// Check if product exists
		_, err = h.uow.Product(ctx).FindByID(ctx, cmd.ProductID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound("Product not found")
			}
			return fmt.Errorf("ToggleWishlistHandler: error finding product: %w", err)
		}

		wishlist := entity.NewWishlist(cmd.UserID, cmd.ProductID)
		if err := h.uow.Wishlist(ctx).Add(ctx, wishlist); err != nil {
			return fmt.Errorf("ToggleWishlistHandler: error adding to wishlist: %w", err)
		}

		return nil
	})
}
