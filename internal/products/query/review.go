package query

import (
	"context"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type ReviewQueryHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewReviewQueryHandler(uow unit_of_work.PGUnitOfWork) *ReviewQueryHandler {
	return &ReviewQueryHandler{uow: uow}
}

func (h *ReviewQueryHandler) GetReviewsByProductID(ctx context.Context, productID uint64) ([]*entity.Review, error) {
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
