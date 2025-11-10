package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/errors/phrases"

	"gorm.io/gorm"
)

func (h *ProductCommandHandler) DeleteProductHandler(ctx context.Context, cmd *commands.DeleteProduct) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Find product
		product, err := h.uow.Product(ctx).FindByID(ctx, cmd.ID)
		if err != nil {
			if errors.Is(err, repository.ErrProductNotFound) || errors.Is(err, gorm.ErrRecordNotFound) {
				return apperrors.NotFound(phrases.UserNotFound, "Product not found")
			}
			return fmt.Errorf("ProductCommandHandler.DeleteProductHandler error finding product: %w", err)
		}

		// Delete associated entities first
		if err := h.uow.Product(ctx).ClearAllAssociations(ctx, product); err != nil {
			return fmt.Errorf("ProductCommandHandler.DeleteProductHandler error deleting associations: %w", err)
		}

		// Delete product (soft or hard delete)
		if err := h.uow.Product(ctx).Remove(ctx, product, cmd.SoftDelete); err != nil {
			return fmt.Errorf("ProductCommandHandler.DeleteProductHandler error deleting product: %w", err)
		}

		return nil
	})

	return err
}
