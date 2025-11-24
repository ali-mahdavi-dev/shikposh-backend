package query

import (
	"context"
	"fmt"
	"strconv"

	"shikposh-backend/internal/seller/adapter/repository"
	sellerentity "shikposh-backend/internal/seller/domain/entity"
	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

type SellerQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewSellerQueryHandler(uow unitofwork.PGUnitOfWork) *SellerQueryHandler {
	return &SellerQueryHandler{
		uow: uow,
	}
}

func (h *SellerQueryHandler) GetSellerByID(ctx context.Context, id string) (map[string]interface{}, error) {
	sellerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid seller ID: %w", err)
	}

	var seller *sellerentity.Seller
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		// Use BaseRepository's FindByID method (expects uint64)
		seller, err = h.uow.Seller(ctx).FindByID(ctx, sellerID)
		return err
	})
	if err != nil {
		if err == repository.ErrSellerNotFound {
			return nil, repository.ErrSellerNotFound
		}
		return nil, fmt.Errorf("failed to get seller: %w", err)
	}

	// Calculate total products for this seller
	totalProducts, err := h.calculateTotalProducts(ctx, sellerentity.SellerID(sellerID))
	if err != nil {
		logging.Warn("Failed to calculate total products for seller").
			WithError(err).
			WithInt64("seller_id", int64(sellerID)).
			Log()
	}

	result := seller.ToMap()
	result["total_products"] = totalProducts

	return result, nil
}

func (h *SellerQueryHandler) GetAllSellers(ctx context.Context) ([]map[string]interface{}, error) {
	var sellers []*sellerentity.Seller
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		sellers, err = h.uow.Seller(ctx).GetAll(ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get sellers: %w", err)
	}

	result := make([]map[string]interface{}, len(sellers))
	for i, seller := range sellers {
		sellerID := uint64(seller.ID)
		totalProducts, err := h.calculateTotalProducts(ctx, seller.ID)
		if err != nil {
			logging.Warn("Failed to calculate total products for seller").
				WithError(err).
				WithInt64("seller_id", int64(sellerID)).
				Log()
		}
		summary := seller.ToSummary()
		summary["total_products"] = totalProducts
		result[i] = summary
	}

	return result, nil
}

func (h *SellerQueryHandler) GetSellersByCategory(ctx context.Context, category string) ([]map[string]interface{}, error) {
	var sellers []*sellerentity.Seller
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		sellers, err = h.uow.Seller(ctx).FindByCategory(ctx, category)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get sellers by category: %w", err)
	}

	result := make([]map[string]interface{}, len(sellers))
	for i, seller := range sellers {
		sellerID := uint64(seller.ID)
		totalProducts, err := h.calculateTotalProducts(ctx, seller.ID)
		if err != nil {
			logging.Warn("Failed to calculate total products for seller").
				WithError(err).
				WithInt64("seller_id", int64(sellerID)).
				Log()
		}
		summary := seller.ToSummary()
		summary["total_products"] = totalProducts
		result[i] = summary
	}

	return result, nil
}

func (h *SellerQueryHandler) SearchSellers(ctx context.Context, query string) ([]map[string]interface{}, error) {
	var sellers []*sellerentity.Seller
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		sellers, err = h.uow.Seller(ctx).Search(ctx, query)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search sellers: %w", err)
	}

	result := make([]map[string]interface{}, len(sellers))
	for i, seller := range sellers {
		sellerID := uint64(seller.ID)
		totalProducts, err := h.calculateTotalProducts(ctx, seller.ID)
		if err != nil {
			logging.Warn("Failed to calculate total products for seller").
				WithError(err).
				WithInt64("seller_id", int64(sellerID)).
				Log()
		}
		summary := seller.ToSummary()
		summary["total_products"] = totalProducts
		result[i] = summary
	}

	return result, nil
}

func (h *SellerQueryHandler) GetSellerByProductID(ctx context.Context, productID string) (map[string]interface{}, error) {
	// Try to parse productID as numeric ID first
	productIDNum, err := strconv.ParseUint(productID, 10, 64)
	if err == nil {
		// It's a numeric ID, try to get from database
		err = h.uow.Do(ctx, func(ctx context.Context) error {
			product, err := h.uow.Product(ctx).FindByID(ctx, productIDNum)
			if err != nil {
				return err
			}
			// Check if product has sellerId field
			// For now, products don't have seller_id column, so we return an error
			_ = product
			return fmt.Errorf("products table does not have seller_id field yet")
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get product: %w", err)
		}
	}

	// Try as slug
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		product, err := h.uow.Product(ctx).FindBySlug(ctx, productID)
		if err != nil {
			return err
		}
		// Check if product has sellerId
		// For now, products don't have seller_id column, so we return an error
		_ = product
		return fmt.Errorf("products table does not have seller_id field yet")
	})

	// If we reach here, product was not found or doesn't have seller_id
	return nil, fmt.Errorf("product not found or seller_id not available in products table: productId=%s", productID)
}

// calculateTotalProducts calculates the total number of products for a seller
// This is a placeholder - in a real scenario, products would have a seller_id field
func (h *SellerQueryHandler) calculateTotalProducts(ctx context.Context, sellerID sellerentity.SellerID) (int, error) {
	// TODO: Once products have seller_id, implement this:
	// SELECT COUNT(*) FROM products WHERE seller_id = ? AND deleted_at IS NULL
	// For now, return 0
	return 0, nil
}
