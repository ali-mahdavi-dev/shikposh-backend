package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity/product_aggregate"
)

func (h *ProductQueryHandler) GetAllTags(ctx context.Context) ([]*product_aggregate.Tag, error) {
	var tags []*product_aggregate.Tag
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		tags, err = h.uow.Tag(ctx).GetAll(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return tags, err
}

func (h *ProductQueryHandler) CreateTag(ctx context.Context, name string) (*product_aggregate.Tag, error) {
	var tag *product_aggregate.Tag
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		tag, err = h.uow.Tag(ctx).FindOrCreateByName(ctx, name)
		if err != nil {
			return err
		}
		return nil
	})
	return tag, err
}
