package query

import (
	"context"
	"fmt"
	"strconv"

	"shikposh-backend/internal/seller/adapter/repository"
	sellerentity "shikposh-backend/internal/seller/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *SellerQueryHandler) GetSellerByID(ctx context.Context, id string) (map[string]interface{}, error) {
	sellerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid seller ID: %w", err)
	}

	var seller *sellerentity.Seller
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		// Use BaseRepository's FindByID method (expects uint64)
		seller, err = h.uow.Seller(ctx).FindByID(ctx, sellerID)
		return err
	})
	if err != nil {
		if err == repository.ErrSellerNotFound {
			return nil, repository.ErrSellerNotFound
		}
		return nil, fmt.Errorf("failed to get seller: %w", err)
	}

	// Calculate total products for this seller
	totalProducts, err := h.calculateTotalProducts(ctx, sellerentity.SellerID(sellerID))
	if err != nil {
		logging.Warn("Failed to calculate total products for seller").
			WithError(err).
			WithInt64("seller_id", int64(sellerID)).
			Log()
	}

	result := seller.ToMap()
	result["total_products"] = totalProducts

	return result, nil
}

