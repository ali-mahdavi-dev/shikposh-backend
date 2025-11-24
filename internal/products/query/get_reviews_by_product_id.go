package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
)

func (h *ReviewQueryHandler) GetReviewsByProductID(ctx context.Context, productID productaggregate.ProductID) ([]*entity.Review, error) {
	var reviews []*entity.Review
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		reviews, err = h.uow.Review(ctx).FindByProductID(ctx, productID)
		if err != nil {
			return err
		}
		return nil
	})
	return reviews, err
}

