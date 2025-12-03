package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/product/adapter/repository"
	"shikposh-backend/internal/product/domain/commands"
	"shikposh-backend/internal/product/domain/entity"

	appadapter "github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *WishlistCommandHandler) AddToWishlistHandler(ctx context.Context, cmd *commands.AddToWishlist) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Check if product exists
		_, err := h.uow.Product(ctx).FindByID(ctx, cmd.ProductID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound("Product not found")
			}
			return fmt.Errorf("AddToWishlistHandler: error finding product: %w", err)
		}

		// Check if already in wishlist
		_, err = h.uow.Wishlist(ctx).FindByUserAndProduct(ctx, cmd.UserID, cmd.ProductID)
		if err == nil {
			// Already exists, return success (idempotent)
			return nil
		}
		if !errors.Is(err, repository.ErrWishlistNotFound) {
			return fmt.Errorf("AddToWishlistHandler: error checking existing wishlist: %w", err)
		}

		// Add to wishlist
		wishlist := entity.NewWishlist(cmd.UserID, cmd.ProductID)
		if err := h.uow.Wishlist(ctx).Add(ctx, wishlist); err != nil {
			return fmt.Errorf("AddToWishlistHandler: error adding to wishlist: %w", err)
		}

		return nil
	})
}
