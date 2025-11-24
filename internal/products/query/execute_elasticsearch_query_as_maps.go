package query

import (
	"context"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// executeElasticsearchQueryAsMaps executes the Elasticsearch query and returns results as maps (no database lookup)
func (h *ProductQueryHandler) executeElasticsearchQueryAsMaps(ctx context.Context, query map[string]interface{}) ([]map[string]interface{}, error) {
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

	maps := make([]map[string]interface{}, 0, len(hitsArray))
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

		maps = append(maps, source)
	}

	return maps, nil
}

