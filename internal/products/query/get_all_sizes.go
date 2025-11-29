package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity/product_aggregate"
)

func (h *ProductQueryHandler) GetAllSizes(ctx context.Context) ([]*product_aggregate.Size, error) {
	var sizes []*product_aggregate.Size
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		sizes, err = h.uow.Size(ctx).GetAll(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return sizes, err
}

