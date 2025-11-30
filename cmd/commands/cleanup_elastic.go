package commands

import (
	"context"
	"fmt"
	"strconv"

	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/spf13/cobra"
)

func cleanupElasticCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup-elastic",
		Short: "Remove products from Elasticsearch that don't exist in database",
		RunE: func(_ *cobra.Command, _ []string) error {
			initializeConfigs()
			return cleanupElasticsearch()
		},
	}

	return cmd
}

func cleanupElasticsearch() error {
	// Initialize database connection
	db, err := initializeDatabase(&cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer closeDatabase(db)

	// Initialize Elasticsearch connection
	elasticsearch, err := initializeElasticsearch(&cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize elasticsearch: %w", err)
	}

	if elasticsearch == nil {
		return fmt.Errorf("elasticsearch is not available")
	}

	// Initialize unit of work
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)

	ctx := context.Background()
	indexName := "products"

	logging.Info("Starting cleanup process").
		WithString("index", indexName).
		Log()

	// Step 1: Get all product IDs from database
	var dbProductIDs map[uint64]bool
	err = uow.Do(ctx, func(ctx context.Context) error {
		productRepo := uow.Product(ctx)
		allProducts, err := productRepo.FindAllForReindex(ctx)
		if err != nil {
			return fmt.Errorf("failed to get all products from database: %w", err)
		}

		dbProductIDs = make(map[uint64]bool, len(allProducts))
		for _, product := range allProducts {
			dbProductIDs[uint64(product.ID)] = true
		}

		logging.Info("Products found in database").
			WithInt("count", len(dbProductIDs)).
			Log()

		return nil
	})

	if err != nil {
		return err
	}

	// Step 2: Get all document IDs from Elasticsearch
	// Use scroll API to get all documents
	elasticsearchIDs := make(map[string]bool)
	from := 0
	size := 1000

	for {
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
			"_source": false, // Don't return source, only IDs
			"from":    from,
			"size":    size,
		}

		result, err := elasticsearch.Search(ctx, indexName, query)
		if err != nil {
			return fmt.Errorf("failed to search elasticsearch: %w", err)
		}

		// Extract hits from result
		hits, ok := result["hits"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid elasticsearch response format")
		}

		hitsArray, ok := hits["hits"].([]interface{})
		if !ok {
			return fmt.Errorf("invalid elasticsearch response format: missing hits array")
		}

		// Extract IDs from hits
		for _, hit := range hitsArray {
			hitMap, ok := hit.(map[string]interface{})
			if !ok {
				continue
			}

			// Get _id from hit
			id, ok := hitMap["_id"].(string)
			if !ok {
				continue
			}

			elasticsearchIDs[id] = true
		}

		// Check if we have more results
		total, ok := hits["total"].(map[string]interface{})
		if !ok {
			// Try as number
			totalNum, ok := hits["total"].(float64)
			if !ok {
				break
			}
			if float64(from+len(hitsArray)) >= totalNum {
				break
			}
		} else {
			totalValue, ok := total["value"].(float64)
			if !ok {
				break
			}
			if float64(from+len(hitsArray)) >= totalValue {
				break
			}
		}

		if len(hitsArray) < size {
			break
		}

		from += size
	}

	logging.Info("Documents found in Elasticsearch").
		WithInt("count", len(elasticsearchIDs)).
		Log()

	// Step 3: Find IDs that are in Elasticsearch but not in database
	idsToDelete := make([]string, 0)
	for esID := range elasticsearchIDs {
		productID, err := strconv.ParseUint(esID, 10, 64)
		if err != nil {
			logging.Warn("Invalid product ID in Elasticsearch").
				WithString("id", esID).
				Log()
			// Still delete it as it's invalid
			idsToDelete = append(idsToDelete, esID)
			continue
		}

		if !dbProductIDs[productID] {
			idsToDelete = append(idsToDelete, esID)
		}
	}

	logging.Info("Products to delete from Elasticsearch").
		WithInt("count", len(idsToDelete)).
		Log()

	// Step 4: Delete products from Elasticsearch
	deletedCount := 0
	errorCount := 0

	for _, id := range idsToDelete {
		if err := elasticsearch.DeleteDocument(ctx, indexName, id); err != nil {
			errorCount++
			logging.Error("Failed to delete product from Elasticsearch").
				WithString("product_id", id).
				WithError(err).
				Log()
			continue
		}

		deletedCount++
		if deletedCount%10 == 0 {
			logging.Info("Cleanup progress").
				WithInt("deleted", deletedCount).
				WithInt("errors", errorCount).
				Log()
		}
	}

	logging.Info("Cleanup completed").
		WithInt("deleted", deletedCount).
		WithInt("errors", errorCount).
		Log()

	return nil
}

