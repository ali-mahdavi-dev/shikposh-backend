package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity/product_aggregate"
)

func (h *ProductQueryHandler) GetAllColors(ctx context.Context) ([]*product_aggregate.Color, error) {
	var colors []*product_aggregate.Color
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		colors, err = h.uow.Color(ctx).GetAll(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return colors, err
}
