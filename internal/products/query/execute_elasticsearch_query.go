package query

import (
	"context"
	"fmt"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// executeElasticsearchQuery executes the Elasticsearch query and converts results to products
func (h *ProductQueryHandler) executeElasticsearchQuery(ctx context.Context, query map[string]interface{}) ([]*productaggregate.Product, error) {
	if query == nil {
		return nil, fmt.Errorf("elasticsearch query cannot be nil")
	}

	result, err := h.elasticsearch.Search(ctx, h.indexName, query)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search operation failed (index=%s): %w", h.indexName, err)
	}

	if result == nil {
		return nil, fmt.Errorf("empty response from elasticsearch (index=%s)", h.indexName)
	}

	// Extract hits from result
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid elasticsearch response format: missing or invalid 'hits' field (index=%s)", h.indexName)
	}

	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid elasticsearch response format: missing or invalid 'hits.hits' array (index=%s)", h.indexName)
	}

	products := make([]*productaggregate.Product, 0, len(hitsArray))
	var conversionErrors []error

	for i, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			logging.Warn("Skipping invalid hit format in elasticsearch response").
				WithInt("hit_index", i).
				Log()
			continue
		}

		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			logging.Warn("Skipping hit with missing _source field").
				WithInt("hit_index", i).
				Log()
			continue
		}

		product, err := h.mapToProduct(ctx, source)
		if err != nil {
			conversionErrors = append(conversionErrors, fmt.Errorf("hit at index %d: %w", i, err))
			logging.Warn("Failed to convert Elasticsearch hit to product").
				WithError(err).
				WithInt("hit_index", i).
				Log()
			continue
		}

		products = append(products, product)
	}

	// If we had conversion errors but no successful conversions, return an error
	if len(conversionErrors) > 0 && len(products) == 0 {
		return nil, fmt.Errorf("failed to convert any elasticsearch hits to products: %v", conversionErrors[0])
	}

	return products, nil
}

