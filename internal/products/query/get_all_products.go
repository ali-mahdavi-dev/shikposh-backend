package query

import (
	"context"
	"fmt"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *ProductQueryHandler) GetAllProducts(ctx context.Context) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	// Use match_all query to get all products
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 100,
	}

	products, err := h.executeElasticsearchQuery(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all products from elasticsearch: %w", err)
	}

	logging.Debug("All products retrieved from Elasticsearch").
		WithInt("count", len(products)).
		Log()

	return products, nil
}
