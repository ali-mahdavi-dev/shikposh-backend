package query

import (
	"context"
	"fmt"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *ProductQueryHandler) GetFeaturedProducts(ctx context.Context) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	// Search for featured products in Elasticsearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"is_featured": true,
			},
		},
		"size": 100,
	}

	products, err := h.executeElasticsearchQuery(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve featured products from elasticsearch: %w", err)
	}

	logging.Debug("Featured products retrieved from Elasticsearch").
		WithInt("count", len(products)).
		Log()

	return products, nil
}
