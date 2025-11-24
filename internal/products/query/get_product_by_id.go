package query

import (
	"context"
	"fmt"
	"strconv"

	"shikposh-backend/internal/products/adapter/repository"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

func (h *ProductQueryHandler) GetProductByID(ctx context.Context, id uint64) (*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch connection is not initialized")
	}

	if id == 0 {
		return nil, fmt.Errorf("product id cannot be zero")
	}

	productID := strconv.FormatUint(id, 10)
	doc, err := h.elasticsearch.GetDocument(ctx, h.indexName, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product document from elasticsearch (id=%d, index=%s): %w", id, h.indexName, err)
	}

	if doc == nil {
		return nil, repository.ErrProductNotFound
	}

	// Extract _source from Elasticsearch response
	source, ok := doc["_source"].(map[string]interface{})
	if !ok {
		// If _source doesn't exist, try using the doc itself
		if len(doc) == 0 {
			return nil, fmt.Errorf("empty document returned from elasticsearch for product id=%d", id)
		}
		source = doc
	}

	// Convert Elasticsearch document to Product entity
	product, err := h.mapToProduct(ctx, source)
	if err != nil {
		return nil, fmt.Errorf("failed to convert elasticsearch document to product (id=%d): %w", id, err)
	}

	logging.Debug("Product retrieved from Elasticsearch").
		WithInt64("product_id", int64(id)).
		Log()

	return product, nil
}
