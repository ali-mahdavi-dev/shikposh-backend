package query

import (
	"context"
	"fmt"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *ProductQueryHandler) SearchProducts(ctx context.Context, searchQuery string) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	if searchQuery == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	products, err := h.searchInElasticsearch(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to search products in elasticsearch (query=%s): %w", searchQuery, err)
	}

	logging.Debug("Products searched from Elasticsearch").
		WithString("query", searchQuery).
		WithInt("count", len(products)).
		Log()

	return products, nil
}
