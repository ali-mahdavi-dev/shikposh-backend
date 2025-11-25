package handler

import (
	"shikposh-backend/internal/products/domain/entity"
	unitofwork "shikposh-backend/internal/unit_of_work"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	"github.com/ali-mahdavi-dev/shikposh-framework/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

type WishlistHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewWishlistHandler(uow unitofwork.PGUnitOfWork) *WishlistHandler {
	return &WishlistHandler{uow: uow}
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

// GetWishlist returns all wishlist items for the authenticated user
func (w *WishlistHandler) GetWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Get("user_id")
	if userID == "" {
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

// AddToWishlist adds a product to the user's wishlist
func (w *WishlistHandler) AddToWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Get("user_id")
	if userID == "" {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	var req AddToWishlistRequest
	if err := httpapi.ParseJSON(c, &req); err != nil {
		return httpapi.ResError(c, err)
	}

	wishlist := entity.NewWishlist(cast.ToUint64(userID), req.ProductID)
	if err := w.uow.Wishlist(ctx).Add(ctx, wishlist); err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success": true,
		"message": "محصول به علاقه‌مندی‌ها اضافه شد",
	})
}

// RemoveFromWishlist removes a product from the user's wishlist
func (w *WishlistHandler) RemoveFromWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Get("user_id")
	if userID == "" {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	productID := cast.ToUint64(c.Params("productId"))

	err := w.uow.Wishlist(ctx).DeleteByUserAndProduct(ctx, cast.ToUint64(userID), productID)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success": true,
		"message": "محصول از علاقه‌مندی‌ها حذف شد",
	})
}

// ToggleWishlist toggles a product in the user's wishlist
func (w *WishlistHandler) ToggleWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Get("user_id")
	if userID == "" {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	var req ToggleWishlistRequest
	if err := httpapi.ParseJSON(c, &req); err != nil {
		return httpapi.ResError(c, err)
	}

	uid := cast.ToUint64(userID)

	// Check if item exists
	_, err := w.uow.Wishlist(ctx).FindByUserAndProduct(ctx, uid, req.ProductID)
	if err == nil {
		// Item exists, remove it
		if err := w.uow.Wishlist(ctx).DeleteByUserAndProduct(ctx, uid, req.ProductID); err != nil {
			return httpapi.ResError(c, err)
		}
		return httpapi.ResSuccess(c, map[string]interface{}{
			"success": true,
			"added":   false,
			"message": "محصول از علاقه‌مندی‌ها حذف شد",
		})
	}

	// Item doesn't exist, add it
	wishlist := entity.NewWishlist(uid, req.ProductID)
	if err := w.uow.Wishlist(ctx).Add(ctx, wishlist); err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success": true,
		"added":   true,
		"message": "محصول به علاقه‌مندی‌ها اضافه شد",
	})
}

// SyncWishlist syncs local wishlist with server
func (w *WishlistHandler) SyncWishlist(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Get("user_id")
	if userID == "" {
		return httpapi.ResError(c, errors.Unauthorized("USER_NOT_AUTHENTICATED"))
	}

	var req SyncWishlistRequest
	if err := httpapi.ParseJSON(c, &req); err != nil {
		return httpapi.ResError(c, err)
	}

	uid := cast.ToUint64(userID)

	// Get existing wishlist
	existing, err := w.uow.Wishlist(ctx).GetProductIDs(ctx, uid)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	existingMap := make(map[uint64]bool)
	for _, id := range existing {
		existingMap[id] = true
	}

	// Add new items from local storage
	for _, productID := range req.ProductIDs {
		if !existingMap[productID] {
			wishlist := entity.NewWishlist(uid, productID)
			_ = w.uow.Wishlist(ctx).Add(ctx, wishlist)
		}
	}

	// Get updated list
	productIDs, err := w.uow.Wishlist(ctx).GetProductIDs(ctx, uid)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success":     true,
		"product_ids": productIDs,
	})
}
