package query

import (
	"context"
	"fmt"
	"strconv"
)

func (h *SellerQueryHandler) GetSellerByProductID(ctx context.Context, productID string) (map[string]interface{}, error) {
	// Try to parse productID as numeric ID first
	productIDNum, err := strconv.ParseUint(productID, 10, 64)
	if err == nil {
		// It's a numeric ID, try to get from database
		err = h.uow.Do(ctx, func(ctx context.Context) error {
			product, err := h.uow.Product(ctx).FindByID(ctx, productIDNum)
			if err != nil {
				return err
			}
			// Check if product has sellerId field
			// For now, products don't have seller_id column, so we return an error
			_ = product
			return fmt.Errorf("products table does not have seller_id field yet")
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get product: %w", err)
		}
	}

	// Try as slug
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		product, err := h.uow.Product(ctx).FindBySlug(ctx, productID)
		if err != nil {
			return err
		}
		// Check if product has sellerId
		// For now, products don't have seller_id column, so we return an error
		_ = product
		return fmt.Errorf("products table does not have seller_id field yet")
	})

	// If we reach here, product was not found or doesn't have seller_id
	return nil, fmt.Errorf("product not found or seller_id not available in products table: productId=%s", productID)
}

