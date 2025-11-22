package query

import (
	"context"
	"fmt"
	"strconv"

	"shikposh-backend/internal/products/adapter/repository"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	unitofwork "shikposh-backend/internal/unit_of_work"

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

func (h *ProductQueryHandler) GetAllProducts(ctx context.Context) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
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
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("All products retrieved from Elasticsearch").
		WithInt("count", len(products)).
		Log()

	return products, nil
}

// GetAllProductsAsMaps returns products from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetAllProductsAsMaps(ctx context.Context) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 100,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("All products retrieved from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}

func (h *ProductQueryHandler) GetProductByID(ctx context.Context, id uint64) (*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	productID := strconv.FormatUint(id, 10)
	doc, err := h.elasticsearch.GetDocument(ctx, h.indexName, productID)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch get document failed: %w", err)
	}

	// Extract _source from Elasticsearch response
	source, ok := doc["_source"].(map[string]interface{})
	if !ok {
		// If _source doesn't exist, try using the doc itself
		source = doc
	}

	// Convert Elasticsearch document to Product entity
	product, err := h.mapToProduct(ctx, source)
	if err != nil {
		return nil, fmt.Errorf("failed to convert Elasticsearch document to product: %w", err)
	}

	logging.Debug("Product retrieved from Elasticsearch").
		WithInt64("product_id", int64(id)).
		Log()

	return product, nil
}

// GetProductByIDAsMap returns product from Elasticsearch as map (no database lookup)
func (h *ProductQueryHandler) GetProductByIDAsMap(ctx context.Context, id uint64) (map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	productID := strconv.FormatUint(id, 10)
	doc, err := h.elasticsearch.GetDocument(ctx, h.indexName, productID)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch get document failed: %w", err)
	}

	// Extract _source from Elasticsearch response
	source, ok := doc["_source"].(map[string]interface{})
	if !ok {
		// If _source doesn't exist, try using the doc itself
		source = doc
	}

	logging.Debug("Product retrieved from Elasticsearch as map").
		WithInt64("product_id", int64(id)).
		Log()

	return source, nil
}

func (h *ProductQueryHandler) GetProductBySlug(ctx context.Context, slug string) (*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
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
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	if len(products) == 0 {
		return nil, repository.ErrProductNotFound
	}

	logging.Debug("Product retrieved from Elasticsearch by slug").
		WithString("slug", slug).
		Log()

	return products[0], nil
}

// GetProductBySlugAsMap returns product from Elasticsearch as map (no database lookup)
func (h *ProductQueryHandler) GetProductBySlugAsMap(ctx context.Context, slug string) (map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"slug": slug,
			},
		},
		"size": 1,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	if len(maps) == 0 {
		return nil, repository.ErrProductNotFound
	}

	logging.Debug("Product retrieved from Elasticsearch by slug as map").
		WithString("slug", slug).
		Log()

	return maps[0], nil
}

func (h *ProductQueryHandler) GetFeaturedProducts(ctx context.Context) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	// Search for featured products in Elasticsearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"is_featured": true,
			},
		},
		"size": 100,
	}

	products, err := h.executeElasticsearchQuery(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Featured products retrieved from Elasticsearch").
		WithInt("count", len(products)).
		Log()

	return products, nil
}

// GetFeaturedProductsAsMaps returns featured products from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetFeaturedProductsAsMaps(ctx context.Context) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"is_featured": true,
			},
		},
		"size": 100,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Featured products retrieved from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}

func (h *ProductQueryHandler) GetProductsByCategory(ctx context.Context, categorySlug string) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	// First, get category ID from slug
	var categoryID uint64
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		category, err := h.uow.Category(ctx).FindBySlug(ctx, categorySlug)
		if err != nil {
			return err
		}
		categoryID = uint64(category.ID)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get category ID from slug: %w", err)
	}

	// Search for products by category_id in Elasticsearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"category_id": categoryID,
			},
		},
		"size": 100,
	}

	products, err := h.executeElasticsearchQuery(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Products by category retrieved from Elasticsearch").
		WithString("category", categorySlug).
		WithInt64("category_id", int64(categoryID)).
		WithInt("count", len(products)).
		Log()

	return products, nil
}

// GetProductsByCategoryAsMaps returns products by category from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetProductsByCategoryAsMaps(ctx context.Context, categorySlug string) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	// First, get category ID from slug
	var categoryID uint64
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		category, err := h.uow.Category(ctx).FindBySlug(ctx, categorySlug)
		if err != nil {
			return err
		}
		categoryID = uint64(category.ID)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get category ID from slug: %w", err)
	}

	// Search for products by category_id in Elasticsearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"category_id": categoryID,
			},
		},
		"size": 100,
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Products by category retrieved from Elasticsearch as maps").
		WithString("category", categorySlug).
		WithInt64("category_id", int64(categoryID)).
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}

func (h *ProductQueryHandler) SearchProducts(ctx context.Context, searchQuery string) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	products, err := h.searchInElasticsearch(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Products searched from Elasticsearch").
		WithString("query", searchQuery).
		WithInt("count", len(products)).
		Log()

	return products, nil
}

func (h *ProductQueryHandler) GetFilteredProducts(ctx context.Context, filters repository.ProductFilters) ([]*productaggregate.Product, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	// Convert category slug to category_id if needed for Elasticsearch
	var categoryID *uint64
	if filters.Category != nil && *filters.Category != "" {
		err := h.uow.Do(ctx, func(ctx context.Context) error {
			category, err := h.uow.Category(ctx).FindBySlug(ctx, *filters.Category)
			if err != nil {
				return err
			}
			id := uint64(category.ID)
			categoryID = &id
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to convert category slug to ID: %w", err)
		}
	}

	products, err := h.searchInElasticsearchWithFilters(ctx, filters, categoryID)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Products filtered from Elasticsearch").
		WithInt("count", len(products)).
		Log()

	return products, nil
}

// GetFilteredProductsAsMaps returns filtered products from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetFilteredProductsAsMaps(ctx context.Context, filters repository.ProductFilters) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	// Convert category slug to category_id if needed for Elasticsearch
	var categoryID *uint64
	if filters.Category != nil && *filters.Category != "" {
		err := h.uow.Do(ctx, func(ctx context.Context) error {
			category, err := h.uow.Category(ctx).FindBySlug(ctx, *filters.Category)
			if err != nil {
				return err
			}
			id := uint64(category.ID)
			categoryID = &id
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to convert category slug to ID: %w", err)
		}
	}

	query := h.buildElasticsearchQueryWithFilters(filters, categoryID)
	maps, err := h.executeElasticsearchQueryAsMaps(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Products filtered from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}

// buildElasticsearchQueryWithFilters builds Elasticsearch query from filters
func (h *ProductQueryHandler) buildElasticsearchQueryWithFilters(filters repository.ProductFilters, categoryID *uint64) map[string]interface{} {
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
	if categoryID != nil {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"category_id": *categoryID,
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

// searchInElasticsearchWithFilters performs a search with all filters applied in Elasticsearch
func (h *ProductQueryHandler) searchInElasticsearchWithFilters(ctx context.Context, filters repository.ProductFilters, categoryID *uint64) ([]*productaggregate.Product, error) {
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

	// Add category filter using category_id (converted from slug in GetFilteredProducts)
	if categoryID != nil {
		boolQuery["filter"] = append(boolQuery["filter"].([]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"category_id": *categoryID,
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
func (h *ProductQueryHandler) executeElasticsearchQuery(ctx context.Context, query map[string]interface{}) ([]*productaggregate.Product, error) {
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

	products := make([]*productaggregate.Product, 0, len(hitsArray))
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

// executeElasticsearchQueryAsMaps executes the Elasticsearch query and returns results as maps (no database lookup)
func (h *ProductQueryHandler) executeElasticsearchQueryAsMaps(ctx context.Context, query map[string]interface{}) ([]map[string]interface{}, error) {
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

	maps := make([]map[string]interface{}, 0, len(hitsArray))
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}

		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		maps = append(maps, source)
	}

	return maps, nil
}

// GetProductsByIDsAsMaps returns products by IDs from Elasticsearch as maps (no database lookup)
func (h *ProductQueryHandler) GetProductsByIDsAsMaps(ctx context.Context, productIDs []string) ([]map[string]interface{}, error) {
	if h.elasticsearch == nil {
		return nil, fmt.Errorf("elasticsearch is not available")
	}

	if len(productIDs) == 0 {
		return []map[string]interface{}{}, nil
	}

	// Convert string IDs to interface slice for terms query
	ids := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		ids[i] = id
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"terms": map[string]interface{}{
				"id": ids,
			},
		},
		"size": len(productIDs),
	}

	maps, err := h.executeElasticsearchQueryAsMaps(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	logging.Debug("Products by IDs retrieved from Elasticsearch as maps").
		WithInt("count", len(maps)).
		Log()

	return maps, nil
}

// mapToProduct converts a map (from Elasticsearch) to Product entity
func (h *ProductQueryHandler) mapToProduct(ctx context.Context, data map[string]interface{}) (*productaggregate.Product, error) {
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
	var product *productaggregate.Product
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
