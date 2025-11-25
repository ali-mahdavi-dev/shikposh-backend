package query

import (
	"context"
	"fmt"

	sellerentity "shikposh-backend/internal/seller/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *SellerQueryHandler) SearchSellers(ctx context.Context, query string) ([]map[string]interface{}, error) {
	var sellers []*sellerentity.Seller
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		sellers, err = h.uow.Seller(ctx).Search(ctx, query)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search sellers: %w", err)
	}

	result := make([]map[string]interface{}, len(sellers))
	for i, seller := range sellers {
		sellerID := uint64(seller.ID)
		totalProducts, err := h.calculateTotalProducts(ctx, seller.ID)
		if err != nil {
			logging.Warn("Failed to calculate total products for seller").
				WithError(err).
				WithInt64("seller_id", int64(sellerID)).
				Log()
		}
		summary := seller.ToSummary()
		summary["total_products"] = totalProducts
		result[i] = summary
	}

	return result, nil
}
