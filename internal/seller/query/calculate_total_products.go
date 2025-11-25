package query

import (
	"context"

	sellerentity "shikposh-backend/internal/seller/domain/entity"
)

// calculateTotalProducts calculates the total number of products for a seller
// This is a placeholder - in a real scenario, products would have a seller_id field
func (h *SellerQueryHandler) calculateTotalProducts(ctx context.Context, sellerID sellerentity.SellerID) (int, error) {
	// TODO: Once products have seller_id, implement this:
	// SELECT COUNT(*) FROM products WHERE seller_id = ? AND deleted_at IS NULL
	// For now, return 0
	return 0, nil
}
