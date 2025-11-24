package query

import (
	"context"
	"fmt"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *ProductQueryHandler) GetProductsByCategory(ctx context.Context, categorySlug string) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	if categorySlug == "" {
		return nil, fmt.Errorf("category slug cannot be empty")
	}

	// First, get category ID from slug
	var categoryID uint64
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		category, err := h.uow.Category(ctx).FindBySlug(ctx, categorySlug)
		if err != nil {
			return err
		}
		categoryID = uint64(category.ID)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get category ID from slug (slug=%s): %w", categorySlug, err)
	}

	if categoryID == 0 {
		return nil, fmt.Errorf("invalid category ID retrieved for slug=%s", categorySlug)
	}

	// Search for products by category_id in Elasticsearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"category_id": categoryID,
			},
		},
		"size": 100,
	}

	products, err := h.executeElasticsearchQuery(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products by category from elasticsearch (category_slug=%s, category_id=%d): %w", categorySlug, categoryID, err)
	}

	logging.Debug("Products by category retrieved from Elasticsearch").
		WithString("category", categorySlug).
		WithInt64("category_id", int64(categoryID)).
		WithInt("count", len(products)).
		Log()

	return products, nil
}
