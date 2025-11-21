package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/phrases"
	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/domain/entity/shared"
	"shikposh-backend/internal/products/domain/specification"

	appadapter "github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *ProductCommandHandler) CreateProductHandler(ctx context.Context, cmd *commands.CreateProduct) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Verify category exists
		_, err := h.uow.Category(ctx).FindByID(ctx, cmd.CategoryID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound(phrases.CategoryNotFound)
			}

			return fmt.Errorf("ProductCommandHandler.CreateProductHandler error finding category: %w", err)
		}

		// Check if slug already exists
		_, err = h.uow.Product(ctx).FindBySlug(ctx, cmd.Slug)
		if err == nil {
			return apperrors.Conflict(phrases.ProductSlugExists)
		} else if !errors.Is(err, repository.ErrProductNotFound) {
			return fmt.Errorf("ProductCommandHandler.CreateProductHandler error checking slug: %w", err)
		}

		// Create product
		cmd.Slug = GenerateSlug(cmd.Name)
		product := product_aggregate.NewProduct(cmd)

		// Convert Features
		if len(cmd.Features) > 0 {
			product.Features = make([]product_aggregate.ProductFeature, len(cmd.Features))
			for i, f := range cmd.Features {
				product.Features[i] = product_aggregate.NewProductFeature(0, f.Feature, f.Order)
			}
		}

		// Convert Details
		if len(cmd.Details) > 0 {
			product.Details = make([]product_aggregate.ProductDetail, len(cmd.Details))
			for i, d := range cmd.Details {
				product.Details[i] = product_aggregate.NewProductDetail(0, d)

				// Convert image paths to attachments
				if len(d.Images) > 0 {
					product.Details[i].Images = make([]shared.Attachment, len(d.Images))
					for j, imgPath := range d.Images {
						product.Details[i].Images[j] = shared.NewAttachment(imgPath, "image")
					}
				}
			}
		}

		// Convert Specs
		if len(cmd.Specs) > 0 {
			product.Specs = make([]product_aggregate.ProductSpec, len(cmd.Specs))
			for i, s := range cmd.Specs {
				product.Specs[i] = product_aggregate.NewProductSpec(0, s)
			}
		}

		// Convert Tags - find or create tags by name
		if len(cmd.Tags) > 0 {
			tags := make([]shared.Tag, 0, len(cmd.Tags))
			for _, tagName := range cmd.Tags {
				tag, err := h.uow.Tag(ctx).FindOrCreateByName(ctx, tagName)
				if err != nil {
					return fmt.Errorf("ProductCommandHandler.CreateProductHandler error finding/creating tag: %w", err)
				}
				tags = append(tags, *tag)
			}
			product.Tags = tags
		}

		// Convert Sizes - find or create sizes by name
		if len(cmd.Sizes) > 0 {
			sizes := make([]shared.Size, 0, len(cmd.Sizes))
			for _, sizeName := range cmd.Sizes {
				size, err := h.uow.Size(ctx).FindOrCreateByName(ctx, sizeName)
				if err != nil {
					return fmt.Errorf("ProductCommandHandler.CreateProductHandler error finding/creating size: %w", err)
				}
				sizes = append(sizes, *size)
			}
			product.Sizes = sizes
		}

		// Validate product using specification pattern
		canBePublishedSpec := specification.NewProductCanBePublishedSpecification()
		if !canBePublishedSpec.IsSatisfiedBy(product) {
			return apperrors.Validation("", "Product must have a name, slug, category, and at least one detail with price to be created")
		}

		// Save product (GORM will handle associations)
		if err := h.uow.Product(ctx).Save(ctx, product); err != nil {
			return fmt.Errorf("ProductCommandHandler.CreateProductHandler error saving product: %w", err)
		}

		return nil
	})
}
