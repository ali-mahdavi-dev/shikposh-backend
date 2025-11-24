package query

import (
	"context"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetProductsByIDsAsMaps returns products by IDs from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetProductsByIDsAsMaps(ctx context.Context, productIDs []string) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	if len(productIDs) == 0 {
		return []map[string]interface{}{}, nil
	}

	// Validate and filter out empty IDs
	validIDs := make([]interface{}, 0, len(productIDs))
	for i, id := range productIDs {
		if id == "" {
			logging.Warn("Skipping empty product ID in request").
				WithInt("index", i).
				Log()
			continue
		}
		validIDs = append(validIDs, id)
	}

	if len(validIDs) == 0 {
		return nil, fmt.Errorf("no valid product IDs provided")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"terms": map[string]interface{}{
				"id": validIDs,
			},
		},
		"size": len(validIDs),
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products by IDs from elasticsearch (count=%d): %w", len(validIDs), err)
	}

	logging.Debug("Products by IDs retrieved from Elasticsearch as maps").
		WithInt("requested_count", len(validIDs)).
		WithInt("returned_count", len(maps)).
		Log()

	return maps, nil
}
