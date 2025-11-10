package event_handler

import (
	"context"

	"shikposh-backend/internal/products/domain/events"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type ProductEventHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewProductEventHandler(uow unit_of_work.PGUnitOfWork) *ProductEventHandler {
	return &ProductEventHandler{uow: uow}
}

// ProductCreatedEvent handles the ProductCreatedEvent
// Currently disabled/closed - not active yet
func (h *ProductEventHandler) ProductCreatedEvent(ctx context.Context, event *events.ProductCreatedEvent) error {
	// TODO: Implement handler logic when needed

	logging.Info("ProductCreatedEvent received (handler disabled)").
		WithInt64("product_id", int64(*event.ProductID)).
		WithString("product_name", event.Name).
		WithString("slug", event.Slug).
		Log()

	return nil
}
