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

	successCount := 0
	errorCount := 0

	// Get all products from database and index them
	err = uow.Do(ctx, func(ctx context.Context) error {
		productRepo := uow.Product(ctx)
		allProducts, err := productRepo.GetAll(ctx)
		if err != nil {
			return fmt.Errorf("failed to get products: %w", err)
		}

		logging.Info("Starting reindex process").
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
