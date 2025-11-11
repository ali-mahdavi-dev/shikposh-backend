package query

import (
	"context"
	"fmt"
	"strconv"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/entity"
	elasticsearchx "shikposh-backend/pkg/framework/infrastructure/elasticsearch"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type ProductQueryHandler struct {
	uow           unit_of_work.PGUnitOfWork
	elasticsearch elasticsearchx.Connection
	indexName     string
}

func NewProductQueryHandler(uow unit_of_work.PGUnitOfWork, elasticsearch elasticsearchx.Connection) *ProductQueryHandler {
	return &ProductQueryHandler{
		uow:           uow,
		elasticsearch: elasticsearch,
		indexName:     "products",
	}
}

func (h *ProductQueryHandler) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).GetAll(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) GetProductByID(ctx context.Context, id uint64) (*entity.Product, error) {
	// Try Elasticsearch first if available
	if h.elasticsearch != nil {
		productID := strconv.FormatUint(id, 10)
		doc, err := h.elasticsearch.GetDocument(ctx, h.indexName, productID)
		if err == nil {
			// Extract _source from Elasticsearch response
			source, ok := doc["_source"].(map[string]interface{})
			if !ok {
				// If _source doesn't exist, try using the doc itself
				source = doc
			}

			// Convert Elasticsearch document to Product entity
			product, err := h.mapToProduct(ctx, source)
			if err == nil {
				logging.Debug("Product retrieved from Elasticsearch").
					WithInt64("product_id", int64(id)).
					Log()
				return product, nil
			}
			logging.Warn("Failed to convert Elasticsearch document to product, falling back to database").
				WithInt64("product_id", int64(id)).
				WithError(err).
				Log()
		} else {
			logging.Debug("Product not found in Elasticsearch, falling back to database").
				WithInt64("product_id", int64(id)).
				Log()
		}
	}

	// Fallback to database
	var product *entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindByID(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return product, err
}

func (h *ProductQueryHandler) GetProductBySlug(ctx context.Context, slug string) (*entity.Product, error) {
	var product *entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindBySlug(ctx, slug)
		if err != nil {
			return err
		}
		return nil
	})
	return product, err
}

func (h *ProductQueryHandler) GetFeaturedProducts(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).FindFeatured(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) GetProductsByCategory(ctx context.Context, categorySlug string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).FindByCategorySlug(ctx, categorySlug)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) SearchProducts(ctx context.Context, searchQuery string) ([]*entity.Product, error) {
	// Try Elasticsearch first if available
	if h.elasticsearch != nil {
		products, err := h.searchInElasticsearch(ctx, searchQuery)
		if err == nil {
			logging.Debug("Products searched from Elasticsearch").
				WithString("query", searchQuery).
				WithInt("count", len(products)).
				Log()
			return products, nil
		}
		logging.Warn("Elasticsearch search failed, falling back to database").
			WithString("query", searchQuery).
			WithError(err).
			Log()
	}

	// Fallback to database
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).Search(ctx, searchQuery)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

func (h *ProductQueryHandler) GetFilteredProducts(ctx context.Context, filters repository.ProductFilters) ([]*entity.Product, error) {
	// Try Elasticsearch first if available
	if h.elasticsearch != nil {
		products, err := h.searchInElasticsearchWithFilters(ctx, filters)
		if err == nil {
			logging.Debug("Products filtered from Elasticsearch").
				WithInt("count", len(products)).
				Log()
			return products, nil
		}
		logging.Warn("Elasticsearch search failed, falling back to database").
			WithError(err).
			Log()
	}

	// Fallback to database
	var products []*entity.Product
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		products, err = h.uow.Product(ctx).Filter(ctx, filters)
		if err != nil {
			return err
		}
		return nil
	})
	return products, err
}

// searchInElasticsearch performs a search query in Elasticsearch
func (h *ProductQueryHandler) searchInElasticsearch(ctx context.Context, query string) ([]*entity.Product, error) {
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

// searchInElasticsearchWithFilters performs a search with all filters applied in Elasticsearch
func (h *ProductQueryHandler) searchInElasticsearchWithFilters(ctx context.Context, filters repository.ProductFilters) ([]*entity.Product, error) {
	// Build bool query with must, should, and filter clauses
	boolQuery := map[string]interface{}{
		"must":   []interface{}{},
		"should": []interface{}{},
		"filter": []interface{}{},
	}

	// Add search query if provided
	if filters.Query != nil && *filters.Query != "" {
		boolQuery["must"] = append(boolQuery["must"].([]interface{}), map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     *filters.Query,
				"fields":    []string{"name^3", "description^2", "brand"},
				"type":      "best_fields",
				"fuzziness": "AUTO",
			},
		})
	}

	// Add category filter
	if filters.Category != nil && *filters.Category != "" {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"category": *filters.Category,
			},
		})
	}

	// Add price range filter
	if filters.MinPrice != nil || filters.MaxPrice != nil {
		priceRange := map[string]interface{}{}
		if filters.MinPrice != nil {
			priceRange["gte"] = *filters.MinPrice
		}
		if filters.MaxPrice != nil {
			priceRange["lte"] = *filters.MaxPrice
		}
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"price": priceRange,
			},
		})
	}

	// Add rating filter
	if filters.Rating != nil {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"rating": map[string]interface{}{
					"gte": *filters.Rating,
				},
			},
		})
	}

	// Add featured filter
	if filters.Featured != nil && *filters.Featured {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"is_featured": true,
			},
		})
	}

	// Add tags filter
	if len(filters.Tags) > 0 {
		tagQueries := make([]interface{}, 0, len(filters.Tags))
		for _, tag := range filters.Tags {
			tagQueries = append(tagQueries, map[string]interface{}{
				"term": map[string]interface{}{
					"tags": tag,
				},
			})
		}
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"bool": map[string]interface{}{
				"should":               tagQueries,
				"minimum_should_match": 1,
			},
		})
	}

	// Build the final query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": boolQuery,
		},
		"size": 100,
	}

	// Add sorting if provided
	if filters.Sort != nil {
		sort := h.buildSortClause(*filters.Sort)
		if len(sort) > 0 {
			query["sort"] = sort
		}
	}

	return h.executeElasticsearchQuery(ctx, query)
}

// buildSortClause builds Elasticsearch sort clause
func (h *ProductQueryHandler) buildSortClause(sort string) []map[string]interface{} {
	switch sort {
	case "price_asc":
		return []map[string]interface{}{
			{"price": map[string]interface{}{"order": "asc"}},
		}
	case "price_desc":
		return []map[string]interface{}{
			{"price": map[string]interface{}{"order": "desc"}},
		}
	case "rating":
		return []map[string]interface{}{
			{"rating": map[string]interface{}{"order": "desc"}},
		}
	case "newest":
		return []map[string]interface{}{
			{"created_at": map[string]interface{}{"order": "desc"}},
		}
	default:
		// Default: relevance score (no sort clause needed)
		return []map[string]interface{}{}
	}
}

// executeElasticsearchQuery executes the Elasticsearch query and converts results to products
func (h *ProductQueryHandler) executeElasticsearchQuery(ctx context.Context, query map[string]interface{}) ([]*entity.Product, error) {
	result, err := h.elasticsearch.Search(ctx, h.indexName, query)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	// Extract hits from result
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid elasticsearch response format")
	}

	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid hits format")
	}

	products := make([]*entity.Product, 0, len(hitsArray))
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}

		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		product, err := h.mapToProduct(ctx, source)
		if err != nil {
			logging.Warn("Failed to convert Elasticsearch hit to product").
				WithError(err).
				Log()
			continue
		}

		products = append(products, product)
	}

	return products, nil
}

// mapToProduct converts a map (from Elasticsearch) to Product entity
func (h *ProductQueryHandler) mapToProduct(ctx context.Context, data map[string]interface{}) (*entity.Product, error) {
	// Get product ID
	idStr, ok := data["id"].(string)
	if !ok {
		return nil, fmt.Errorf("product id is missing or invalid")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse product id: %w", err)
	}

	// Get product from database to get full entity with relationships
	var product *entity.Product
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindByID(ctx, id)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get product from database: %w", err)
	}

	return product, nil
}
