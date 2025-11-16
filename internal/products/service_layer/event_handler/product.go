package event_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/domain/events"
	"github.com/shikposh/framework/infrastructure/logging"
	"shikposh-backend/internal/unit_of_work"
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
