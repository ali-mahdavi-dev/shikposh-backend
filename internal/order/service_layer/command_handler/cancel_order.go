package command_handler

import (
	"context"
	"errors"

	"shikposh-backend/internal/order/adapter/repository"
	"shikposh-backend/internal/order/domain/entity"
)

func (h *OrderCommandHandler) CancelOrderHandler(ctx context.Context, orderID uint64) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		orderRepo := h.uow.Order(ctx)

		order, err := orderRepo.FindByID(ctx, orderID)
		if err != nil {
			if errors.Is(err, repository.ErrOrderNotFound) {
				return errors.New("order not found")
			}
			return err
		}

		// Check if order can be cancelled
		if order.Status != entity.OrderStatusPaymentConfirmed && order.Status != entity.OrderStatusProcessing {
			return errors.New("order cannot be cancelled in current status")
		}

		// Update order status
		order.Status = entity.OrderStatusCancelled

		err = orderRepo.Modify(ctx, order)
		if err != nil {
			return err
		}

		return nil
	})
}
