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

	// Search for products by category slug in Elasticsearch using nested query
	// Categories is a nested field, so we need to use nested query
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "categories",
				"query": map[string]interface{}{
					"term": map[string]interface{}{
						"categories.slug": categorySlug,
					},
				},
			},
		},
		"size": 100,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products by category as maps from elasticsearch (category_slug=%s): %w", categorySlug, err)
	}

	logging.Debug("Products by category retrieved from Elasticsearch as maps").
		WithString("category", categorySlug).
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}
