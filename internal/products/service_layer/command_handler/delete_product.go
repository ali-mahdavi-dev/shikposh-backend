package command_handler

import (
	"context"
	"errors"
	"fmt"

	"shikposh-backend/internal/products/adapter/phrases"
	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/events"

	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

func (h *ProductCommandHandler) DeleteProductHandler(ctx context.Context, cmd *commands.DeleteProduct) error {
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Find product including soft deleted ones (for delete operation)
		product, err := h.uow.Product(ctx).FindByIDIncludingDeleted(ctx, cmd.ID)
		if err != nil {
			if errors.Is(err, repository.ErrProductNotFound) {
				return apperrors.NotFound(phrases.ProductNotFound)
			}
			return fmt.Errorf("ProductCommandHandler.DeleteProductHandler error finding product: %w", err)
		}

		// Delete associated entities first
		if err := h.uow.Product(ctx).ClearAllAssociations(ctx, product); err != nil {
			return fmt.Errorf("ProductCommandHandler.DeleteProductHandler error deleting associations: %w", err)
		}

		// Emit ProductDeletedEvent BEFORE deleting (so Unit of Work can collect it)
		productID := uint64(product.ID)
		product.AddEvent(&events.ProductDeletedEvent{
			ProductID:  &productID,
			SoftDelete: cmd.SoftDelete,
		})

		// Delete product (soft or hard delete)
		// Remove will add product to Seen() so events can be collected
		if err := h.uow.Product(ctx).Remove(ctx, product, cmd.SoftDelete); err != nil {
			return fmt.Errorf("ProductCommandHandler.DeleteProductHandler error deleting product: %w", err)
		}

		return nil
	})

	return err
}
