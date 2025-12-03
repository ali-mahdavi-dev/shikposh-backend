package command

import (
	"context"
	"fmt"

	"shikposh-backend/internal/product/domain/entity"
	product_aggregate "shikposh-backend/internal/product/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func importFromElasticCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-elastic",
		Short: "Import products from Elasticsearch to database",
		RunE: func(_ *cobra.Command, _ []string) error {
			initializeConfigs()
			return importFromElastic()
		},
	}

	return cmd
}

func importFromElastic() error {
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

	ctx := context.Background()
	indexName := "products"

	// Search all documents from Elasticsearch
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 10000,
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

	logging.Info("Found documents in Elasticsearch").
		WithInt("count", len(hitsArray)).
		Log()

	successCount := 0
	errorCount := 0

	// Process each product
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		doc, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			product := &product_aggregate.Product{
				Title:       cast.ToString(doc["title"]),
				Slug:        cast.ToString(doc["slug"]),
				Brand:       cast.ToString(doc["brand"]),
				Description: strPtr(cast.ToString(doc["description"])),
				Thumbnail:   cast.ToString(doc["thumbnail"]),
				Discount:    cast.ToInt(doc["discount"]),
				Stock:       cast.ToInt(doc["stock"]),
				IsNew:       cast.ToBool(doc["is_new"]),
				IsFeatured:  cast.ToBool(doc["is_featured"]),
				Rating:      cast.ToFloat64(doc["rating"]),
				Price:       cast.ToInt64(doc["price"]),
			}

			// Handle discount as origin_price
			if discount := cast.ToInt64(doc["discount"]); discount > 0 {
				originPrice := product.Price * (100 + discount) / 100
				product.OriginPrice = &originPrice
			}

			// Handle categories - ensure category exists
			if categoriesRaw, ok := doc["categories"].([]interface{}); ok {
				for _, catRaw := range categoriesRaw {
					if catMap, ok := catRaw.(map[string]interface{}); ok {
						catSlug := cast.ToString(catMap["slug"])
						catName := cast.ToString(catMap["name"])
						if catSlug != "" {
							var cat entity.Category
							err := tx.Where("slug = ?", catSlug).First(&cat).Error
							if err == gorm.ErrRecordNotFound {
								cat = entity.Category{
									Name: catName,
									Slug: catSlug,
								}
								if err := tx.Create(&cat).Error; err != nil {
									logging.Warn("Failed to create category").
										WithString("slug", catSlug).
										WithError(err).
										Log()
								}
							}
							product.Categories = append(product.Categories, cat)
						}
					}
				}
			}

			// Handle colors
			if colorsRaw, ok := doc["colors"].([]interface{}); ok {
				for _, colorRaw := range colorsRaw {
					if colorMap, ok := colorRaw.(map[string]interface{}); ok {
						colorName := cast.ToString(colorMap["name"])
						colorHex := cast.ToString(colorMap["hex"])
						if colorName != "" {
							var color product_aggregate.Color
							err := tx.Where("name = ?", colorName).First(&color).Error
							if err == gorm.ErrRecordNotFound {
								color = product_aggregate.Color{
									Name: colorName,
									Hex:  colorHex,
								}
								if err := tx.Create(&color).Error; err != nil {
									continue
								}
							}
							product.Colors = append(product.Colors, color)
						}
					}
				}
			}

			// Handle sizes
			if sizesRaw, ok := doc["sizes"].([]interface{}); ok {
				for _, sizeRaw := range sizesRaw {
					if sizeMap, ok := sizeRaw.(map[string]interface{}); ok {
						sizeName := cast.ToString(sizeMap["name"])
						if sizeName != "" {
							var size product_aggregate.Size
							err := tx.Where("name = ?", sizeName).First(&size).Error
							if err == gorm.ErrRecordNotFound {
								size = product_aggregate.Size{
									Name: sizeName,
									Slug: sizeName,
								}
								if err := tx.Create(&size).Error; err != nil {
									continue
								}
							}
							product.Sizes = append(product.Sizes, size)
						}
					}
				}
			}

			// Handle tags
			if tagsRaw, ok := doc["tags"].([]interface{}); ok {
				for _, tagRaw := range tagsRaw {
					tagName := cast.ToString(tagRaw)
					if tagName != "" {
						var tag product_aggregate.Tag
						err := tx.Where("name = ?", tagName).First(&tag).Error
						if err == gorm.ErrRecordNotFound {
							tag = product_aggregate.Tag{
								Name: tagName,
								Slug: tagName,
							}
							if err := tx.Create(&tag).Error; err != nil {
								continue
							}
						}
						product.Tags = append(product.Tags, tag)
					}
				}
			}

			// Handle features
			if featuresRaw, ok := doc["features"].([]interface{}); ok {
				for _, featureRaw := range featuresRaw {
					featureStr := cast.ToString(featureRaw)
					if featureStr != "" {
						product.Features = append(product.Features, product_aggregate.ProductFeature{
							Feature: featureStr,
						})
					}
				}
			}

			// Handle specs
			if specsRaw, ok := doc["specs"].([]interface{}); ok {
				for _, specRaw := range specsRaw {
					if specMap, ok := specRaw.(map[string]interface{}); ok {
						product.Specs = append(product.Specs, product_aggregate.ProductSpec{
							Key:   cast.ToString(specMap["key"]),
							Value: cast.ToString(specMap["value"]),
						})
					}
				}
			}

			// Save product
			return tx.Create(product).Error
		})

		if err != nil {
			errorCount++
			logging.Error("Failed to save product").
				WithString("slug", cast.ToString(doc["slug"])).
				WithError(err).
				Log()
		} else {
			successCount++
			if successCount%10 == 0 {
				logging.Info("Import progress").
					WithInt("success", successCount).
					WithInt("errors", errorCount).
					Log()
			}
		}
	}

	logging.Info("Import completed").
		WithInt("success", successCount).
		WithInt("errors", errorCount).
		Log()

	return nil
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
