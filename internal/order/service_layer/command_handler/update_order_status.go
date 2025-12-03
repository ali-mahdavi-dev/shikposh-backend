package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/order/adapter/repository"
	"shikposh-backend/internal/order/domain/commands"
	"shikposh-backend/internal/order/domain/entity"
)

// UpdateOrderStatusHandler handles updating order status, payment status and notes
func (h *OrderCommandHandler) UpdateOrderStatusHandler(ctx context.Context, cmd *commands.UpdateOrderStatus) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		orderRepo := h.uow.Order(ctx)

		order, err := orderRepo.FindByID(ctx, cmd.OrderID)
		if err != nil {
			if errors.Is(err, repository.ErrOrderNotFound) {
				return errors.New("order not found")
			}
			return err
		}

		if cmd.Status != nil && *cmd.Status != "" {
			order.Status = entity.OrderStatus(*cmd.Status)
		}

		if cmd.PaymentStatus != nil && *cmd.PaymentStatus != "" {
			order.PaymentStatus = entity.PaymentStatus(*cmd.PaymentStatus)
		}

		if cmd.Notes != nil {
			order.Notes = cmd.Notes
		}

		if err := orderRepo.Modify(ctx, order); err != nil {
			return fmt.Errorf("failed to update order status: %w", err)
		}

		return nil
	})
}


