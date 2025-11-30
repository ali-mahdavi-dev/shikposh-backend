package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/spf13/cobra"
)

func testCategoryQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-category-query",
		Short: "Test category query in Elasticsearch",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			initializeConfigs()
			categorySlug := args[0]
			return testCategoryQuery(categorySlug)
		},
	}

	return cmd
}

func testCategoryQuery(categorySlug string) error {
	// Initialize Elasticsearch connection
	elasticsearch, err := initializeElasticsearch(&cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize elasticsearch: %w", err)
	}

	if elasticsearch == nil {
		return fmt.Errorf("elasticsearch is not available")
	}

	ctx := context.Background()
	indexName := "products"

	logging.Info("Testing category query in Elasticsearch").
		WithString("category_slug", categorySlug).
		WithString("index", indexName).
		Log()

	// Test 1: term query (exact match, case-sensitive)
	searchQuery1 := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"nested": map[string]interface{}{
							"path": "categories",
							"query": map[string]interface{}{
								"term": map[string]interface{}{
									"categories.slug": categorySlug,
								},
							},
						},
					},
				},
			},
		},
		"size": 100,
	}

	// Test 2: match query
	searchQuery2 := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"nested": map[string]interface{}{
							"path": "categories",
							"query": map[string]interface{}{
								"match": map[string]interface{}{
									"categories.slug": categorySlug,
								},
							},
						},
					},
				},
			},
		},
		"size": 100,
	}

	// Test 3: term query with keyword field explicitly
	searchQuery3 := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"nested": map[string]interface{}{
							"path": "categories",
							"query": map[string]interface{}{
								"term": map[string]interface{}{
									"categories.slug.keyword": categorySlug,
								},
							},
						},
					},
				},
			},
		},
		"size": 100,
	}

	queries := []map[string]interface{}{searchQuery1, searchQuery2, searchQuery3}
	queryNames := []string{"term query", "match query", "term query with .keyword"}

	for idx, searchQuery := range queries {
		fmt.Printf("\n=== Test %d: %s ===\n", idx+1, queryNames[idx])

		// Pretty print the query
		queryJSON, _ := json.MarshalIndent(searchQuery, "", "  ")
		fmt.Println("Query:")
		fmt.Println(string(queryJSON))
		fmt.Println()

		searchResult, err := elasticsearch.Search(ctx, indexName, searchQuery)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Extract hits
		hits, ok := searchResult["hits"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid response format")
			continue
		}

		total, _ := hits["total"].(map[string]interface{})
		totalValue := 0
		if total != nil {
			if val, ok := total["value"].(float64); ok {
				totalValue = int(val)
			}
		}

		hitsArray, ok := hits["hits"].([]interface{})
		if !ok {
			fmt.Println("Invalid hits format")
			continue
		}

		fmt.Printf("Total results: %d\n", totalValue)
		fmt.Printf("Returned hits: %d\n", len(hitsArray))

		if len(hitsArray) > 0 {
			fmt.Println("✓ Found products!")
			// Show first result
			hitMap, ok := hitsArray[0].(map[string]interface{})
			if ok {
				source, ok := hitMap["_source"].(map[string]interface{})
				if ok {
					id := "unknown"
					if idVal, ok := source["id"]; ok {
						id = fmt.Sprintf("%v", idVal)
					}
					title := "unknown"
					if titleVal, ok := source["title"].(string); ok {
						title = titleVal
					}
					fmt.Printf("First product: ID=%s, Title=%s\n", id, title)
				}
			}
		} else {
			fmt.Println("✗ No products found")
		}
	}

	return nil
}
