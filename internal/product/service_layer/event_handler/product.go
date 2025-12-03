package event_handler

import (
	"context"
	"fmt"
	"strconv"

	"shikposh-backend/internal/product/domain/events"
	unitofwork "shikposh-backend/internal/unit_of_work"

	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

type ProductEventHandler struct {
	uow           unitofwork.PGUnitOfWork
	elasticsearch elasticsearchx.Connection
	indexName     string
}

func NewProductEventHandler(uow unitofwork.PGUnitOfWork, elasticsearch elasticsearchx.Connection) *ProductEventHandler {
	return &ProductEventHandler{
		uow:           uow,
		elasticsearch: elasticsearch,
		indexName:     "products",
	}
}

// ProductCreatedEvent handles the ProductCreatedEvent
// Indexes the product directly in Elasticsearch
func (h *ProductEventHandler) ProductCreatedEvent(ctx context.Context, event *events.ProductCreatedEvent) error {
	if event.ProductID == nil {
		return fmt.Errorf("product_id is nil in ProductCreatedEvent")
	}

	if h.elasticsearch == nil {
		logging.Warn("Elasticsearch not available, skipping ProductCreatedEvent indexing").
			WithInt64("product_id", int64(*event.ProductID)).
			Log()
		return nil
	}

	// Get full product from database
	var productMap map[string]interface{}
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		product, err := h.uow.Product(ctx).FindByID(ctx, *event.ProductID)
		if err != nil {
			return fmt.Errorf("failed to get product from database: %w", err)
		}

		// Convert product to map using ToMap method
		productMap = product.ToMap()
		return nil
	})

	if err != nil {
		logging.Error("Failed to get product for indexing").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return err
	}

	// Index product in Elasticsearch
	productIDStr := strconv.FormatUint(*event.ProductID, 10)
	if err := h.elasticsearch.IndexDocument(ctx, h.indexName, productIDStr, productMap); err != nil {
		logging.Error("Failed to index product in Elasticsearch").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return fmt.Errorf("failed to index product in elasticsearch: %w", err)
	}

	logging.Info("Product indexed in Elasticsearch").
		WithInt64("product_id", int64(*event.ProductID)).
		WithString("product_name", event.Name).
		WithString("index", h.indexName).
		Log()

	return nil
}

// ProductUpdatedEvent handles the ProductUpdatedEvent
// Updates the product in Elasticsearch
func (h *ProductEventHandler) ProductUpdatedEvent(ctx context.Context, event *events.ProductUpdatedEvent) error {
	if event.ProductID == nil {
		return fmt.Errorf("product_id is nil in ProductUpdatedEvent")
	}

	if h.elasticsearch == nil {
		logging.Warn("Elasticsearch not available, skipping ProductUpdatedEvent indexing").
			WithInt64("product_id", int64(*event.ProductID)).
			Log()
		return nil
	}

	// Get full product from database
	var productMap map[string]interface{}
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		product, err := h.uow.Product(ctx).FindByID(ctx, *event.ProductID)
		if err != nil {
			return fmt.Errorf("failed to get product from database: %w", err)
		}

		// Convert product to map using ToMap method
		productMap = product.ToMap()
		return nil
	})

	if err != nil {
		logging.Error("Failed to get product for indexing").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return err
	}

	// Update product in Elasticsearch (IndexDocument updates if exists)
	productIDStr := strconv.FormatUint(*event.ProductID, 10)
	if err := h.elasticsearch.IndexDocument(ctx, h.indexName, productIDStr, productMap); err != nil {
		logging.Error("Failed to update product in Elasticsearch").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return fmt.Errorf("failed to update product in elasticsearch: %w", err)
	}

	logging.Info("Product updated in Elasticsearch").
		WithInt64("product_id", int64(*event.ProductID)).
		WithString("product_name", event.Name).
		WithString("index", h.indexName).
		Log()

	return nil
}

// ProductDeletedEvent handles the ProductDeletedEvent
// Deletes the product from Elasticsearch
func (h *ProductEventHandler) ProductDeletedEvent(ctx context.Context, event *events.ProductDeletedEvent) error {
	if event.ProductID == nil {
		return fmt.Errorf("product_id is nil in ProductDeletedEvent")
	}

	if h.elasticsearch == nil {
		logging.Warn("Elasticsearch not available, skipping ProductDeletedEvent deletion").
			WithInt64("product_id", int64(*event.ProductID)).
			Log()
		return nil
	}

	// Delete product from Elasticsearch
	productIDStr := strconv.FormatUint(*event.ProductID, 10)
	if err := h.elasticsearch.DeleteDocument(ctx, h.indexName, productIDStr); err != nil {
		logging.Error("Failed to delete product from Elasticsearch").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return fmt.Errorf("failed to delete product from elasticsearch: %w", err)
	}

	logging.Info("Product deleted from Elasticsearch").
		WithInt64("product_id", int64(*event.ProductID)).
		WithBool("soft_delete", event.SoftDelete).
		WithString("index", h.indexName).
		Log()

	return nil
}
