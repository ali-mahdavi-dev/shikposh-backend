package query

import (
	"context"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetFeaturedProductsAsMaps returns featured products from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetFeaturedProductsAsMaps(ctx context.Context) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"is_featured": true,
			},
		},
		"size": 100,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve featured products as maps from elasticsearch: %w", err)
	}

	logging.Debug("Featured products retrieved from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}
