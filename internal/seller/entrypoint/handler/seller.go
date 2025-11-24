package handler

import (
	"errors"

	"shikposh-backend/internal/seller/adapter/repository"
	"shikposh-backend/internal/seller/query"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"

	"github.com/gofiber/fiber/v3"
)

type SellerHandler struct {
	sellerQueryHandler *query.SellerQueryHandler
}

func NewSellerHandler(sellerQueryHandler *query.SellerQueryHandler) *SellerHandler {
	return &SellerHandler{
		sellerQueryHandler: sellerQueryHandler,
	}
}

func (h *SellerHandler) RegisterRoutes(r fiber.Router) {
	publicRoute := r.Group("/api/v1/public")
	{
		// Sellers
		publicRoute.Get("/sellers", h.GetAllSellers)
		publicRoute.Get("/sellers/:id", h.GetSellerByID)
		publicRoute.Get("/products/:productId/seller", h.GetSellerByProductID)
	}
}

// GetAllSellers godoc
//
//	@Summary		Get all sellers
//	@Description	Retrieves all sellers with optional filtering
//	@Tags			sellers
//	@Accept			json
//	@Produce		json
//	@Param			q			query		string	false	"Search query"
//	@Param			categories_like	query		string	false	"Filter by category"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/sellers [get]
func (h *SellerHandler) GetAllSellers(c fiber.Ctx) error {
	ctx := c.Context()

	// Check for search query
	if q := c.Query("q"); q != "" {
		sellers, err := h.sellerQueryHandler.SearchSellers(ctx, q)
		if err != nil {
			return httpapi.ResError(c, err)
		}
		return httpapi.ResSuccess(c, sellers)
	}

	// Check for category filter
	if category := c.Query("categories_like"); category != "" {
		sellers, err := h.sellerQueryHandler.GetSellersByCategory(ctx, category)
		if err != nil {
			return httpapi.ResError(c, err)
		}
		return httpapi.ResSuccess(c, sellers)
	}

	// Get all sellers
	sellers, err := h.sellerQueryHandler.GetAllSellers(ctx)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, sellers)
}

// GetSellerByID godoc
//
//	@Summary		Get seller by ID
//	@Description	Retrieves a single seller by its ID
//	@Tags			sellers
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Seller ID"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/sellers/{id} [get]
func (h *SellerHandler) GetSellerByID(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	if id == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "id is required"))
	}

	seller, err := h.sellerQueryHandler.GetSellerByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrSellerNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Seller not found"))
		}
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, seller)
}

// GetSellerByProductID godoc
//
//	@Summary		Get seller by product ID
//	@Description	Retrieves the seller for a specific product
//	@Tags			sellers
//	@Accept			json
//	@Produce		json
//	@Param			productId	path		string	true	"Product ID"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/{productId}/seller [get]
func (h *SellerHandler) GetSellerByProductID(c fiber.Ctx) error {
	ctx := c.Context()
	productID := c.Params("productId")
	if productID == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "productId is required"))
	}

	seller, err := h.sellerQueryHandler.GetSellerByProductID(ctx, productID)
	if err != nil {
		if errors.Is(err, repository.ErrSellerNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Seller not found for this product"))
		}
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, seller)
}
