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

func reindexCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reindex",
		Short: "reindex all products in Elasticsearch",
		RunE: func(_ *cobra.Command, _ []string) error {
			initializeConfigs()
			return reindexProducts()
		},
	}

	return cmd
}

func reindexProducts() error {
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

	// Create index with mapping if it doesn't exist
	mapping := map[string]interface{}{
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type": "keyword",
			},
			"seller_id": map[string]interface{}{
				"type": "integer",
			},
			"brand": map[string]interface{}{
				"type": "text",
			},
			"title": map[string]interface{}{
				"type": "text",
			},
			"slug": map[string]interface{}{
				"type": "keyword",
			},
			"description": map[string]interface{}{
				"type": "text",
			},
			"thumbnail": map[string]interface{}{
				"type": "keyword",
			},
			"discount": map[string]interface{}{
				"type": "integer",
			},
			"stock": map[string]interface{}{
				"type": "integer",
			},
			"original_price": map[string]interface{}{
				"type": "integer",
			},
			"price": map[string]interface{}{
				"type": "integer",
			},
			"rating": map[string]interface{}{
				"type": "float",
			},
			"is_featured": map[string]interface{}{
				"type": "boolean",
			},
			"is_new": map[string]interface{}{
				"type": "boolean",
			},
			"created_at": map[string]interface{}{
				"type": "date",
			},
			"categories": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"name": map[string]interface{}{
						"type": "text",
					},
				},
			},
			"colors": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"name": map[string]interface{}{
						"type": "text",
					},
					"hex": map[string]interface{}{
						"type": "keyword",
					},
				},
			},
			"sizes": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "integer",
					},
					"name": map[string]interface{}{
						"type": "keyword",
					},
				},
			},
			"tags": map[string]interface{}{
				"type": "keyword",
			},
			"features": map[string]interface{}{
				"type": "text",
			},
			"specs": map[string]interface{}{
				"type": "nested",
				"properties": map[string]interface{}{
					"key": map[string]interface{}{
						"type": "keyword",
					},
					"value": map[string]interface{}{
						"type": "text",
					},
				},
			},
			"images": map[string]interface{}{
				"type": "object",
			},
			"variant": map[string]interface{}{
				"type": "object",
			},
		},
	}

	if err := elasticsearch.CreateIndex(ctx, indexName, mapping); err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}

	logging.Info("Index created or already exists").WithString("index", indexName).Log()

	successCount := 0
	errorCount := 0

	// Get all products from database and index them
	err = uow.Do(ctx, func(ctx context.Context) error {
		productRepo := uow.Product(ctx)
		// Use FindAllForReindex to get all products without Images preload
		// This avoids the preload issue with Images
		allProducts, err := productRepo.FindAllForReindex(ctx)
		if err != nil {
			return fmt.Errorf("failed to get all products: %w", err)
		}

		logging.Info("Starting reindex process for all products").
			WithInt("total_products", len(allProducts)).
			Log()

		for _, product := range allProducts {
			productMap := product.ToMap()
			productIDStr := strconv.FormatUint(uint64(product.ID), 10)

			if err := elasticsearch.IndexDocument(ctx, indexName, productIDStr, productMap); err != nil {
				errorCount++
				logging.Error("Failed to index product").
					WithInt64("product_id", int64(product.ID)).
					WithError(err).
					Log()
				continue
			}

			successCount++
			if successCount%10 == 0 {
				logging.Info("Reindexing progress").
					WithInt("success", successCount).
					WithInt("errors", errorCount).
					Log()
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	logging.Info("Reindex completed").
		WithInt("success", successCount).
		WithInt("errors", errorCount).
		Log()

	return nil
}
