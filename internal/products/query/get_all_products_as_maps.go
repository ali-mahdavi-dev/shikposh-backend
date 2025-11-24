package query

import (
	"context"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetAllProductsAsMaps returns products from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetAllProductsAsMaps(ctx context.Context) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 100,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all products as maps from elasticsearch: %w", err)
	}

	logging.Debug("All products retrieved from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}

