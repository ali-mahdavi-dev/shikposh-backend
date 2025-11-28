package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"

	appadapter "github.com/ali-mahdavi-dev/shikposh-framework/adapter"
)

func (h *WishlistCommandHandler) SyncWishlistHandler(ctx context.Context, cmd *commands.SyncWishlist) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Get existing wishlist
		existing, err := h.uow.Wishlist(ctx).GetProductIDs(ctx, cmd.UserID)
		if err != nil {
			return fmt.Errorf("SyncWishlistHandler: error getting existing wishlist: %w", err)
		}

		existingMap := make(map[uint64]bool)
		for _, id := range existing {
			existingMap[id] = true
		}

		// Add new items from local storage
		for _, productID := range cmd.ProductIDs {
			if !existingMap[productID] {
				// Check if product exists
				_, err := h.uow.Product(ctx).FindByID(ctx, productID)
				if err != nil {
					if errors.Is(err, appadapter.ErrEntityNotFound) {
						// Skip invalid product IDs
						continue
					}
					return fmt.Errorf("SyncWishlistHandler: error finding product: %w", err)
				}

				wishlist := entity.NewWishlist(cmd.UserID, productID)
				if err := h.uow.Wishlist(ctx).Add(ctx, wishlist); err != nil {
					// Continue on error (might be duplicate key)
					continue
				}
			}
		}

		return nil
	})
}

