package commands

import (
	"context"
	"fmt"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func updateImagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-images",
		Short: "update product image URLs and reindex",
		RunE: func(_ *cobra.Command, _ []string) error {
			initializeConfigs()
			return updateImagesAndReindex()
		},
	}

	return cmd
}

func updateImagesAndReindex() error {
	// Initialize database connection
	db, err := initializeDatabase(&cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer closeDatabase(db)

	// Update image URLs in database
	logging.Info("Updating product image URLs in database").Log()

	// Map of old URLs to new URLs
	imageUpdates := map[string]string{
		"https://example.com/images/women-suit.jpg": "/images/Women-Formal.avif",
		"https://example.com/images/handbag.jpg":    "/images/handbag.jpg",
		"https://example.com/images/jewelry.jpg":    "/images/jewelry.jpg",
		"https://example.com/images/suit-Top.jpg":   "/images/suit-Top.jpg",
		"https://example.com/images/shoes.jpg":      "/images/shoes.jpg",
		"https://example.com/images/harir.jpeg":     "/images/harir.jpeg",
	}

	ctx := context.Background()
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for oldURL, newURL := range imageUpdates {
			result := tx.Model(&struct {
				ID    uint64 `gorm:"primaryKey"`
				Image string `gorm:"image"`
			}{}).
				Table("products").
				Where("image = ?", oldURL).
				Update("image", newURL)

			if result.Error != nil {
				return fmt.Errorf("failed to update image URL %s: %w", oldURL, result.Error)
			}

			if result.RowsAffected > 0 {
				logging.Info("Updated image URL").
					WithString("old_url", oldURL).
					WithString("new_url", newURL).
					WithInt64("rows_affected", result.RowsAffected).
					Log()
			}
		}

		// Also update category images
		categoryImageUpdates := map[string]string{
			"https://example.com/images/women-clothing.jpg": "/images/Women-Formal.avif",
			"https://example.com/images/bags-shoes.jpg":     "/images/handbag.jpg",
			"https://example.com/images/jewelry.jpg":        "/images/jewelry.jpg",
			"https://example.com/images/men-clothing.jpg":   "/images/suit-Top.jpg",
			"https://example.com/images/accessories.jpg":    "/images/harir.jpeg",
		}

		for oldURL, newURL := range categoryImageUpdates {
			result := tx.Model(&struct {
				ID    uint64 `gorm:"primaryKey"`
				Image string `gorm:"image"`
			}{}).
				Table("categories").
				Where("image = ?", oldURL).
				Update("image", newURL)

			if result.Error != nil {
				return fmt.Errorf("failed to update category image URL %s: %w", oldURL, result.Error)
			}

			if result.RowsAffected > 0 {
				logging.Info("Updated category image URL").
					WithString("old_url", oldURL).
					WithString("new_url", newURL).
					WithInt64("rows_affected", result.RowsAffected).
					Log()
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to update images: %w", err)
	}

	logging.Info("Image URLs updated successfully").Log()

	// Now reindex Elasticsearch
	logging.Info("Starting reindex process").Log()
	return reindexProducts()
}
