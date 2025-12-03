package query

import (
	"context"
	"fmt"

	"shikposh-backend/internal/product/adapter/repository"
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type WishlistQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewWishlistQueryHandler(uow unitofwork.PGUnitOfWork) *WishlistQueryHandler {
	return &WishlistQueryHandler{uow: uow}
}

// GetProductIDs returns wishlist product IDs for a user
func (h *WishlistQueryHandler) GetProductIDs(ctx context.Context, userID uint64) ([]uint64, error) {
	var productIDs []uint64

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		productIDs, err = h.uow.Wishlist(ctx).GetProductIDs(ctx, userID)
		return err
	})

	return productIDs, err
}

// Exists returns whether a wishlist entry exists for user and product
func (h *WishlistQueryHandler) Exists(ctx context.Context, userID, productID uint64) (bool, error) {
	var exists bool

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		_, err := h.uow.Wishlist(ctx).FindByUserAndProduct(ctx, userID, productID)
		if err != nil {
			if err == repository.ErrWishlistNotFound {
				exists = false
				return nil
			}
			return err
		}

		exists = true
		return nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to check wishlist existence: %w", err)
	}

	return exists, nil
}
