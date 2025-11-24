package handler

import (
	"errors"
	"strconv"
	"strings"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/query"
	"shikposh-backend/internal/products/service_layer/command_handler"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

// convertProductsToMap converts a slice of products to map format for JSON response
func convertProductsToMap(products []*productaggregate.Product) []map[string]interface{} {
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
	productHandler       *command_handler.ProductCommandHandler
	bus                  messagebus.MessageBus
}

func NewProductHandler(
	productQueryHandler *query.ProductQueryHandler,
	categoryQueryHandler *query.CategoryQueryHandler,
	reviewQueryHandler *query.ReviewQueryHandler,
	reviewHandler *command_handler.ReviewCommandHandler,
	productHandler *command_handler.ProductCommandHandler,
	bus messagebus.MessageBus,
) *ProductHandler {
	return &ProductHandler{
		productQueryHandler:  productQueryHandler,
		categoryQueryHandler: categoryQueryHandler,
		reviewQueryHandler:   reviewQueryHandler,
		reviewHandler:        reviewHandler,
		productHandler:       productHandler,
		bus:                  bus,
	}
}

func (p *ProductHandler) RegisterRoutes(r fiber.Router) {
	publicRoute := r.Group("/api/v1/public")
	{
		// Products
		publicRoute.Get("/products", p.GetAllProducts)
		publicRoute.Get("/products/featured", p.GetFeaturedProducts)
		publicRoute.Get("/products/category/:category", p.GetProductsByCategory)
		// Reviews route must come before /products/:slug to avoid route conflict
		publicRoute.Get("/products/:slug/reviews", p.GetReviewsByProductID)
		publicRoute.Get("/products/:slug", p.GetProductBySlug)
		publicRoute.Post("/products/cart", p.GetProductsForCart)

		// Categories
		publicRoute.Get("/categories", p.GetAllCategories)

		// Reviews
		publicRoute.Get("/reviews", p.GetReviews)
		publicRoute.Post("/reviews", p.CreateReview)
		publicRoute.Patch("/reviews/:slug", p.UpdateReviewHelpful)
	}

	// Admin routes for product CRUD
	adminRoute := r.Group("/api/v1/admin")
	{
		adminRoute.Post("/products", p.CreateProduct)
		adminRoute.Put("/products/:id", p.UpdateProduct)
		adminRoute.Delete("/products/:id", p.DeleteProduct)
	}
}

// parseProductFilters parses query parameters into ProductFilters
func parseProductFilters(c fiber.Ctx) repository.ProductFilters {
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
		tagList := strings.Split(tags, ",")
		cleanedTags := make([]string, 0, len(tagList))
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

	return filters
}

// hasFilters checks if any filters are set
func hasFilters(filters repository.ProductFilters) bool {
	return filters.Query != nil || filters.Category != nil || filters.MinPrice != nil ||
		filters.MaxPrice != nil || filters.Rating != nil || filters.Featured != nil ||
		len(filters.Tags) > 0 || filters.Sort != nil
}

// GetAllProducts godoc
//
//	@Summary		Get all products
//	@Description	Retrieves all products with optional filtering (uses Elasticsearch only)
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
	filters := parseProductFilters(c)

	var products []map[string]interface{}
	var err error

	if hasFilters(filters) {
		products, err = p.productQueryHandler.GetFilteredProductsAsMaps(ctx, filters)
	} else {
		products, err = p.productQueryHandler.GetAllProductsAsMaps(ctx)
	}

	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, products)
}

// GetProductBySlug godoc
//
//	@Summary		Get product by slug
//	@Description	Retrieves a single product by its slug (uses Elasticsearch only)
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

	productMap, err := p.productQueryHandler.GetProductBySlugAsMap(ctx, slug)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Product not found"))
		}
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, productMap)
}

// GetFeaturedProducts godoc
//
//	@Summary		Get featured products
//	@Description	Retrieves all featured products (uses Elasticsearch only)
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/featured [get]
func (p *ProductHandler) GetFeaturedProducts(c fiber.Ctx) error {
	ctx := c.Context()

	products, err := p.productQueryHandler.GetFeaturedProductsAsMaps(ctx)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, products)
}

// GetProductsByCategory godoc
//
//	@Summary		Get products by category
//	@Description	Retrieves all products in a specific category (uses Elasticsearch only)
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			category	path		string	true	"Category slug"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/category/{category} [get]
func (p *ProductHandler) GetProductsByCategory(c fiber.Ctx) error {
	ctx := c.Context()
	categorySlug := c.Params("category")

	products, err := p.productQueryHandler.GetProductsByCategoryAsMaps(ctx, categorySlug)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Category not found"))
		}
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, products)
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
//	@Summary		Get reviews by product slug
//	@Description	Retrieves all reviews for a specific product by slug
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"Product slug"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/{slug}/reviews [get]
func (p *ProductHandler) GetReviewsByProductID(c fiber.Ctx) error {
	ctx := c.Context()
	slug := c.Params("slug")
	if slug == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "product slug is required"))
	}

	// Find product by slug first
	product, err := p.productQueryHandler.GetProductBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Product not found"))
		}
		return httpapi.ResError(c, err)
	}

	productID := product.ID

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

// GetReviews godoc
//
//	@Summary		Get reviews by product slug or ID
//	@Description	Retrieves all reviews for a specific product. Accepts productId as query parameter (can be slug or numeric ID).
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			productId	query		string	true	"Product slug or numeric ID"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/reviews [get]
func (p *ProductHandler) GetReviews(c fiber.Ctx) error {
	ctx := c.Context()
	productIdOrSlug := c.Query("productId")
	if productIdOrSlug == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "productId query parameter is required"))
	}

	var productID productaggregate.ProductID

	// Check if the parameter is numeric (ID)
	if id, err := strconv.ParseUint(productIdOrSlug, 10, 64); err == nil {
		// It's a numeric ID
		productID = productaggregate.ProductID(id)
	} else {
		// It's a slug, need to find product by slug first
		product, err := p.productQueryHandler.GetProductBySlug(ctx, productIdOrSlug)
		if err != nil {
			if errors.Is(err, repository.ErrProductNotFound) {
				return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Product not found"))
			}
			return httpapi.ResError(c, err)
		}
		productID = product.ID
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

	err := p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateReviewHelpful godoc
//
//	@Summary		Update review helpful count
//	@Description	Increments helpful or notHelpful count for a review
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		uint64							true	"Review ID"
//	@Param			request	body		commands.UpdateReviewHelpful	true	"UpdateReviewHelpful request"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/reviews/{slug} [patch]
func (p *ProductHandler) UpdateReviewHelpful(c fiber.Ctx) error {
	ctx := c.Context()
	reviewID, err := strconv.ParseUint(c.Params("slug"), 10, 64)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	cmd := new(commands.UpdateReviewHelpful)
	cmd.ReviewID = reviewID

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	err = p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Creates a new product with all its details, features, and specs
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		commands.CreateProduct	true	"CreateProduct request"
//	@Success		201		{object}	httpapi.ResponseResult
//	@Router			/api/v1/admin/products [post]
func (p *ProductHandler) CreateProduct(c fiber.Ctx) error {
	ctx := c.Context()
	cmd := new(commands.CreateProduct)

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	err := p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Updates an existing product. Only provided fields will be updated.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		uint64					true	"Product ID"
//	@Param			request	body		commands.UpdateProduct	true	"UpdateProduct request"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/admin/products/{id} [put]
func (p *ProductHandler) UpdateProduct(c fiber.Ctx) error {
	ctx := c.Context()
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	cmd := new(commands.UpdateProduct)
	cmd.ID = id

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	err = p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Deletes a product. Can perform soft delete or hard delete.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id			path		uint64	true	"Product ID"
//	@Param			soft_delete	query		boolean	false	"Soft delete (default: true)"
//	@Success		200			{object}	httpapi.ResponseResult
//	@Router			/api/v1/admin/products/{id} [delete]
func (p *ProductHandler) DeleteProduct(c fiber.Ctx) error {
	ctx := c.Context()
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	cmd := &commands.DeleteProduct{
		ID:         id,
		SoftDelete: true, // Default to soft delete
	}

	// Check query parameter for soft_delete
	if softDelete := c.Query("soft_delete"); softDelete != "" {
		cmd.SoftDelete = cast.ToBool(softDelete)
	}

	err = p.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetProductsForCart godoc
//
//	@Summary		Get products for cart
//	@Description	Retrieves products by IDs with only essential fields (name, image, price, discount)
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			request	body		[]string	true	"Array of product IDs"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/products/cart [post]
func (p *ProductHandler) GetProductsForCart(c fiber.Ctx) error {
	ctx := c.Context()

	var productIDs []string
	// Parse JSON body directly without validation (since it's a simple slice)
	if err := c.Bind().Body(&productIDs); err != nil {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "Invalid request body: expected array of product IDs"))
	}

	if len(productIDs) == 0 {
		return httpapi.ResSuccess(c, []interface{}{})
	}

	products, err := p.productQueryHandler.GetProductsByIDsAsMaps(ctx, productIDs)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	// Filter to only return name, image, price, discount, and id
	filteredProducts := make([]map[string]interface{}, 0, len(products))
	for _, product := range products {
		filtered := map[string]interface{}{
			"id": product["id"],
		}

		if name, ok := product["name"].(string); ok {
			filtered["name"] = name
		}
		if image, ok := product["image"].(string); ok {
			filtered["image"] = image
		}
		if price, ok := product["price"].(float64); ok {
			filtered["price"] = price
		} else if price, ok := product["price"].(int); ok {
			filtered["price"] = float64(price)
		}

		// Check if discount exists
		hasDiscount := false
		if discount, ok := product["discount"].(float64); ok && discount > 0 {
			hasDiscount = true
			filtered["discount"] = discount
		} else if discount, ok := product["discount"].(int); ok && discount > 0 {
			hasDiscount = true
			filtered["discount"] = float64(discount)
		}
		filtered["hasDiscount"] = hasDiscount

		filteredProducts = append(filteredProducts, filtered)
	}

	return httpapi.ResSuccess(c, filteredProducts)
}
