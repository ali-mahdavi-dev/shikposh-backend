package query

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type ProductQueryHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewProductQueryHandler(uow unit_of_work.PGUnitOfWork) *ProductQueryHandler {
	return &ProductQueryHandler{uow: uow}
}

func (h *ProductQueryHandler) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).GetAll(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) GetProductByID(ctx context.Context, id uint64) (*entity.Product, error) {
	var product *entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindByID(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return product, err
}

func (h *ProductQueryHandler) GetFeaturedProducts(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).FindFeatured(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) GetProductsByCategory(ctx context.Context, categorySlug string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).FindByCategorySlug(ctx, categorySlug)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) SearchProducts(ctx context.Context, query string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).Search(ctx, query)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) GetFilteredProducts(ctx context.Context, filters repository.ProductFilters) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).Filter(ctx, filters)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}
