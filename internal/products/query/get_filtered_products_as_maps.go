package query

import (
	"context"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetFilteredProductsAsMaps returns filtered products from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetFilteredProductsAsMaps(ctx context.Context, filters repository.ProductFilters) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	// Validate price range if both min and max are provided
	if filters.MinPrice != nil && filters.MaxPrice != nil {
		if *filters.MinPrice > *filters.MaxPrice {
			return nil, fmt.Errorf("invalid price range: min_price (%d) cannot be greater than max_price (%d)", *filters.MinPrice, *filters.MaxPrice)
		}
		if *filters.MinPrice < 0 || *filters.MaxPrice < 0 {
			return nil, fmt.Errorf("invalid price range: prices cannot be negative")
		}
	}

	// Validate rating if provided
	if filters.Rating != nil {
		if *filters.Rating < 0 || *filters.Rating > 5 {
			return nil, fmt.Errorf("invalid rating: rating must be between 0 and 5, got %.2f", *filters.Rating)
		}
	}

	// Get category slug for Elasticsearch nested query
	var categorySlug *string
	if filters.Category != nil && *filters.Category != "" {
		// Verify category exists in database before searching
		err := h.uow.Do(ctx, func(ctx context.Context) error {
			category, err := h.uow.Category(ctx).FindBySlug(ctx, *filters.Category)
			if err != nil {
				return err
			}
			if category == nil {
				return fmt.Errorf("category not found with slug: %s", *filters.Category)
			}
			categorySlug = filters.Category
			return nil
		})
		if err != nil {
			logging.Warn("Category not found in database").
				WithString("category_slug", *filters.Category).
				WithError(err).
				Log()
			return nil, fmt.Errorf("category not found (slug=%s): %w", *filters.Category, err)
		}
	} else if filters.CategoryName != nil && *filters.CategoryName != "" {
		// Convert category name to slug
		err := h.uow.Do(ctx, func(ctx context.Context) error {
			category, err := h.uow.Category(ctx).FindByName(ctx, *filters.CategoryName)
			if err != nil {
				return err
			}
			categorySlug = &category.Slug
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to convert category name to slug (name=%s): %w", *filters.CategoryName, err)
		}
	}

	query := h.buildElasticsearchQueryWithFilters(filters, categorySlug)
	maps, err := h.executeElasticsearchQueryAsMaps(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve filtered products as maps from elasticsearch: %w", err)
	}

	logging.Debug("Products filtered from Elasticsearch as maps").
		WithInt("count", len(maps)).
		WithString("category_slug", func() string {
			if categorySlug != nil {
				return *categorySlug
			}
			return ""
		}()).
		Log()

	// Log warning if category filter was used but no products found
	if categorySlug != nil && len(maps) == 0 {
		logging.Warn("No products found for category in Elasticsearch").
			WithString("category_slug", *categorySlug).
			Log()
	}

	return maps, nil
}
