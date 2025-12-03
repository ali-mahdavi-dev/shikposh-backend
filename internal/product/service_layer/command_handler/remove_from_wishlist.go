package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/product/adapter/repository"
	"shikposh-backend/internal/product/domain/commands"

	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *WishlistCommandHandler) RemoveFromWishlistHandler(ctx context.Context, cmd *commands.RemoveFromWishlist) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		err := h.uow.Wishlist(ctx).DeleteByUserAndProduct(ctx, cmd.UserID, cmd.ProductID)
		if err != nil {
			if errors.Is(err, repository.ErrWishlistNotFound) {
				return apperrors.NotFound("Product not in wishlist")
			}
			return fmt.Errorf("RemoveFromWishlistHandler: error removing from wishlist: %w", err)
		}

		return nil
	})
}

