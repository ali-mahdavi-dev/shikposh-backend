package query

import (
	"context"
	"fmt"
	"shikposh-backend/internal/products/adapter/repository"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	unitofwork "shikposh-backend/internal/unit_of_work"
	"strconv"

	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

type ProductQueryHandler struct {
	uow           unitofwork.PGUnitOfWork
	elasticsearch elasticsearchx.Connection
	indexName     string
}

func NewProductQueryHandler(uow unitofwork.PGUnitOfWork, elasticsearch elasticsearchx.Connection) *ProductQueryHandler {
	return &ProductQueryHandler{
		uow:           uow,
		elasticsearch: elasticsearch,
		indexName:     "products",
	}
}

type CategoryQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewCategoryQueryHandler(uow unitofwork.PGUnitOfWork) *CategoryQueryHandler {
	return &CategoryQueryHandler{uow: uow}
}

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

// buildElasticsearchQueryWithFilters builds Elasticsearch query from filters
func (h *ProductQueryHandler) buildElasticsearchQueryWithFilters(filters repository.ProductFilters, categorySlug *string) map[string]interface{} {
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

	// Add category filter using nested query on categories.slug
	// Use term query with .keyword for exact match on keyword fields (slug is defined as keyword type)
	if categorySlug != nil && *categorySlug != "" {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "categories",
				"query": map[string]interface{}{
					"term": map[string]interface{}{
						"categories.slug.keyword": *categorySlug,
					},
				},
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

	return query
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
	case "discount_desc":
		return []map[string]interface{}{
			{"discount": map[string]interface{}{"order": "desc"}},
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

// searchInElasticsearchWithFilters performs a search with all filters applied in Elasticsearch
func (h *ProductQueryHandler) searchInElasticsearchWithFilters(ctx context.Context, filters repository.ProductFilters, categorySlug *string) ([]*productaggregate.Product, error) {
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

	// Add category filter using nested query on categories.slug
	// Use term query with .keyword for exact match on keyword fields (slug is defined as keyword type)
	if categorySlug != nil && *categorySlug != "" {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "categories",
				"query": map[string]interface{}{
					"term": map[string]interface{}{
						"categories.slug.keyword": *categorySlug,
					},
				},
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

	// Determine size
	size := 100
	if filters.Limit != nil && *filters.Limit > 0 {
		size = *filters.Limit
	}

	// Build the final query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": boolQuery,
		},
		"size": size,
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

// searchInElasticsearch performs a search query in Elasticsearch
func (h *ProductQueryHandler) searchInElasticsearch(ctx context.Context, query string) ([]*productaggregate.Product, error) {
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

// mapToProduct converts a map (from Elasticsearch) to Product entity
func (h *ProductQueryHandler) mapToProduct(ctx context.Context, data map[string]interface{}) (*productaggregate.Product, error) {
	if data == nil {
		return nil, fmt.Errorf("product data cannot be nil")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("product data is empty")
	}

	// Get product ID
	idValue, exists := data["id"]
	if !exists {
		return nil, fmt.Errorf("product id field is missing from elasticsearch document")
	}

	idStr, ok := idValue.(string)
	if !ok {
		// Try to convert from number if it's stored as a number
		if idNum, ok := idValue.(float64); ok {
			idStr = strconv.FormatFloat(idNum, 'f', 0, 64)
		} else {
			return nil, fmt.Errorf("product id has invalid type: expected string or number, got %T", idValue)
		}
	}

	if idStr == "" {
		return nil, fmt.Errorf("product id cannot be empty")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse product id (value=%s): %w", idStr, err)
	}

	if id == 0 {
		return nil, fmt.Errorf("product id cannot be zero")
	}

	// Get product from database to get full entity with relationships
	var product *productaggregate.Product
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindByID(ctx, id)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get product from database (id=%d): %w", id, err)
	}

	if product == nil {
		return nil, fmt.Errorf("product not found in database (id=%d)", id)
	}

	return product, nil
}
