package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/product/adapter/phrases"
	"shikposh-backend/internal/product/adapter/repository"
	"shikposh-backend/internal/product/domain/commands"
	"shikposh-backend/internal/product/domain/entity"
	"shikposh-backend/internal/product/domain/entity/product_aggregate"
	"shikposh-backend/internal/product/domain/specification"

	appadapter "github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *ProductCommandHandler) CreateProductHandler(ctx context.Context, cmd *commands.CreateProduct) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Verify at least one category exists
		if len(cmd.Categories) == 0 {
			return apperrors.Validation("", "At least one category is required")
		}

		for _, catID := range cmd.Categories {
			_, err := h.uow.Category(ctx).FindByID(ctx, catID)
			if err != nil {
				if errors.Is(err, appadapter.ErrEntityNotFound) {
					return apperrors.NotFound(phrases.CategoryNotFound)
				}
				return fmt.Errorf("CreateProductHandler: error finding category: %w", err)
			}
		}

		// Check if slug already exists
		cmd.Slug = GenerateSlug(cmd.Title)
		_, err := h.uow.Product(ctx).FindBySlug(ctx, cmd.Slug)
		if err == nil {
			return apperrors.Conflict(phrases.ProductSlugExists)
		} else if !errors.Is(err, repository.ErrProductNotFound) {
			return fmt.Errorf("CreateProductHandler: error checking slug: %w", err)
		}

		// Create product using aggregate constructor
		product := product_aggregate.NewProduct(cmd)

		// Categories (M:N)
		categories := make([]entity.Category, 0, len(cmd.Categories))
		for _, catID := range cmd.Categories {
			cat, err := h.uow.Category(ctx).FindByID(ctx, catID)
			if err != nil {
				continue
			}
			categories = append(categories, entity.Category{
				ID:   cat.ID,
				Name: cat.Name,
				Slug: cat.Slug,
			})
		}
		product.Categories = categories

		// Colors (M:N)
		if len(cmd.Colors) > 0 {
			colors := make([]product_aggregate.Color, 0, len(cmd.Colors))
			for _, colorID := range cmd.Colors {
				color, err := h.uow.Color(ctx).FindByID(ctx, colorID)
				if err != nil {
					continue
				}
				colors = append(colors, product_aggregate.Color{
					ID:   color.ID,
					Name: color.Name,
					Slug: color.Slug,
					Hex:  color.Hex,
				})
			}
			product.Colors = colors
		}

		// Sizes (M:N)
		if len(cmd.Sizes) > 0 {
			sizes := make([]product_aggregate.Size, 0, len(cmd.Sizes))
			for _, sizeID := range cmd.Sizes {
				size, err := h.uow.Size(ctx).FindByID(ctx, sizeID)
				if err != nil {
					continue
				}
				sizes = append(sizes, product_aggregate.Size{
					ID:   size.ID,
					Name: size.Name,
					Slug: size.Slug,
				})
			}
			product.Sizes = sizes
		}

		// Tags
		if len(cmd.Tags) > 0 {
			tags := make([]product_aggregate.Tag, 0, len(cmd.Tags))
			for _, tagName := range cmd.Tags {
				tag, err := h.uow.Tag(ctx).FindOrCreateByName(ctx, tagName)
				if err != nil {
					return fmt.Errorf("CreateProductHandler: error finding/creating tag: %w", err)
				}
				tags = append(tags, product_aggregate.Tag{
					ID:   tag.ID,
					Name: tag.Name,
					Slug: tag.Slug,
				})
			}
			product.Tags = tags
		}

		// Features
		if len(cmd.Features) > 0 {
			features := make([]product_aggregate.ProductFeature, len(cmd.Features))
			for i, f := range cmd.Features {
				features[i] = product_aggregate.ProductFeature{
					Feature: f,
					Order:   i,
				}
			}
			product.Features = features
		}

		// Specs
		if len(cmd.Specs) > 0 {
			specs := make([]product_aggregate.ProductSpec, len(cmd.Specs))
			for i, s := range cmd.Specs {
				specs[i] = product_aggregate.ProductSpec{
					Key:   s.Key,
					Value: s.Value,
					Order: s.Order,
				}
			}
			product.Specs = specs
		}

		// Variants (color_id + size_id + stock)
		if len(cmd.Variants) > 0 {
			variants := make([]product_aggregate.ProductVariant, len(cmd.Variants))
			for i, v := range cmd.Variants {
				variants[i] = product_aggregate.ProductVariant{
					ColorID: product_aggregate.ColorID(v.ColorID),
					SizeID:  product_aggregate.SizeID(v.SizeID),
					Stock:   v.Stock,
				}
			}
			product.Variants = variants
		}

		// Images (color_id + urls)
		if len(cmd.Images) > 0 {
			var images []product_aggregate.ProductImage
			for _, img := range cmd.Images {
				for order, url := range img.URLs {
					images = append(images, product_aggregate.ProductImage{
						ColorID:   product_aggregate.ColorID(img.ColorID),
						URL:       url,
						SortOrder: order,
					})
				}
			}
			product.Images = images
		}

		// Validate
		canBePublishedSpec := specification.NewProductCanBePublishedSpecification()
		if !canBePublishedSpec.IsSatisfiedBy(product) {
			return apperrors.Validation("", "Product must have title, slug, category, and price")
		}

		// Save
		if err := h.uow.Product(ctx).Save(ctx, product); err != nil {
			return fmt.Errorf("CreateProductHandler: error saving product: %w", err)
		}

		// Emit event after successful save (product ID is now available)
		product.EmitCreatedEvent()

		return nil
	})
}
