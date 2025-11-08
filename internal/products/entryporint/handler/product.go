package handler

import (
	"errors"
	"strconv"
	"strings"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/query"
	"shikposh-backend/internal/products/service_layer/command_handler"
	httpapi "shikposh-backend/pkg/framework/api/http"
	"shikposh-backend/pkg/framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

// convertProductsToMap converts a slice of products to map format for JSON response
func convertProductsToMap(products []*entity.Product) []map[string]interface{} {
	result := make([]map[string]interface{}, len(products))
	for i, product := range products {
		result[i] = product.ToMap()
	}
	return result
}

type ProductHandler struct {
	productQueryHandler  *query.ProductQueryHandler
	categoryQueryHandler *query.CategoryQueryHandler
	reviewQueryHandler   *query.ReviewQueryHandler
	reviewHandler        *command_handler.ReviewCommandHandler
	bus                  messagebus.MessageBus
}

func NewProductHandler(
	productQueryHandler *query.ProductQueryHandler,
	categoryQueryHandler *query.CategoryQueryHandler,
	reviewQueryHandler *query.ReviewQueryHandler,
	reviewHandler *command_handler.ReviewCommandHandler,
	bus messagebus.MessageBus,
) *ProductHandler {
	return &ProductHandler{
		productQueryHandler:  productQueryHandler,
		categoryQueryHandler: categoryQueryHandler,
		reviewQueryHandler:   reviewQueryHandler,
		reviewHandler:        reviewHandler,
		bus:                  bus,
	}
}

func (p *ProductHandler) RegisterRoutes(r fiber.Router) {
	publicRoute := r.Group("/api/v1/public")
	{
		// Products
		publicRoute.Get("/products", p.GetAllProducts)
		publicRoute.Get("/products/:slug", p.GetProductBySlug)
		publicRoute.Get("/products/featured", p.GetFeaturedProducts)
		publicRoute.Get("/products/category/:category", p.GetProductsByCategory)

		// Categories
		publicRoute.Get("/categories", p.GetAllCategories)

		// Reviews
		publicRoute.Get("/products/:id/reviews", p.GetReviewsByProductID)
		publicRoute.Post("/reviews", p.CreateReview)
		publicRoute.Patch("/reviews/:id", p.UpdateReviewHelpful)
	}
}

// GetAllProducts godoc
//
//	@Summary		Get all products
//	@Description	Retrieves all products with optional filtering
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			q			query		string	false	"Search query"
//	@Param			category	query		string	false	"Category slug"
//	@Param			min			query		number	false	"Minimum price"
//	@Param			max			query		number	false	"Maximum price"
//	@Param			rating		query		number	false	"Minimum rating"
//	@Param			featured	query		boolean	false	"Featured products only"
//	@Param			tags		query		string	false	"Comma-separated tags"
//	@Param			sort		query		string	false	"Sort order (price_asc, price_desc, rating, newest)"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products [get]
func (p *ProductHandler) GetAllProducts(c fiber.Ctx) error {
	ctx := c.Context()

	// Parse query parameters
	filters := repository.ProductFilters{}
	if q := c.Query("q"); q != "" {
		filters.Query = &q
	}
	if category := c.Query("category"); category != "" {
		filters.Category = &category
	}
	if min := c.Query("min"); min != "" {
		if minPrice := cast.ToFloat64(min); minPrice > 0 {
			filters.MinPrice = &minPrice
		}
	}
	if max := c.Query("max"); max != "" {
		if maxPrice := cast.ToFloat64(max); maxPrice > 0 {
			filters.MaxPrice = &maxPrice
		}
	}
	if rating := c.Query("rating"); rating != "" {
		if ratingVal := cast.ToFloat64(rating); ratingVal > 0 {
			filters.Rating = &ratingVal
		}
	}
	if featured := c.Query("featured"); featured == "true" {
		featuredVal := true
		filters.Featured = &featuredVal
	}
	if tags := c.Query("tags"); tags != "" {
		// Parse comma-separated tags
		tagList := strings.Split(tags, ",")
		cleanedTags := []string{}
		for _, tag := range tagList {
			if trimmed := strings.TrimSpace(tag); trimmed != "" {
				cleanedTags = append(cleanedTags, trimmed)
			}
		}
		if len(cleanedTags) > 0 {
			filters.Tags = cleanedTags
		}
	}
	if sort := c.Query("sort"); sort != "" {
		filters.Sort = &sort
	}

	// Use filter if any filters are set, otherwise get all
	var productsList []*entity.Product
	var err error
	if filters.Query != nil || filters.Category != nil || filters.MinPrice != nil ||
		filters.MaxPrice != nil || filters.Rating != nil || filters.Featured != nil ||
		len(filters.Tags) > 0 || filters.Sort != nil {
		productsList, err = p.productQueryHandler.GetFilteredProducts(ctx, filters)
	} else {
		productsList, err = p.productQueryHandler.GetAllProducts(ctx)
	}

	if err != nil {
		return httpapi.ResError(c, err)
	}

	// Convert to map format
	productsMap := convertProductsToMap(productsList)
	return httpapi.ResSuccess(c, productsMap)
}

// GetProductBySlug godoc
//
//	@Summary		Get product by slug
//	@Description	Retrieves a single product by its slug
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"Product slug"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/{slug} [get]
func (p *ProductHandler) GetProductBySlug(c fiber.Ctx) error {
	ctx := c.Context()
	slug := c.Params("slug")
	if slug == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "slug is required"))
	}

	product, err := p.productQueryHandler.GetProductBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Product not found"))
		}
		return httpapi.ResError(c, err)
	}

	// Convert to map format
	productMap := product.ToMap()
	return httpapi.ResSuccess(c, productMap)
}

func (p *ProductHandler) GetProductByID(c fiber.Ctx) error {
	ctx := c.Context()
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	product, err := p.productQueryHandler.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Product not found"))
		}
		return httpapi.ResError(c, err)
	}

	// Convert to map format
	productMap := product.ToMap()
	return httpapi.ResSuccess(c, productMap)
}

// GetFeaturedProducts godoc
//
//	@Summary		Get featured products
//	@Description	Retrieves all featured products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/featured [get]
func (p *ProductHandler) GetFeaturedProducts(c fiber.Ctx) error {
	ctx := c.Context()

	products, err := p.productQueryHandler.GetFeaturedProducts(ctx)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	// Convert to map format
	productsMap := convertProductsToMap(products)
	return httpapi.ResSuccess(c, productsMap)
}

// GetProductsByCategory godoc
//
//	@Summary		Get products by category
//	@Description	Retrieves all products in a specific category
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			category	path		string	true	"Category slug"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/category/{category} [get]
func (p *ProductHandler) GetProductsByCategory(c fiber.Ctx) error {
	ctx := c.Context()
	categorySlug := c.Params("category")

	products, err := p.productQueryHandler.GetProductsByCategory(ctx, categorySlug)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Category not found"))
		}
		return httpapi.ResError(c, err)
	}

	// Convert to map format
	productsMap := convertProductsToMap(products)
	return httpapi.ResSuccess(c, productsMap)
}

// GetAllCategories godoc
//
//	@Summary		Get all categories
//	@Description	Retrieves all product categories
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/categories [get]
func (p *ProductHandler) GetAllCategories(c fiber.Ctx) error {
	ctx := c.Context()

	categories, err := p.categoryQueryHandler.GetAllCategories(ctx)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, categories)
}

// GetReviewsByProductID godoc
//
//	@Summary		Get reviews by product ID
//	@Description	Retrieves all reviews for a specific product
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Product ID"
//	@Success		200	{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/{id}/reviews [get]
func (p *ProductHandler) GetReviewsByProductID(c fiber.Ctx) error {
	ctx := c.Context()
	productID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	reviews, err := p.reviewQueryHandler.GetReviewsByProductID(ctx, productID)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Product not found"))
		}
		return httpapi.ResError(c, err)
	}

	// Return paginated response
	pr := &httpapi.PaginationResult{
		Total: int64(len(reviews)),
		Skip:  0,
		Limit: int64(len(reviews)),
	}
	return httpapi.ResPage(c, reviews, pr)
}

// CreateReview godoc
//
//	@Summary		Create a review
//	@Description	Creates a new review for a product
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			request	body		commands.CreateReview	true	"CreateReview request"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/reviews [post]
func (p *ProductHandler) CreateReview(c fiber.Ctx) error {
	ctx := c.Context()
	cmd := new(commands.CreateReview)

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	// Get user ID from context (set by auth middleware)
	userID := c.Get("user_id")
	if userID != "" {
		cmd.UserID = cast.ToUint64(userID)
	}

	result, err := p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, result)
}

// UpdateReviewHelpful godoc
//
//	@Summary		Update review helpful count
//	@Description	Increments helpful or notHelpful count for a review
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint64							true	"Review ID"
//	@Param			request	body		commands.UpdateReviewHelpful	true	"UpdateReviewHelpful request"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/reviews/{id} [patch]
func (p *ProductHandler) UpdateReviewHelpful(c fiber.Ctx) error {
	ctx := c.Context()
	reviewID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	cmd := new(commands.UpdateReviewHelpful)
	cmd.ReviewID = reviewID

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	result, err := p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, result)
}
