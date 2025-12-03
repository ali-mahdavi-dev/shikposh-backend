package query

import (
	"context"
	"fmt"

	"shikposh-backend/internal/product/adapter/repository"
	productaggregate "shikposh-backend/internal/product/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *ProductQueryHandler) GetProductBySlug(ctx context.Context, slug string) (*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	if slug == "" {
		return nil, fmt.Errorf("product slug cannot be empty")
	}

	// Search for product by slug in Elasticsearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"slug": slug,
			},
		},
		"size": 1,
	}

	products, err := h.executeElasticsearchQuery(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to search product by slug from elasticsearch (slug=%s): %w", slug, err)
	}

	if len(products) == 0 {
		return nil, repository.ErrProductNotFound
	}

	logging.Debug("Product retrieved from Elasticsearch by slug").
		WithString("slug", slug).
		Log()

	return products[0], nil
}
