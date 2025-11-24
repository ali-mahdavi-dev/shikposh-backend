package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity"
)

func (h *CategoryQueryHandler) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
	var categories []*entity.Category
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		categories, err = h.uow.Category(ctx).GetAll(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return categories, err
}

