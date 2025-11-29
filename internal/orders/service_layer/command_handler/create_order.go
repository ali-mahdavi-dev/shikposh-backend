package command_handler

import (
	"context"
	"fmt"

	"shikposh-backend/internal/orders/domain/commands"
	"shikposh-backend/internal/orders/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *OrderCommandHandler) CreateOrderHandler(ctx context.Context, cmd *commands.CreateOrder) (*entity.Order, error) {
	var createdOrder *entity.Order

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		orderRepo := h.uow.Order(ctx)

		// Generate unique order number
		orderNumber := entity.GenerateOrderNumber()

		// Create order items
		orderItems := make([]entity.OrderItem, len(cmd.Items))
		for i, item := range cmd.Items {
			orderItems[i] = entity.OrderItem{
				ProductID:    item.ProductID,
				ProductName:  item.ProductName,
				ProductSlug:  item.ProductSlug,
				ProductImage: item.ProductImage,
				Quantity:     item.Quantity,
				Price:        item.Price,
				Discount:     item.Discount,
				Color:        item.Color,
				Size:         item.Size,
				VariantID:    item.VariantID,
			}
		}

		// Create shipping address if provided
		var shippingAddress *entity.OrderAddress
		if cmd.ShippingAddress != nil {
			shippingAddress = &entity.OrderAddress{
				FullName:   cmd.ShippingAddress.FullName,
				Phone:      cmd.ShippingAddress.Phone,
				Address:    cmd.ShippingAddress.Address,
				City:       cmd.ShippingAddress.City,
				Province:   cmd.ShippingAddress.Province,
				PostalCode: cmd.ShippingAddress.PostalCode,
			}
		}

		// Create order with pending status
		paymentMethod := cmd.PaymentMethod
		order := &entity.Order{
			OrderNumber:     orderNumber,
			UserID:          cmd.UserID,
			Status:          entity.OrderStatusPending,
			TotalAmount:     cmd.TotalAmount,
			DiscountAmount:  cmd.DiscountAmount,
			ShippingCost:    cmd.ShippingCost,
			FinalAmount:     cmd.FinalAmount,
			PaymentMethod:   &paymentMethod,
			PaymentStatus:   entity.PaymentStatusPending,
			Items:           orderItems,
			ShippingAddress: shippingAddress,
		}

		// Save order
		err := orderRepo.Save(ctx, order)
		if err != nil {
			logging.Error("CreateOrderHandler: failed to create order").
				WithError(err).
				WithInt64("user_id", int64(cmd.UserID)).
				Log()
			return fmt.Errorf("failed to create order: %w", err)
		}

		createdOrder = order

		logging.Info("CreateOrderHandler: order created successfully").
			WithInt64("order_id", int64(order.ID)).
			WithString("order_number", order.OrderNumber).
			WithInt64("user_id", int64(cmd.UserID)).
			Log()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}
