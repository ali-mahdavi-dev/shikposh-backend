package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/domain/entity/shared"
	"shikposh-backend/internal/products/domain/specification"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/errors/phrases"
)

func (h *ProductCommandHandler) UpdateProductHandler(ctx context.Context, cmd *commands.UpdateProduct) error {
	cmd.Slug = GenerateSlug(cmd.Slug)

	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Find existing product
		product, err := h.uow.Product(ctx).FindByID(ctx, cmd.ID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound(phrases.UserNotFound, "Product not found")
			}

			return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error finding product: %w", err)
		}

		// Verify category exists
		_, err = h.uow.Category(ctx).FindByID(ctx, cmd.CategoryID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound(phrases.UserNotFound, "Category not found")
			}
			return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error finding category: %w", err)
		}

		// Check if new slug already exists (and is not the current product)
		if cmd.Slug != product.Slug {
			existingProduct, err := h.uow.Product(ctx).FindBySlug(ctx, cmd.Slug)
			if err == nil && existingProduct != nil && existingProduct.ID != product.ID {
				return apperrors.Conflict("", fmt.Sprintf("Product with slug '%s' already exists", cmd.Slug))
			}
			if err != nil && !errors.Is(err, repository.ErrProductNotFound) {
				return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error checking slug: %w", err)
			}
		}

		// Update required fields
		product.Name = cmd.Name
		product.Slug = cmd.Slug
		product.Brand = cmd.Brand
		product.CategoryID = cmd.CategoryID

		product.Description = cmd.Description
		if cmd.Tags != nil {
			product.Tags = cmd.Tags
		}
		if cmd.Sizes != nil {
			product.Sizes = cmd.Sizes
		}
		if cmd.Image != nil {
			product.Image = *cmd.Image
		}
		if cmd.IsNew != nil {
			product.IsNew = *cmd.IsNew
		}
		if cmd.IsFeatured != nil {
			product.IsFeatured = *cmd.IsFeatured
		}

		// Update Features if provided
		if cmd.Features != nil {
			// Delete existing features
			if err := h.uow.Product(ctx).ClearFeatures(ctx, product); err != nil {
				return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error deleting features: %w", err)
			}

			// Create new features
			if len(cmd.Features) > 0 {
				product.Features = make([]product_aggregate.ProductFeature, len(cmd.Features))
				for i, f := range cmd.Features {
					product.Features[i] = product_aggregate.NewProductFeature(product.ID, f.Feature, f.Order)
				}
			} else {
				product.Features = []product_aggregate.ProductFeature{}
			}
		}

		// Update Details if provided
		if cmd.Details != nil {
			// Delete existing details and their attachments
			if err := h.uow.Product(ctx).ClearDetails(ctx, product); err != nil {
				return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error deleting details: %w", err)
			}

			// Create new details
			if len(cmd.Details) > 0 {
				product.Details = make([]product_aggregate.ProductDetail, len(cmd.Details))
				for i, d := range cmd.Details {
					product.Details[i] = product_aggregate.NewProductDetail(product.ID, d)

					// Convert image paths to attachments
					if len(d.Images) > 0 {
						product.Details[i].Images = make([]shared.Attachment, len(d.Images))
						for j, imgPath := range d.Images {
							product.Details[i].Images[j] = shared.NewAttachment(imgPath, "image")
						}
					}
				}
			} else {
				product.Details = []product_aggregate.ProductDetail{}
			}
		}

		// Update Specs if provided
		if cmd.Specs != nil {
			// Delete existing specs
			if err := h.uow.Product(ctx).ClearSpecs(ctx, product); err != nil {
				return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error deleting specs: %w", err)
			}

			// Create new specs
			if len(cmd.Specs) > 0 {
				product.Specs = make([]product_aggregate.ProductSpec, len(cmd.Specs))
				for i, s := range cmd.Specs {
					product.Specs[i] = product_aggregate.NewProductSpec(product.ID, s)
				}
			} else {
				product.Specs = []product_aggregate.ProductSpec{}
			}
		}

		// Validate product using specification pattern
		canBePublishedSpec := specification.NewProductCanBePublishedSpecification()
		if !canBePublishedSpec.IsSatisfiedBy(product) {
			return apperrors.Validation("", "Product must have a name, slug, category, and at least one detail with price to be updated")
		}

		// Save product
		if err := h.uow.Product(ctx).Modify(ctx, product); err != nil {
			return fmt.Errorf("ProductCommandHandler.UpdateProductHandler error saving product: %w", err)
		}

		return nil
	})
}
