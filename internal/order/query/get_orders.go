package query

import (
	"context"
	"shikposh-backend/internal/order/adapter/repository"
	"shikposh-backend/internal/order/domain/entity"
	unitofwork "shikposh-backend/internal/unit_of_work"

	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

type OrderQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewOrderQueryHandler(uow unitofwork.PGUnitOfWork) *OrderQueryHandler {
	return &OrderQueryHandler{
		uow: uow,
	}
}

func (h *OrderQueryHandler) GetOrdersByUserID(ctx context.Context, userID uint64, filters repository.OrderFilters) ([]*entity.Order, error) {
	var orders []*entity.Order
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		orders, err = h.uow.Order(ctx).GetByUserID(ctx, userID, filters)
		return err
	})
	return orders, err
}

func (h *OrderQueryHandler) GetOrderByID(ctx context.Context, orderID uint64, userID uint64) (*entity.Order, error) {
	var order *entity.Order
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		order, err = h.uow.Order(ctx).GetByIDWithItems(ctx, orderID)
		if err != nil {
			return err
		}

		// Verify order belongs to user
		if order.UserID != userID {
			return apperrors.Forbidden("access denied")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (h *OrderQueryHandler) GetOrderByOrderNumber(ctx context.Context, orderNumber string) (*entity.Order, error) {
	var order *entity.Order
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		order, err = h.uow.Order(ctx).GetByOrderNumber(ctx, orderNumber)
		return err
	})
	return order, err
}
