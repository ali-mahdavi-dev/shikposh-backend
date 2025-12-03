package query

import (
	"context"
	"fmt"

	"shikposh-backend/internal/product/adapter/repository"
	productaggregate "shikposh-backend/internal/product/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

// GetProductBySlugAsMap returns product from Elasticsearch as map, with fallback to database
func (h *ProductQueryHandler) GetProductBySlugAsMap(ctx context.Context, slug string) (map[string]interface{}, error) {
	if slug == "" {
		return nil, fmt.Errorf("product slug cannot be empty")
	}

	// Try Elasticsearch first if available
	if h.elasticsearch != nil {
		searchQuery := map[string]interface{}{
			"query": map[string]interface{}{
				"term": map[string]interface{}{
					"slug": slug,
				},
			},
			"size": 1,
		}

		maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
		if err == nil && len(maps) > 0 {
			logging.Debug("Product retrieved from Elasticsearch by slug as map").
				WithString("slug", slug).
				Log()
			return maps[0], nil
		}

		// Log Elasticsearch error but continue to database fallback
		if err != nil {
			logging.Warn("Elasticsearch query failed, falling back to database").
				WithString("slug", slug).
				WithError(err).
				Log()
		} else {
			logging.Debug("Product not found in Elasticsearch, falling back to database").
				WithString("slug", slug).
				Log()
		}
	} else {
		logging.Debug("Elasticsearch not available, using database").
			WithString("slug", slug).
			Log()
	}

	// Fallback to database
	var product *productaggregate.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindBySlug(ctx, slug)
		return err
	})

	if err != nil {
		if err == repository.ErrProductNotFound {
			return nil, repository.ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to get product from database (slug=%s): %w", slug, err)
	}

	if product == nil {
		return nil, repository.ErrProductNotFound
	}

	logging.Debug("Product retrieved from database by slug as map").
		WithString("slug", slug).
		Log()

	return product.ToMap(), nil
}
