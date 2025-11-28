package handler

import (
	"shikposh-backend/internal/products/domain/commands"
	unitofwork "shikposh-backend/internal/unit_of_work"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	"github.com/ali-mahdavi-dev/shikposh-framework/errors"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

type WishlistHandler struct {
	uow unitofwork.PGUnitOfWork
	bus messagebus.MessageBus
}

func NewWishlistHandler(uow unitofwork.PGUnitOfWork, bus messagebus.MessageBus) *WishlistHandler {
	return &WishlistHandler{uow: uow, bus: bus}
}

func (w *WishlistHandler) RegisterRoutes(r fiber.Router) {
	wishlistRoute := r.Group("/api/v1/wishlist")
	{
		wishlistRoute.Get("/", w.GetWishlist)
		wishlistRoute.Post("/", w.AddToWishlist)
		wishlistRoute.Delete("/:productId", w.RemoveFromWishlist)
		wishlistRoute.Post("/toggle", w.ToggleWishlist)
		wishlistRoute.Post("/sync", w.SyncWishlist)
	}
}

type AddToWishlistRequest struct {
	ProductID uint64 `json:"product_id"`
}

type ToggleWishlistRequest struct {
	ProductID uint64 `json:"product_id"`
}

type SyncWishlistRequest struct {
	ProductIDs []uint64 `json:"product_ids"`
}

// GetWishlist godoc
//
//	@Summary		Get user's wishlist
//	@Description	Returns all wishlist product IDs for the authenticated user
//	@Tags			wishlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	httpapi.ResponseResult
//	@Failure		401	{object}	httpapi.ResponseResult	"User not authenticated"
//	@Failure		500	{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/wishlist [get]
func (w *WishlistHandler) GetWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Locals("user_id")
	if userID == nil {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	productIDs, err := w.uow.Wishlist(ctx).GetProductIDs(ctx, cast.ToUint64(userID))
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"product_ids": productIDs,
	})
}

// AddToWishlist godoc
//
//	@Summary		Add product to wishlist
//	@Description	Adds a product to the authenticated user's wishlist
//	@Tags			wishlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		AddToWishlistRequest	true	"Product ID to add"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body"
//	@Failure		401		{object}	httpapi.ResponseResult	"User not authenticated"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/wishlist [post]
func (w *WishlistHandler) AddToWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Locals("user_id")
	if userID == nil {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	var req AddToWishlistRequest
	if err := httpapi.ParseJSON(c, &req); err != nil {
		return httpapi.ResError(c, err)
	}

	cmd := &commands.AddToWishlist{
		UserID:    cast.ToUint64(userID),
		ProductID: req.ProductID,
	}

	if err := w.bus.Handle(ctx, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success": true,
		"message": "محصول به علاقه‌مندی‌ها اضافه شد",
	})
}

// RemoveFromWishlist godoc
//
//	@Summary		Remove product from wishlist
//	@Description	Removes a product from the authenticated user's wishlist
//	@Tags			wishlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			productId	path		uint64	true	"Product ID to remove"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Failure		401			{object}	httpapi.ResponseResult	"User not authenticated"
//	@Failure		404			{object}	httpapi.ResponseResult	"Product not in wishlist"
//	@Failure		500			{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/wishlist/{productId} [delete]
func (w *WishlistHandler) RemoveFromWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Locals("user_id")
	if userID == nil {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	productID := cast.ToUint64(c.Params("productId"))

	cmd := &commands.RemoveFromWishlist{
		UserID:    cast.ToUint64(userID),
		ProductID: productID,
	}

	if err := w.bus.Handle(ctx, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success": true,
		"message": "محصول از علاقه‌مندی‌ها حذف شد",
	})
}

// ToggleWishlist godoc
//
//	@Summary		Toggle product in wishlist
//	@Description	Adds product to wishlist if not exists, removes if exists
//	@Tags			wishlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		ToggleWishlistRequest	true	"Product ID to toggle"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body"
//	@Failure		401		{object}	httpapi.ResponseResult	"User not authenticated"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/wishlist/toggle [post]
func (w *WishlistHandler) ToggleWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Locals("user_id")
	if userID == nil {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	var req ToggleWishlistRequest
	if err := httpapi.ParseJSON(c, &req); err != nil {
		return httpapi.ResError(c, err)
	}

	uid := cast.ToUint64(userID)

	// Check if item exists before toggling to determine response
	_, err := w.uow.Wishlist(ctx).FindByUserAndProduct(ctx, uid, req.ProductID)
	wasInWishlist := err == nil

	cmd := &commands.ToggleWishlist{
		UserID:    uid,
		ProductID: req.ProductID,
	}

	if err := w.bus.Handle(ctx, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	// Determine if item was added or removed based on previous state
	added := !wasInWishlist
	message := "محصول به علاقه‌مندی‌ها اضافه شد"
	if !added {
		message = "محصول از علاقه‌مندی‌ها حذف شد"
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success": true,
		"added":   added,
		"message": message,
	})
}

// SyncWishlist godoc
//
//	@Summary		Sync wishlist with server
//	@Description	Merges local wishlist with server wishlist for authenticated user
//	@Tags			wishlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		SyncWishlistRequest	true	"Product IDs from local storage"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body"
//	@Failure		401		{object}	httpapi.ResponseResult	"User not authenticated"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/wishlist/sync [post]
func (w *WishlistHandler) SyncWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Locals("user_id")
	if userID == nil {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	var req SyncWishlistRequest
	if err := httpapi.ParseJSON(c, &req); err != nil {
		return httpapi.ResError(c, err)
	}

	uid := cast.ToUint64(userID)

	cmd := &commands.SyncWishlist{
		UserID:     uid,
		ProductIDs: req.ProductIDs,
	}

	if err := w.bus.Handle(ctx, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	// Get updated list after sync
	productIDs, err := w.uow.Wishlist(ctx).GetProductIDs(ctx, uid)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success":     true,
		"product_ids": productIDs,
	})
}
