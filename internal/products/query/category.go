package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/unit_of_work"
)

type CategoryQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewCategoryQueryHandler(uow unitofwork.PGUnitOfWork) *CategoryQueryHandler {
	return &CategoryQueryHandler{uow: uow}
}

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
