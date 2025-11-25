package outbox

import (
	"context"
	"fmt"
	"strconv"

	unitofwork "shikposh-backend/internal/unit_of_work"

	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	frameworkoutbox "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/outbox"
)

// Consumer wraps the framework outbox consumer for products module
type Consumer struct {
	*frameworkoutbox.Consumer
}

// ProductEventHandler implements frameworkoutbox.EventHandler for products
type ProductEventHandler struct {
	uow           unitofwork.PGUnitOfWork
	elasticsearch elasticsearchx.Connection
	indexName     string
}

func NewConsumer(
	uow unitofwork.PGUnitOfWork,
	elasticsearch elasticsearchx.Connection,
	kafkaService frameworkoutbox.MessageConsumer,
) *Consumer {
	if elasticsearch == nil {
		logging.Warn("Elasticsearch not available, consumer will not start").Log()
		return nil
	}

	handler := &ProductEventHandler{
		uow:           uow,
		elasticsearch: elasticsearch,
		indexName:     "products",
	}

	frameworkConsumer := frameworkoutbox.NewConsumer(kafkaService, handler, "product.events")
	return &Consumer{
		Consumer: frameworkConsumer,
	}
}

// updateCategoryInProductMap updates category slug and name in productMap's categories array
func (h *ProductEventHandler) updateCategoryInProductMap(ctx context.Context, productMap map[string]interface{}, categoryID uint64) {
	if categoryID == 0 {
		return
	}

	category, err := h.uow.Category(ctx).FindByID(ctx, categoryID)
	if err != nil || category == nil {
		return
	}

	// Update categories array in productMap
	if categoriesRaw, ok := productMap["categories"].([]interface{}); ok && len(categoriesRaw) > 0 {
		if categoryMap, ok := categoriesRaw[0].(map[string]interface{}); ok {
			categoryMap["slug"] = category.Slug
			categoryMap["name"] = category.Name
		}
	}
}

// HandleEvent implements frameworkoutbox.EventHandler
func (h *ProductEventHandler) HandleEvent(ctx context.Context, eventType string, payload map[string]interface{}) error {
	switch eventType {
	case "ProductCreatedEvent":
		return h.handleProductCreatedEvent(ctx, payload)
	case "ProductUpdatedEvent":
		return h.handleProductUpdatedEvent(ctx, payload)
	case "ProductDeletedEvent":
		return h.handleProductDeletedEvent(ctx, payload)
	default:
		logging.Warn("Unknown event type, skipping").
			WithString("event_type", eventType).
			Log()
		return nil
	}
}

func (h *ProductEventHandler) handleProductCreatedEvent(ctx context.Context, payload map[string]interface{}) error {
	// Extract product_id from payload
	productIDRaw, ok := payload["product_id"]
	if !ok {
		return fmt.Errorf("product_id is missing in payload")
	}

	var productID uint64
	switch v := productIDRaw.(type) {
	case float64:
		productID = uint64(v)
	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse product_id: %w", err)
		}
		productID = parsed
	default:
		return fmt.Errorf("product_id has invalid type")
	}

	// Get full product from database
	var productMap map[string]interface{}
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		product, err := h.uow.Product(ctx).FindByID(ctx, productID)
		if err != nil {
			return fmt.Errorf("failed to get product from database: %w", err)
		}

		// Convert product to map using ToMap method
		productMap = product.ToMap()

		// Update category slug and name in categories array
		h.updateCategoryInProductMap(ctx, productMap, product.CategoryID)

		return nil
	})

	if err != nil {
		return err
	}

	// Index product in Elasticsearch
	productIDStr := strconv.FormatUint(productID, 10)
	if err := h.elasticsearch.IndexDocument(ctx, h.indexName, productIDStr, productMap); err != nil {
		return fmt.Errorf("failed to index product in elasticsearch: %w", err)
	}

	logging.Info("Product indexed in Elasticsearch from Kafka").
		WithInt64("product_id", int64(productID)).
		WithString("index", h.indexName).
		Log()

	return nil
}

func (h *ProductEventHandler) handleProductUpdatedEvent(ctx context.Context, payload map[string]interface{}) error {
	// Extract product_id from payload
	productIDRaw, ok := payload["product_id"]
	if !ok {
		return fmt.Errorf("product_id is missing in payload")
	}

	var productID uint64
	switch v := productIDRaw.(type) {
	case float64:
		productID = uint64(v)
	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse product_id: %w", err)
		}
		productID = parsed
	default:
		return fmt.Errorf("product_id has invalid type")
	}

	// Get full product from database
	var productMap map[string]interface{}
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		product, err := h.uow.Product(ctx).FindByID(ctx, productID)
		if err != nil {
			return fmt.Errorf("failed to get product from database: %w", err)
		}

		// Convert product to map using ToMap method
		productMap = product.ToMap()

		// Update category slug and name in categories array
		h.updateCategoryInProductMap(ctx, productMap, product.CategoryID)

		return nil
	})

	if err != nil {
		return err
	}

	// Update product in Elasticsearch
	productIDStr := strconv.FormatUint(productID, 10)
	if err := h.elasticsearch.IndexDocument(ctx, h.indexName, productIDStr, productMap); err != nil {
		return fmt.Errorf("failed to update product in elasticsearch: %w", err)
	}

	logging.Info("Product updated in Elasticsearch from Kafka").
		WithInt64("product_id", int64(productID)).
		WithString("index", h.indexName).
		Log()

	return nil
}

func (h *ProductEventHandler) handleProductDeletedEvent(ctx context.Context, payload map[string]interface{}) error {
	// Extract product_id from payload
	productIDRaw, ok := payload["product_id"]
	if !ok {
		return fmt.Errorf("product_id is missing in payload")
	}

	var productID uint64
	switch v := productIDRaw.(type) {
	case float64:
		productID = uint64(v)
	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse product_id: %w", err)
		}
		productID = parsed
	default:
		return fmt.Errorf("product_id has invalid type")
	}

	// Delete product from Elasticsearch
	productIDStr := strconv.FormatUint(productID, 10)
	if err := h.elasticsearch.DeleteDocument(ctx, h.indexName, productIDStr); err != nil {
		return fmt.Errorf("failed to delete product from elasticsearch: %w", err)
	}

	logging.Info("Product deleted from Elasticsearch from Kafka").
		WithInt64("product_id", int64(productID)).
		WithString("index", h.indexName).
		Log()

	return nil
}
