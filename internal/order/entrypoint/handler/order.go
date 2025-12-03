package handler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"shikposh-backend/internal/order/adapter/payment"
	"shikposh-backend/internal/order/adapter/repository"
	"shikposh-backend/internal/order/domain/commands"
	"shikposh-backend/internal/order/domain/entity"
	"shikposh-backend/internal/order/query"
	"shikposh-backend/internal/order/service_layer/command_handler"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

type OrderHandler struct {
	orderQueryHandler   *query.OrderQueryHandler
	orderCommandHandler *command_handler.OrderCommandHandler
	zarinPalService     *payment.ZarinPalService
}

func NewOrderHandler(
	orderQueryHandler *query.OrderQueryHandler,
	orderCommandHandler *command_handler.OrderCommandHandler,
	zarinPalService *payment.ZarinPalService,
) *OrderHandler {
	return &OrderHandler{
		orderQueryHandler:   orderQueryHandler,
		orderCommandHandler: orderCommandHandler,
		zarinPalService:     zarinPalService,
	}
}

func (h *OrderHandler) RegisterRoutes(r fiber.Router) {
	// Protected routes - require authentication
	protectedRoute := r.Group("/api/v1/orders")
	{
		protectedRoute.Get("/", h.GetOrders)
		protectedRoute.Get("/:id", h.GetOrderByID)
		protectedRoute.Post("/", h.CreateOrder)
		protectedRoute.Post("/:id/cancel", h.CancelOrder)
	}

	// Payment routes
	paymentRoute := r.Group("/api/v1/payments")
	{
		paymentRoute.Post("/zarinpal/verify", h.VerifyZarinPalPayment)
		paymentRoute.Get("/zarinpal/callback", h.ZarinPalCallback)
	}
}

// GetOrders godoc
//
//	@Summary		Get user orders
//	@Description	Retrieves all orders for the authenticated user with optional filtering
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			status	query		string	false	"Filter by status (payment_confirmed, processing, confirmed, shipped, delivered, cancelled, refunded)"
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
//	@Description	Cancels an order (only if status is payment_confirmed or processing)
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

// CreateOrder godoc
//
//	@Summary		Create order and get payment URL
//	@Description	Creates a new order and returns ZarinPal payment URL
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		commands.CreateOrder	true	"Order data"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request"
//	@Failure		401		{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/orders [post]
func (h *OrderHandler) CreateOrder(c fiber.Ctx) error {
	ctx := c.Context()

	// Get user ID from context
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusUnauthorized, "user not authenticated"))
	}

	var cmd commands.CreateOrder
	if err := c.Bind().Body(&cmd); err != nil {
		logging.Error("CreateOrder: failed to parse request body").
			WithError(err).
			Log()
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "invalid request body"))
	}

	// Set user ID from context
	cmd.UserID = userID

	// Validate payment method
	if cmd.PaymentMethod != "zarinpal" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "invalid payment method"))
	}

	// Create order
	order, err := h.orderCommandHandler.CreateOrderHandler(ctx, &cmd)
	if err != nil {
		logging.Error("CreateOrder: failed to create order").
			WithError(err).
			WithInt64("user_id", int64(userID)).
			Log()
		return httpapi.ResError(c, err)
	}

	// Request payment from ZarinPal
	description := fmt.Sprintf("سفارش %s", order.OrderNumber)
	authority, err := h.zarinPalService.RequestPayment(ctx, order.FinalAmount, description)
	if err != nil {
		logging.Error("CreateOrder: failed to request payment").
			WithError(err).
			WithInt64("order_id", int64(order.ID)).
			Log()

		// Check if it's a configuration error (should return 400 Bad Request)
		errorMsg := err.Error()
		errorMsgLower := strings.ToLower(errorMsg)

		// Check for various configuration-related error messages
		isConfigError := strings.Contains(errorMsgLower, "merchant id is not configured") ||
			strings.Contains(errorMsgLower, "invalid merchant id") ||
			(strings.Contains(errorMsgLower, "merchant id") &&
				(strings.Contains(errorMsgLower, "not configured") ||
					strings.Contains(errorMsgLower, "invalid") ||
					strings.Contains(errorMsgLower, "not correct")))

		if isConfigError {
			logging.Warn("CreateOrder: configuration error detected, returning 400").
				WithString("error", errorMsg).
				Log()
			// Use framework's Validation error which will return 400 Bad Request
			validationErr := apperrors.Validation("", errorMsg)
			return httpapi.ResError(c, validationErr)
		}

		// For other errors, return 500
		internalErr := apperrors.Internal("failed to initialize payment: " + errorMsg)
		return httpapi.ResError(c, internalErr)
	}

	// Store authority in order notes (temporary storage)
	authorityNote := fmt.Sprintf("payment_authority:%s", authority)
	order.Notes = &authorityNote
	orderRepo := h.orderCommandHandler.GetUOW().Order(ctx)
	if err := orderRepo.Modify(ctx, order); err != nil {
		logging.Error("CreateOrder: failed to update order with authority").
			WithError(err).
			WithInt64("order_id", int64(order.ID)).
			Log()
	}

	// Get payment URL
	paymentURL := h.zarinPalService.GetPaymentURL(authority)

	logging.Info("CreateOrder: order created and payment URL generated").
		WithInt64("order_id", int64(order.ID)).
		WithString("order_number", order.OrderNumber).
		WithString("authority", authority).
		Log()

	return httpapi.ResSuccess(c, map[string]interface{}{
		"order_id":     order.ID,
		"order_number": order.OrderNumber,
		"payment_url":  paymentURL,
		"authority":    authority,
	})
}

// VerifyZarinPalPayment godoc
//
//	@Summary		Verify ZarinPal payment
//	@Description	Verifies a payment after user returns from ZarinPal gateway
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			request	body		object	true	"Payment verification request"	example({"authority":"A0000000000000000000000000000000000000","order_id":1})
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request"
//	@Failure		401		{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/payments/zarinpal/verify [post]
func (h *OrderHandler) VerifyZarinPalPayment(c fiber.Ctx) error {
	ctx := c.Context()

	// Get user ID from context
	userID, ok := c.Locals("user_id").(uint64)
	if !ok {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusUnauthorized, "user not authenticated"))
	}

	var req struct {
		Authority string `json:"authority" validate:"required"`
		OrderID   uint64 `json:"order_id" validate:"required"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "invalid request body"))
	}

	// Get order
	order, err := h.orderQueryHandler.GetOrderByID(ctx, req.OrderID)
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

	// Verify order is in pending status
	if order.Status != entity.OrderStatusPending {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "order is not in pending status"))
	}

	// Verify payment with ZarinPal
	refID, err := h.zarinPalService.VerifyPayment(ctx, req.Authority, order.FinalAmount)
	if err != nil {
		logging.Error("VerifyZarinPalPayment: payment verification failed").
			WithError(err).
			WithInt64("order_id", int64(req.OrderID)).
			WithString("authority", req.Authority).
			Log()

		// Update order status to failed
		order.PaymentStatus = entity.PaymentStatusFailed
		orderRepo := h.orderCommandHandler.GetUOW().Order(ctx)
		if err := orderRepo.Modify(ctx, order); err != nil {
			logging.Error("VerifyZarinPalPayment: failed to update order status").
				WithError(err).
				Log()
		}

		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "payment verification failed"))
	}

	// Update order status
	order.Status = entity.OrderStatusPaymentConfirmed
	order.PaymentStatus = entity.PaymentStatusPaid
	refIDStr := fmt.Sprintf("ref_id:%d", refID)
	order.Notes = &refIDStr

	orderRepo := h.orderCommandHandler.GetUOW().Order(ctx)
	if err := orderRepo.Modify(ctx, order); err != nil {
		logging.Error("VerifyZarinPalPayment: failed to update order").
			WithError(err).
			WithInt64("order_id", int64(req.OrderID)).
			Log()
		return httpapi.ResError(c, fiber.NewError(fiber.StatusInternalServerError, "failed to update order"))
	}

	logging.Info("VerifyZarinPalPayment: payment verified successfully").
		WithInt64("order_id", int64(req.OrderID)).
		WithInt64("ref_id", refID).
		Log()

	return httpapi.ResSuccess(c, map[string]interface{}{
		"order_id":     order.ID,
		"order_number": order.OrderNumber,
		"status":       order.Status,
		"ref_id":       refID,
		"message":      "payment verified successfully",
	})
}

// ZarinPalCallback godoc
//
//	@Summary		ZarinPal payment callback
//	@Description	Handles callback from ZarinPal after payment
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			Status	query		string	true	"Payment status from ZarinPal"
//	@Param			Authority	query		string	true	"Payment authority from ZarinPal"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/payments/zarinpal/callback [get]
func (h *OrderHandler) ZarinPalCallback(c fiber.Ctx) error {
	status := c.Query("Status")
	authority := c.Query("Authority")

	if status == "" || authority == "" {
		return c.Redirect().To("/payment/callback?error=invalid_callback")
	}

	// Status "OK" means payment was successful
	if status != "OK" {
		return c.Redirect().To("/payment/callback?error=payment_failed&authority=" + authority)
	}

	// Redirect to frontend callback page with authority
	return c.Redirect().To("/payment/callback?authority=" + authority)
}
