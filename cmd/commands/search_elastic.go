package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/spf13/cobra"
)

func searchElasticCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search-elastic",
		Short: "Search for a product in Elasticsearch by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			initializeConfigs()
			productID := args[0]
			return searchProductInElasticsearch(productID)
		},
	}

	return cmd
}

func searchProductInElasticsearch(productID string) error {
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

	// Try to get document by ID first
	logging.Info("Searching for product in Elasticsearch").
		WithString("product_id", productID).
		WithString("index", indexName).
		Log()

	// Method 1: Try GetDocument (direct document access)
	result, err := elasticsearch.GetDocument(ctx, indexName, productID)
	if err == nil && result != nil {
		logging.Info("Product found using GetDocument").
			WithString("product_id", productID).
			Log()

		// Extract _source from result
		if source, ok := result["_source"].(map[string]interface{}); ok {
			prettyJSON, _ := json.MarshalIndent(source, "", "  ")
			fmt.Println("Product found:")
			fmt.Println(string(prettyJSON))
			return nil
		}
	}

	// Method 2: Search by term query on id field
	logging.Info("Trying search query method").
		WithString("product_id", productID).
		Log()

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"id": productID,
			},
		},
		"size": 1,
	}

	searchResult, err := elasticsearch.Search(ctx, indexName, searchQuery)
	if err != nil {
		return fmt.Errorf("failed to search elasticsearch: %w", err)
	}

	// Extract hits
	hits, ok := searchResult["hits"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid elasticsearch response format: missing 'hits' field")
	}

	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid elasticsearch response format: missing 'hits.hits' array")
	}

	if len(hitsArray) == 0 {
		logging.Warn("Product not found in Elasticsearch").
			WithString("product_id", productID).
			Log()
		fmt.Printf("Product with ID '%s' not found in Elasticsearch index '%s'\n", productID, indexName)
		return nil
	}

	// Get first hit
	hitMap, ok := hitsArray[0].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid hit format")
	}

	source, ok := hitMap["_source"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("missing _source in hit")
	}

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(source, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal result: %w", err)
	}

	fmt.Println("Product found:")
	fmt.Println(string(prettyJSON))

	// Also show categories if they exist
	if categories, ok := source["categories"].([]interface{}); ok && len(categories) > 0 {
		fmt.Println("\nCategories:")
		for i, cat := range categories {
			if catMap, ok := cat.(map[string]interface{}); ok {
				fmt.Printf("  [%d] ", i+1)
				if slug, ok := catMap["slug"].(string); ok {
					fmt.Printf("slug: %s", slug)
				}
				if name, ok := catMap["name"].(string); ok {
					fmt.Printf(", name: %s", name)
				}
				if id, ok := catMap["id"].(float64); ok {
					fmt.Printf(", id: %.0f", id)
				}
				fmt.Println()
			}
		}
	}

	return nil
}
