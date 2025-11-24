package query

import (
	"context"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetProductBySlugAsMap returns product from Elasticsearch as map (no database lookup)
func (h *ProductQueryHandler) GetProductBySlugAsMap(ctx context.Context, slug string) (map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	if slug == "" {
		return nil, fmt.Errorf("product slug cannot be empty")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"slug": slug,
			},
		},
		"size": 1,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to search product by slug from elasticsearch (slug=%s): %w", slug, err)
	}

	if len(maps) == 0 {
		return nil, repository.ErrProductNotFound
	}

	logging.Debug("Product retrieved from Elasticsearch by slug as map").
		WithString("slug", slug).
		Log()

	return maps[0], nil
}
