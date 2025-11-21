package event_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/domain/events"
	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
)

type ProductEventHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewProductEventHandler(uow unitofwork.PGUnitOfWork) *ProductEventHandler {
	return &ProductEventHandler{
		uow: uow,
	}
}

// ProductCreatedEvent handles the ProductCreatedEvent
// Saves the event to outbox table for later processing
func (h *ProductEventHandler) ProductCreatedEvent(ctx context.Context, event *events.ProductCreatedEvent) error {
	if event.ProductID == nil {
		return fmt.Errorf("product_id is nil in ProductCreatedEvent")
	}

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Convert event to JSON payload
		eventJSON, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event: %w", err)
		}

		var payload map[string]interface{}
		if err := json.Unmarshal(eventJSON, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal event to map: %w", err)
		}

		// Create outbox event
		outboxEvent := &entity.OutboxEvent{
			EventType:     "ProductCreatedEvent",
			AggregateType: "Product",
			AggregateID:   strconv.FormatUint(*event.ProductID, 10),
			Payload:       payload,
			Status:        entity.OutboxStatusPending,
			RetryCount:    0,
			MaxRetries:    5,
		}

		// Save to outbox
		if err := h.uow.Outbox(ctx).Create(ctx, outboxEvent); err != nil {
			return fmt.Errorf("failed to save event to outbox: %w", err)
		}

		logging.Info("ProductCreatedEvent saved to outbox").
			WithInt64("product_id", int64(*event.ProductID)).
			WithString("product_name", event.Name).
			WithInt64("outbox_id", int64(outboxEvent.ID)).
			Log()

		return nil
	})

	if err != nil {
		logging.Error("Failed to handle ProductCreatedEvent").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return err
	}

	return nil
}

// ProductUpdatedEvent handles the ProductUpdatedEvent
// Saves the event to outbox table for later processing
func (h *ProductEventHandler) ProductUpdatedEvent(ctx context.Context, event *events.ProductUpdatedEvent) error {
	if event.ProductID == nil {
		return fmt.Errorf("product_id is nil in ProductUpdatedEvent")
	}

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Convert event to JSON payload
		eventJSON, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event: %w", err)
		}

		var payload map[string]interface{}
		if err := json.Unmarshal(eventJSON, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal event to map: %w", err)
		}

		// Create outbox event
		outboxEvent := &entity.OutboxEvent{
			EventType:     "ProductUpdatedEvent",
			AggregateType: "Product",
			AggregateID:   strconv.FormatUint(*event.ProductID, 10),
			Payload:       payload,
			Status:        entity.OutboxStatusPending,
			RetryCount:    0,
			MaxRetries:    5,
		}

		// Save to outbox
		if err := h.uow.Outbox(ctx).Create(ctx, outboxEvent); err != nil {
			return fmt.Errorf("failed to save event to outbox: %w", err)
		}

		logging.Info("ProductUpdatedEvent saved to outbox").
			WithInt64("product_id", int64(*event.ProductID)).
			WithString("product_name", event.Name).
			WithInt64("outbox_id", int64(outboxEvent.ID)).
			Log()

		return nil
	})

	if err != nil {
		logging.Error("Failed to handle ProductUpdatedEvent").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return err
	}

	return nil
}

// ProductDeletedEvent handles the ProductDeletedEvent
// Saves the event to outbox table for later processing
func (h *ProductEventHandler) ProductDeletedEvent(ctx context.Context, event *events.ProductDeletedEvent) error {
	if event.ProductID == nil {
		return fmt.Errorf("product_id is nil in ProductDeletedEvent")
	}

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Convert event to JSON payload
		eventJSON, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event: %w", err)
		}

		var payload map[string]interface{}
		if err := json.Unmarshal(eventJSON, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal event to map: %w", err)
		}

		// Create outbox event
		outboxEvent := &entity.OutboxEvent{
			EventType:     "ProductDeletedEvent",
			AggregateType: "Product",
			AggregateID:   strconv.FormatUint(*event.ProductID, 10),
			Payload:       payload,
			Status:        entity.OutboxStatusPending,
			RetryCount:    0,
			MaxRetries:    5,
		}

		// Save to outbox
		if err := h.uow.Outbox(ctx).Create(ctx, outboxEvent); err != nil {
			return fmt.Errorf("failed to save event to outbox: %w", err)
		}

		logging.Info("ProductDeletedEvent saved to outbox").
			WithInt64("product_id", int64(*event.ProductID)).
			WithBool("soft_delete", event.SoftDelete).
			WithInt64("outbox_id", int64(outboxEvent.ID)).
			Log()

		return nil
	})

	if err != nil {
		logging.Error("Failed to handle ProductDeletedEvent").
			WithInt64("product_id", int64(*event.ProductID)).
			WithError(err).
			Log()
		return err
	}

	return nil
}
