package query

import (
	"shikposh-backend/internal/products/adapter/repository"
)

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
