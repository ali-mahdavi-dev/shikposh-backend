package query

import (
	"context"

	"shikposh-backend/internal/product/domain/entity"
)

func (h *CategoryQueryHandler) GetCategoryBySlug(ctx context.Context, slug string) (*entity.Category, error) {
	var category *entity.Category
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		category, err = h.uow.Category(ctx).FindBySlug(ctx, slug)
		if err != nil {
			return err
		}
		return nil
	})
	return category, err
}
