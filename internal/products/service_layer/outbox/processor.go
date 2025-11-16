package outbox

import (
	"context"

	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/entity"
	"github.com/ali-mahdavi-dev/framework/adapter"
	frameworkoutbox "github.com/ali-mahdavi-dev/framework/service_layer/outbox"
	"shikposh-backend/internal/unit_of_work"

	"gorm.io/gorm"
)

// Processor wraps the framework outbox processor for products module
type Processor struct {
	*frameworkoutbox.Processor
}

// NewProcessor creates a new outbox processor using the framework processor
func NewProcessor(uow unitofwork.PGUnitOfWork, kafkaProducer frameworkoutbox.MessagePublisher) *Processor {
	// Get the outbox repository from UoW
	repo := uow.Outbox(context.Background())

	// Create a wrapper repository that implements frameworkoutbox.Repository
	frameworkRepo := &repositoryWrapper{repo: repo}

	// Create framework processor config
	config := frameworkoutbox.DefaultProcessorConfig("product.events")

	// Create framework processor
	frameworkProcessor := frameworkoutbox.NewProcessor(frameworkRepo, kafkaProducer, config)

	return &Processor{
		Processor: frameworkProcessor,
	}
}

// repositoryWrapper wraps products OutboxRepository to implement frameworkoutbox.Repository
type repositoryWrapper struct {
	repo repository.OutboxRepository
}

// BaseRepository methods
func (w *repositoryWrapper) FindByID(ctx context.Context, id uint64) (*frameworkoutbox.OutboxEvent, error) {
	event, err := w.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertToFrameworkEvent(event), nil
}

func (w *repositoryWrapper) FindByField(ctx context.Context, field string, value interface{}) (*frameworkoutbox.OutboxEvent, error) {
	event, err := w.repo.FindByField(ctx, field, value)
	if err != nil {
		return nil, err
	}
	return convertToFrameworkEvent(event), nil
}

func (w *repositoryWrapper) Remove(ctx context.Context, model *frameworkoutbox.OutboxEvent, softDelete bool) error {
	productsEvent := convertFromFrameworkEvent(model)
	return w.repo.Remove(ctx, productsEvent, softDelete)
}

func (w *repositoryWrapper) Modify(ctx context.Context, model *frameworkoutbox.OutboxEvent) error {
	productsEvent := convertFromFrameworkEvent(model)
	return w.repo.Modify(ctx, productsEvent)
}

func (w *repositoryWrapper) Save(ctx context.Context, model *frameworkoutbox.OutboxEvent) error {
	productsEvent := convertFromFrameworkEvent(model)
	return w.repo.Save(ctx, productsEvent)
}

func (w *repositoryWrapper) Seen() []adapter.Entity {
	seen := w.repo.Seen()
	result := make([]adapter.Entity, len(seen))
	for i, e := range seen {
		if outboxEvent, ok := e.(*entity.OutboxEvent); ok {
			result[i] = convertToFrameworkEvent(outboxEvent)
		}
	}
	return result
}

func (w *repositoryWrapper) SetSeen(model adapter.Entity) {
	if fe, ok := model.(*frameworkoutbox.OutboxEvent); ok {
		productsEvent := convertFromFrameworkEvent(fe)
		w.repo.SetSeen(productsEvent)
	}
}

// Outbox-specific methods
func (w *repositoryWrapper) Create(ctx context.Context, event *frameworkoutbox.OutboxEvent) error {
	productsEvent := convertFromFrameworkEvent(event)
	return w.repo.Create(ctx, productsEvent)
}

func (w *repositoryWrapper) GetPendingEvents(ctx context.Context, limit int) ([]*frameworkoutbox.OutboxEvent, error) {
	events, err := w.repo.GetPendingEvents(ctx, limit)
	if err != nil {
		return nil, err
	}

	frameworkEvents := make([]*frameworkoutbox.OutboxEvent, len(events))
	for i, e := range events {
		frameworkEvents[i] = convertToFrameworkEvent(e)
	}
	return frameworkEvents, nil
}

func (w *repositoryWrapper) MarkAsProcessing(ctx context.Context, id frameworkoutbox.OutboxEventID) error {
	return w.repo.MarkAsProcessing(ctx, entity.OutboxEventID(id))
}

func (w *repositoryWrapper) MarkAsCompleted(ctx context.Context, id frameworkoutbox.OutboxEventID) error {
	return w.repo.MarkAsCompleted(ctx, entity.OutboxEventID(id))
}

func (w *repositoryWrapper) MarkAsFailed(ctx context.Context, id frameworkoutbox.OutboxEventID, errorMsg string) error {
	return w.repo.MarkAsFailed(ctx, entity.OutboxEventID(id), errorMsg)
}

func (w *repositoryWrapper) IncrementRetry(ctx context.Context, id frameworkoutbox.OutboxEventID) error {
	return w.repo.IncrementRetry(ctx, entity.OutboxEventID(id))
}

func (w *repositoryWrapper) Model(ctx context.Context) *gorm.DB {
	return w.repo.Model(ctx)
}

// Helper functions for conversion
func convertToFrameworkEvent(e *entity.OutboxEvent) *frameworkoutbox.OutboxEvent {
	return &frameworkoutbox.OutboxEvent{
		ID:            frameworkoutbox.OutboxEventID(e.ID),
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		DeletedAt:     e.DeletedAt,
		EventType:     e.EventType,
		AggregateType: e.AggregateType,
		AggregateID:   e.AggregateID,
		Payload:       e.Payload,
		Status:        frameworkoutbox.OutboxEventStatus(e.Status),
		RetryCount:    e.RetryCount,
		MaxRetries:    e.MaxRetries,
		ErrorMessage:  e.ErrorMessage,
		ProcessedAt:   e.ProcessedAt,
	}
}

func convertFromFrameworkEvent(e *frameworkoutbox.OutboxEvent) *entity.OutboxEvent {
	return &entity.OutboxEvent{
		ID:            entity.OutboxEventID(e.ID),
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		DeletedAt:     e.DeletedAt,
		EventType:     e.EventType,
		AggregateType: e.AggregateType,
		AggregateID:   e.AggregateID,
		Payload:       e.Payload,
		Status:        entity.OutboxEventStatus(e.Status),
		RetryCount:    e.RetryCount,
		MaxRetries:    e.MaxRetries,
		ErrorMessage:  e.ErrorMessage,
		ProcessedAt:   e.ProcessedAt,
	}
}
