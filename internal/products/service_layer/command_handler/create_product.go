package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/domain/entity/shared"
	appadapter "shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/errors/phrases"
)

func (h *ProductCommandHandler) CreateProductHandler(ctx context.Context, cmd *commands.CreateProduct) error {
	return h.uow.Do(ctx, func(ctx context.Context) error {
		// Verify category exists
		_, err := h.uow.Category(ctx).FindByID(ctx, cmd.CategoryID)
		if err != nil {
			if errors.Is(err, appadapter.ErrEntityNotFound) {
				return apperrors.NotFound(phrases.UserNotFound, "Category not found")
			}

			return fmt.Errorf("ProductCommandHandler.CreateProductHandler error finding category: %w", err)
		}

		// Check if slug already exists
		_, err = h.uow.Product(ctx).FindBySlug(ctx, cmd.Slug)
		if err == nil {
			return apperrors.Conflict("", fmt.Sprintf("Product with slug '%s' already exists", cmd.Slug))
		} else if !errors.Is(err, repository.ErrProductNotFound) {
			return fmt.Errorf("ProductCommandHandler.CreateProductHandler error checking slug: %w", err)
		}

		// Create product
		cmd.Slug = GenerateSlug(cmd.Name)
		product := entity.NewProduct(cmd)

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

		// Save product (GORM will handle associations)
		if err := h.uow.Product(ctx).Save(ctx, product); err != nil {
			return fmt.Errorf("ProductCommandHandler.CreateProductHandler error saving product: %w", err)
		}

		return nil
	})
}
