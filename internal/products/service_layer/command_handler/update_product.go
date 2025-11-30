package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/phrases"
	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/domain/specification"

	appadapter "github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *ProductCommandHandler) UpdateProductHandler(ctx context.Context, cmd *commands.UpdateProduct) error {
	cmd.Slug = GenerateSlug(cmd.Slug)

	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Find existing product
		product, err := h.uow.Product(ctx).FindByID(ctx, cmd.ID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound(phrases.ProductNotFound)
			}
			return fmt.Errorf("UpdateProductHandler: error finding product: %w", err)
		}

		// Verify categories exist if provided
		if cmd.Categories != nil {
			for _, catID := range cmd.Categories {
				_, err = h.uow.Category(ctx).FindByID(ctx, catID)
				if err != nil {
					if errors.Is(err, appadapter.ErrEntityNotFound) {
						return apperrors.NotFound(phrases.CategoryNotFound)
					}
					return fmt.Errorf("UpdateProductHandler: error finding category: %w", err)
				}
			}
		}

		// Check slug uniqueness
		if cmd.Slug != product.Slug {
			existingProduct, err := h.uow.Product(ctx).FindBySlug(ctx, cmd.Slug)
			if err == nil && existingProduct != nil && existingProduct.ID != product.ID {
				return apperrors.Conflict(phrases.ProductSlugExists)
			}
			if err != nil && !errors.Is(err, repository.ErrProductNotFound) {
				return fmt.Errorf("UpdateProductHandler: error checking slug: %w", err)
			}
		}

		// Update required fields
		product.Title = cmd.Title
		product.Slug = cmd.Slug
		product.Brand = cmd.Brand
		product.Description = cmd.Description
		if cmd.ShortDescription != nil {
			product.ShortDescription = cmd.ShortDescription
		}

		if cmd.SellerID != nil {
			product.SellerID = cmd.SellerID
		}
		if cmd.Thumbnail != nil {
			product.Thumbnail = *cmd.Thumbnail
		}
		if cmd.Discount != nil {
			product.Discount = *cmd.Discount
		}
		if cmd.Stock != nil {
			product.Stock = *cmd.Stock
		}
		if cmd.Price != nil {
			product.Price = *cmd.Price
		}
		if cmd.OriginPrice != nil {
			product.OriginPrice = cmd.OriginPrice
		}
		if cmd.IsNew != nil {
			product.IsNew = *cmd.IsNew
		}
		if cmd.IsFeatured != nil {
			product.IsFeatured = *cmd.IsFeatured
		}

		// Update Categories (M:N)
		// Don't clear here - let Modify handle clearing and appending
		if cmd.Categories != nil {
			if len(cmd.Categories) > 0 {
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
			} else {
				product.Categories = []entity.Category{}
			}
		}

		// Update Colors
		// Don't clear here - let Modify handle clearing and appending
		if cmd.Colors != nil {
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
			} else {
				product.Colors = []product_aggregate.Color{}
			}
		}

		// Update Sizes
		// Don't clear here - let Modify handle clearing and appending
		if cmd.Sizes != nil {
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
			} else {
				product.Sizes = []product_aggregate.Size{}
			}
		}

		// Update Tags
		// Don't clear here - let Modify handle clearing and appending
		if cmd.Tags != nil {
			if len(cmd.Tags) > 0 {
				tags := make([]product_aggregate.Tag, 0, len(cmd.Tags))
				for _, tagName := range cmd.Tags {
					tag, err := h.uow.Tag(ctx).FindOrCreateByName(ctx, tagName)
					if err != nil {
						return fmt.Errorf("UpdateProductHandler: error finding/creating tag: %w", err)
					}
					tags = append(tags, product_aggregate.Tag{
						ID:   tag.ID,
						Name: tag.Name,
						Slug: tag.Slug,
					})
				}
				product.Tags = tags
			} else {
				product.Tags = []product_aggregate.Tag{}
			}
		}

		// Update Features
		// Don't clear here - let Modify handle clearing and inserting
		if cmd.Features != nil {
			if len(cmd.Features) > 0 {
				features := make([]product_aggregate.ProductFeature, len(cmd.Features))
				for i, f := range cmd.Features {
					features[i] = product_aggregate.ProductFeature{
						ProductID: product.ID,
						Feature:   f,
						Order:     i,
					}
				}
				product.Features = features
			} else {
				product.Features = []product_aggregate.ProductFeature{}
			}
		}

		// Update Specs
		// Don't clear here - let Modify handle clearing and inserting
		if cmd.Specs != nil {
			if len(cmd.Specs) > 0 {
				specs := make([]product_aggregate.ProductSpec, len(cmd.Specs))
				for i, s := range cmd.Specs {
					specs[i] = product_aggregate.ProductSpec{
						ProductID: product.ID,
						Key:       s.Key,
						Value:     s.Value,
						Order:     s.Order,
					}
				}
				product.Specs = specs
			} else {
				product.Specs = []product_aggregate.ProductSpec{}
			}
		}

		// Update Variants
		// Don't clear here - let Modify handle clearing and inserting
		if cmd.Variants != nil {
			if len(cmd.Variants) > 0 {
				variants := make([]product_aggregate.ProductVariant, len(cmd.Variants))
				for i, v := range cmd.Variants {
					variants[i] = product_aggregate.ProductVariant{
						ProductID: product.ID,
						ColorID:   product_aggregate.ColorID(v.ColorID),
						SizeID:    product_aggregate.SizeID(v.SizeID),
						Stock:     v.Stock,
					}
				}
				product.Variants = variants
			} else {
				product.Variants = []product_aggregate.ProductVariant{}
			}
		}

		// Update Images
		// Don't clear here - let Modify handle clearing and inserting
		if cmd.Images != nil {
			if len(cmd.Images) > 0 {
				var images []product_aggregate.ProductImage
				for _, img := range cmd.Images {
					for order, url := range img.URLs {
						images = append(images, product_aggregate.ProductImage{
							ProductID: product.ID,
							ColorID:   product_aggregate.ColorID(img.ColorID),
							URL:       url,
							SortOrder: order,
						})
					}
				}
				product.Images = images
			} else {
				product.Images = []product_aggregate.ProductImage{}
			}
		}

		// Validate
		canBePublishedSpec := specification.NewProductCanBePublishedSpecification()
		if !canBePublishedSpec.IsSatisfiedBy(product) {
			return apperrors.Validation("", "Product must have title, slug, category, and price")
		}

		// Emit event BEFORE saving (so Unit of Work can collect it)
		product.EmitUpdatedEvent()

		// Save (this will add product to Seen() and events will be collected)
		if err := h.uow.Product(ctx).Modify(ctx, product); err != nil {
			return fmt.Errorf("UpdateProductHandler: error saving product: %w", err)
		}

		return nil
	})
}
