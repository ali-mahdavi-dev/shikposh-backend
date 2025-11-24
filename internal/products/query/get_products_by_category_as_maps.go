package query

import (
	"context"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetProductsByCategoryAsMaps returns products by category from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetProductsByCategoryAsMaps(ctx context.Context, categorySlug string) ([]map[string]interface{}, error) {
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

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products by category as maps from elasticsearch (category_slug=%s, category_id=%d): %w", categorySlug, categoryID, err)
	}

	logging.Debug("Products by category retrieved from Elasticsearch as maps").
		WithString("category", categorySlug).
		WithInt64("category_id", int64(categoryID)).
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}
