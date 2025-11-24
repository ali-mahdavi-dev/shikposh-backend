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
			return nil, fmt.Errorf("invalid price range: min_price (%.2f) cannot be greater than max_price (%.2f)", *filters.MinPrice, *filters.MaxPrice)
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

	// Convert category slug to category_id if needed for Elasticsearch
	var categoryID *uint64
	if filters.Category != nil && *filters.Category != "" {
		err := h.uow.Do(ctx, func(ctx context.Context) error {
			category, err := h.uow.Category(ctx).FindBySlug(ctx, *filters.Category)
			if err != nil {
				return err
			}
			id := uint64(category.ID)
			categoryID = &id
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to convert category slug to ID (slug=%s): %w", *filters.Category, err)
		}
		if categoryID != nil && *categoryID == 0 {
			return nil, fmt.Errorf("invalid category ID retrieved for slug=%s", *filters.Category)
		}
	}

	query := h.buildElasticsearchQueryWithFilters(filters, categoryID)
	maps, err := h.executeElasticsearchQueryAsMaps(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve filtered products as maps from elasticsearch: %w", err)
	}

	logging.Debug("Products filtered from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}
