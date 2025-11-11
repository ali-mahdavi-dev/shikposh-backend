package outbox

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	elasticsearchx "shikposh-backend/pkg/framework/infrastructure/elasticsearch"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/IBM/sarama"
)

type Consumer struct {
	uow           unit_of_work.PGUnitOfWork
	elasticsearch elasticsearchx.Connection
	kafka         interface {
		ConsumeMessages(topic string, fn func(pc sarama.PartitionConsumer)) error
	}
	indexName string
	stopChan  chan struct{}
}

func NewConsumer(
	uow unit_of_work.PGUnitOfWork,
	elasticsearch elasticsearchx.Connection,
	kafkaService interface {
		ConsumeMessages(topic string, fn func(pc sarama.PartitionConsumer)) error
	},
) *Consumer {
	return &Consumer{
		uow:           uow,
		elasticsearch: elasticsearch,
		kafka:         kafkaService,
		indexName:     "products",
		stopChan:      make(chan struct{}),
	}
}

// Start starts the Kafka consumer
func (c *Consumer) Start(ctx context.Context) error {
	if c.elasticsearch == nil {
		logging.Warn("Elasticsearch not available, consumer will not start").Log()
		return fmt.Errorf("elasticsearch connection is nil")
	}

	handler := func(pc sarama.PartitionConsumer) {
		for {
			select {
			case <-ctx.Done():
				logging.Info("Kafka consumer stopped: context cancelled").Log()
				return
			case <-c.stopChan:
				logging.Info("Kafka consumer stopped: stop signal received").Log()
				return
			case message := <-pc.Messages():
				if err := c.handleMessage(ctx, message); err != nil {
					logging.Error("Failed to handle Kafka message").
						WithError(err).
						WithInt("partition", int(message.Partition)).
						WithInt64("offset", message.Offset).
						Log()
				}
			case err := <-pc.Errors():
				if err != nil {
					logging.Error("Kafka consumer error").
						WithError(err.Err).
						Log()
				}
			}
		}
	}

	return c.kafka.ConsumeMessages(KafkaTopicProductEvents, handler)
}

// Stop stops the Kafka consumer
func (c *Consumer) Stop() {
	close(c.stopChan)
}

func (c *Consumer) handleMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	var kafkaMessage map[string]interface{}
	if err := json.Unmarshal(message.Value, &kafkaMessage); err != nil {
		return fmt.Errorf("failed to unmarshal kafka message: %w", err)
	}

	eventType, ok := kafkaMessage["event_type"].(string)
	if !ok {
		return fmt.Errorf("event_type is missing or invalid")
	}

	// Handle different event types
	switch eventType {
	case "ProductCreatedEvent":
		return c.handleProductCreatedEvent(ctx, kafkaMessage)
	default:
		logging.Warn("Unknown event type, skipping").
			WithString("event_type", eventType).
			Log()
		return nil
	}
}

func (c *Consumer) handleProductCreatedEvent(ctx context.Context, kafkaMessage map[string]interface{}) error {
	payload, ok := kafkaMessage["payload"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("payload is missing or invalid")
	}

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
	err := c.uow.Do(ctx, func(ctx context.Context) error {
		product, err := c.uow.Product(ctx).FindByID(ctx, productID)
		if err != nil {
			return fmt.Errorf("failed to get product from database: %w", err)
		}

		// Convert product to map using ToMap method
		productMap = product.ToMap()
		return nil
	})

	if err != nil {
		return err
	}

	// Index product in Elasticsearch
	productIDStr := strconv.FormatUint(productID, 10)
	if err := c.elasticsearch.IndexDocument(ctx, c.indexName, productIDStr, productMap); err != nil {
		return fmt.Errorf("failed to index product in elasticsearch: %w", err)
	}

	logging.Info("Product indexed in Elasticsearch from Kafka").
		WithInt64("product_id", int64(productID)).
		WithString("index", c.indexName).
		Log()

	return nil
}

