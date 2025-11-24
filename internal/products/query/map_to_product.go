package query

import (
	"context"
	"fmt"
	"strconv"

	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
)

// mapToProduct converts a map (from Elasticsearch) to Product entity
func (h *ProductQueryHandler) mapToProduct(ctx context.Context, data map[string]interface{}) (*productaggregate.Product, error) {
	if data == nil {
		return nil, fmt.Errorf("product data cannot be nil")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("product data is empty")
	}

	// Get product ID
	idValue, exists := data["id"]
	if !exists {
		return nil, fmt.Errorf("product id field is missing from elasticsearch document")
	}

	idStr, ok := idValue.(string)
	if !ok {
		// Try to convert from number if it's stored as a number
		if idNum, ok := idValue.(float64); ok {
			idStr = strconv.FormatFloat(idNum, 'f', 0, 64)
		} else {
			return nil, fmt.Errorf("product id has invalid type: expected string or number, got %T", idValue)
		}
	}

	if idStr == "" {
		return nil, fmt.Errorf("product id cannot be empty")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse product id (value=%s): %w", idStr, err)
	}

	if id == 0 {
		return nil, fmt.Errorf("product id cannot be zero")
	}

	// Get product from database to get full entity with relationships
	var product *productaggregate.Product
	err = h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		product, err = h.uow.Product(ctx).FindByID(ctx, id)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get product from database (id=%d): %w", id, err)
	}

	if product == nil {
		return nil, fmt.Errorf("product not found in database (id=%d)", id)
	}

	return product, nil
}

