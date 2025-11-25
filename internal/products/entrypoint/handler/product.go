package handler

import (
	"errors"
	"fmt"
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
	productHandler       *command_handler.ProductCommandHandler
	bus                  messagebus.MessageBus
}

func NewProductHandler(
	productQueryHandler *query.ProductQueryHandler,
	categoryQueryHandler *query.CategoryQueryHandler,
	productHandler *command_handler.ProductCommandHandler,
	bus messagebus.MessageBus,
) *ProductHandler {
	return &ProductHandler{
		productQueryHandler:  productQueryHandler,
		categoryQueryHandler: categoryQueryHandler,
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
		publicRoute.Get("/products/:slug", p.GetProductBySlug)
		publicRoute.Post("/products/cart", p.GetProductsForCart)

		// Categories
		publicRoute.Get("/categories", p.GetAllCategories)
		publicRoute.Get("/categories/:slug", p.GetCategoryBySlug)
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

// GetCategoryBySlug godoc
//
//	@Summary		Get category by slug
//	@Description	Retrieves a single category by its slug
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			slug	path		string	true	"Category slug"
//	@Success		200		{object}	httpapi.ResponseResult
//	@Router			/api/v1/public/categories/{slug} [get]
func (p *ProductHandler) GetCategoryBySlug(c fiber.Ctx) error {
	ctx := c.Context()
	slug := c.Params("slug")
	if slug == "" {
		return httpapi.ResError(c, fiber.NewError(fiber.StatusBadRequest, "slug is required"))
	}

	category, err := p.categoryQueryHandler.GetCategoryBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			return httpapi.ResError(c, fiber.NewError(fiber.StatusNotFound, "Category not found"))
		}
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, category)
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
		filtered := map[string]interface{}{}

		// Get ID (convert to string if needed)
		id := ""
		if idVal, ok := product["id"]; ok {
			switch v := idVal.(type) {
			case string:
				id = v
			case float64:
				id = fmt.Sprintf("%.0f", v)
			case int:
				id = strconv.Itoa(v)
			case int64:
				id = strconv.FormatInt(v, 10)
			case uint64:
				id = strconv.FormatUint(v, 10)
			default:
				id = fmt.Sprintf("%v", idVal)
			}
		}
		filtered["id"] = id

		// Get name (from "title" field in Elasticsearch)
		if title, ok := product["title"].(string); ok && title != "" {
			filtered["name"] = title
		} else {
			filtered["name"] = ""
		}

		// Get slug
		if slug, ok := product["slug"].(string); ok && slug != "" {
			filtered["slug"] = slug
		} else {
			filtered["slug"] = ""
		}

		// Get image (prefer first image from images object, fallback to thumbnail)
		image := ""
		if imagesMap, ok := product["images"].(map[string]interface{}); ok && imagesMap != nil {
			// Get first image from first color
			for _, colorImages := range imagesMap {
				if imagesArray, ok := colorImages.([]interface{}); ok && len(imagesArray) > 0 {
					if firstImage, ok := imagesArray[0].(string); ok && firstImage != "" {
						image = firstImage
						break
					}
				}
			}
		}
		// Fallback to thumbnail if no image found
		if image == "" {
			if thumbnail, ok := product["thumbnail"].(string); ok && thumbnail != "" {
				image = thumbnail
			}
		}
		filtered["image"] = image

		// Get price
		price := 0.0
		if priceVal, ok := product["price"].(float64); ok {
			price = priceVal
		} else if priceVal, ok := product["price"].(int); ok {
			price = float64(priceVal)
		} else if priceVal, ok := product["price"].(int64); ok {
			price = float64(priceVal)
		}
		filtered["price"] = price

		// Get discount
		discount := 0.0
		hasDiscount := false
		if discountVal, ok := product["discount"].(float64); ok && discountVal > 0 {
			discount = discountVal
			hasDiscount = true
		} else if discountVal, ok := product["discount"].(int); ok && discountVal > 0 {
			discount = float64(discountVal)
			hasDiscount = true
		} else if discountVal, ok := product["discount"].(int64); ok && discountVal > 0 {
			discount = float64(discountVal)
			hasDiscount = true
		}
		filtered["discount"] = discount
		filtered["hasDiscount"] = hasDiscount

		filteredProducts = append(filteredProducts, filtered)
	}

	return httpapi.ResSuccess(c, filteredProducts)
}
