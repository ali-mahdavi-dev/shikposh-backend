package handler

import (
	"errors"
	"strconv"

	"shikposh-backend/internal/orders/adapter/repository"
	"shikposh-backend/internal/orders/domain/entity"
	"shikposh-backend/internal/orders/query"
	"shikposh-backend/internal/orders/service_layer/command_handler"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

type OrderHandler struct {
	orderQueryHandler   *query.OrderQueryHandler
	orderCommandHandler *command_handler.OrderCommandHandler
}

func NewOrderHandler(
	orderQueryHandler *query.OrderQueryHandler,
	orderCommandHandler *command_handler.OrderCommandHandler,
) *OrderHandler {
	return &OrderHandler{
		orderQueryHandler:   orderQueryHandler,
		orderCommandHandler: orderCommandHandler,
	}
}

func (h *OrderHandler) RegisterRoutes(r fiber.Router) {
	// Protected routes - require authentication
	protectedRoute := r.Group("/api/v1/orders")
	{
		protectedRoute.Get("/", h.GetOrders)
		protectedRoute.Get("/:id", h.GetOrderByID)
		protectedRoute.Post("/:id/cancel", h.CancelOrder)
	}
}

// GetOrders godoc
//
//	@Summary		Get user orders
//	@Description	Retrieves all orders for the authenticated user with optional filtering
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			status	query		string	false	"Filter by status (pending, processing, confirmed, shipped, delivered, cancelled, refunded)"
//	@Param			limit	query		int		false	"Limit number of results"
//	@Param			offset	query		int		false	"Offset for pagination"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		401		{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/orders [get]
func (h *OrderHandler) GetOrders(c fiber.Ctx) error {
	ctx := c.Context()

	// Get user ID from context (set by auth middleware)
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		logging.Warn("GetOrders: user_id not found in context").Log()
		return httpapi.ResError(c, fiber.NewError(fiber.StatusUnauthorized, "user not authenticated"))
	}

	logging.Info("GetOrders: fetching orders for user").
		WithInt64("user_id", int64(userID)).
		Log()

	// Parse filters
	filters := parseOrderFilters(c)

	orders, err := h.orderQueryHandler.GetOrdersByUserID(ctx, userID, filters)
	if err != nil {
		logging.Error("GetOrders: failed to fetch orders").
			WithInt64("user_id", int64(userID)).
			WithError(err).
			Log()
		return httpapi.ResError(c, err)
	}

	logging.Info("GetOrders: found orders").
		WithInt64("user_id", int64(userID)).
		WithInt("count", len(orders)).
		Log()

	// Convert to map format for response
	ordersMap := make([]map[string]interface{}, len(orders))
	for i, order := range orders {
		ordersMap[i] = orderToMap(order)
	}

	// Return orders directly as array (frontend expects OrdersResponse with orders array)
	// The ResSuccess will wrap it in { success: true, data: ordersMap }
	// But frontend expects { orders: [...], total: ... } structure
	// So we return the full structure
	return httpapi.ResSuccess(c, map[string]interface{}{
		"orders": ordersMap,
		"total":  len(ordersMap),
	})
}

// GetOrderByID godoc
//
//	@Summary		Get order by ID
//	@Description	Retrieves a single order by its ID (must belong to authenticated user)
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Order ID"
//	@Success		200	{object}	httpapi.ResponseResult
//	@Failure		400	{object}	httpapi.ResponseResult	"ID is required"
//	@Failure		401	{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		404	{object}	httpapi.ResponseResult	"Order not found"
//	@Failure		500	{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	if idStr == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "id is required"))
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "invalid order id"))
	}

	// Get user ID from context
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusUnauthorized, "user not authenticated"))
	}

	order, err := h.orderQueryHandler.GetOrderByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "order not found"))
		}
		return httpapi.ResError(c, err)
	}

	// Verify order belongs to user
	if order.UserID != userID {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusForbidden, "access denied"))
	}

	return httpapi.ResSuccess(c, orderToMap(order))
}

// CancelOrder godoc
//
//	@Summary		Cancel order
//	@Description	Cancels an order (only if status is pending or processing)
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Order ID"
//	@Success		200	{object}	httpapi.ResponseResult
//	@Failure		400	{object}	httpapi.ResponseResult	"Order cannot be cancelled"
//	@Failure		401	{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		404	{object}	httpapi.ResponseResult	"Order not found"
//	@Failure		500	{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/orders/{id}/cancel [post]
func (h *OrderHandler) CancelOrder(c fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	if idStr == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "id is required"))
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "invalid order id"))
	}

	// Get user ID from context
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusUnauthorized, "user not authenticated"))
	}

	// Verify order belongs to user
	order, err := h.orderQueryHandler.GetOrderByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "order not found"))
		}
		return httpapi.ResError(c, err)
	}

	if order.UserID != userID {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusForbidden, "access denied"))
	}

	// Cancel order
	err = h.orderCommandHandler.CancelOrderHandler(ctx, id)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"message": "order cancelled successfully",
	})
}

func parseOrderFilters(c fiber.Ctx) repository.OrderFilters {
	filters := repository.OrderFilters{}

	if status := c.Query("status"); status != "" {
		orderStatus := entity.OrderStatus(status)
		filters.Status = &orderStatus
	}

	if limit := c.Query("limit"); limit != "" {
		if limitVal := cast.ToInt(limit); limitVal > 0 {
			filters.Limit = &limitVal
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if offsetVal := cast.ToInt(offset); offsetVal >= 0 {
			filters.Offset = &offsetVal
		}
	}

	return filters
}

func orderToMap(order *entity.Order) map[string]interface{} {
	items := make([]map[string]interface{}, len(order.Items))
	for i, item := range order.Items {
		items[i] = map[string]interface{}{
			"id":            item.ID,
			"product_id":    item.ProductID,
			"product_name":  item.ProductName,
			"product_slug":  item.ProductSlug,
			"product_image": item.ProductImage,
			"quantity":      item.Quantity,
			"price":         item.Price,
			"discount":      item.Discount,
			"color":         item.Color,
			"size":          item.Size,
			"variant_id":    item.VariantID,
		}
	}

	result := map[string]interface{}{
		"id":              order.ID,
		"order_number":    order.OrderNumber,
		"status":          order.Status,
		"total_amount":    order.TotalAmount,
		"discount_amount": order.DiscountAmount,
		"shipping_cost":   order.ShippingCost,
		"final_amount":    order.FinalAmount,
		"payment_method":  order.PaymentMethod,
		"payment_status":  order.PaymentStatus,
		"tracking_number": order.TrackingNumber,
		"created_at":      order.CreatedAt,
		"updated_at":      order.UpdatedAt,
		"shipped_at":      order.ShippedAt,
		"delivered_at":    order.DeliveredAt,
		"notes":           order.Notes,
		"items":           items,
	}

	if order.ShippingAddress != nil {
		result["shipping_address"] = map[string]interface{}{
			"id":          order.ShippingAddress.ID,
			"full_name":   order.ShippingAddress.FullName,
			"phone":       order.ShippingAddress.Phone,
			"address":     order.ShippingAddress.Address,
			"city":        order.ShippingAddress.City,
			"province":    order.ShippingAddress.Province,
			"postal_code": order.ShippingAddress.PostalCode,
		}
	}

	return result
}
