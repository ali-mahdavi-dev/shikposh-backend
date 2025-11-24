package query

import (
	"context"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
)

// searchInElasticsearch performs a search query in Elasticsearch
func (h *ProductQueryHandler) searchInElasticsearch(ctx context.Context, query string) ([]*productaggregate.Product, error) {
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     query,
				"fields":    []string{"name^3", "description^2", "brand"},
				"type":      "best_fields",
				"fuzziness": "AUTO",
			},
		},
		"size": 100,
	}

	return h.executeElasticsearchQuery(ctx, searchQuery)
}

