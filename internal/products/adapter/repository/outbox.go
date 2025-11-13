package repository

import (
	"context"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/pkg/framework/adapter"
	frameworkoutbox "shikposh-backend/pkg/framework/service_layer/outbox"

	"gorm.io/gorm"
)

type OutboxRepository interface {
	adapter.BaseRepository[*entity.OutboxEvent]
	Model(ctx context.Context) *gorm.DB
	Create(ctx context.Context, event *entity.OutboxEvent) error
	GetPendingEvents(ctx context.Context, limit int) ([]*entity.OutboxEvent, error)
	MarkAsProcessing(ctx context.Context, id entity.OutboxEventID) error
	MarkAsCompleted(ctx context.Context, id entity.OutboxEventID) error
	MarkAsFailed(ctx context.Context, id entity.OutboxEventID, errorMsg string) error
	IncrementRetry(ctx context.Context, id entity.OutboxEventID) error
}

type outboxGormRepository struct {
	adapter.BaseRepository[*entity.OutboxEvent]
	frameworkRepo frameworkoutbox.Repository
}

func NewOutboxRepository(db *gorm.DB) OutboxRepository {
	frameworkRepo := frameworkoutbox.NewGormRepository(db, "")
	return &outboxGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.OutboxEvent](db),
		frameworkRepo:  frameworkRepo,
	}
}

func (r *outboxGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.frameworkRepo.Model(ctx)
}

// convertToFramework converts products entity to framework entity
func convertToFramework(event *entity.OutboxEvent) *frameworkoutbox.OutboxEvent {
	return &frameworkoutbox.OutboxEvent{
		ID:            frameworkoutbox.OutboxEventID(event.ID),
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
		DeletedAt:     event.DeletedAt,
		EventType:     event.EventType,
		AggregateType: event.AggregateType,
		AggregateID:   event.AggregateID,
		Payload:       event.Payload,
		Status:        frameworkoutbox.OutboxEventStatus(event.Status),
		RetryCount:    event.RetryCount,
		MaxRetries:    event.MaxRetries,
		ErrorMessage:  event.ErrorMessage,
		ProcessedAt:   event.ProcessedAt,
	}
}

// convertFromFramework converts framework entity to products entity
func convertFromFramework(event *frameworkoutbox.OutboxEvent) *entity.OutboxEvent {
	return &entity.OutboxEvent{
		ID:            entity.OutboxEventID(event.ID),
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
		DeletedAt:     event.DeletedAt,
		EventType:     event.EventType,
		AggregateType: event.AggregateType,
		AggregateID:   event.AggregateID,
		Payload:       event.Payload,
		Status:        entity.OutboxEventStatus(event.Status),
		RetryCount:    event.RetryCount,
		MaxRetries:    event.MaxRetries,
		ErrorMessage:  event.ErrorMessage,
		ProcessedAt:   event.ProcessedAt,
	}
}

func (r *outboxGormRepository) Create(ctx context.Context, event *entity.OutboxEvent) error {
	frameworkEvent := convertToFramework(event)
	return r.frameworkRepo.Create(ctx, frameworkEvent)
}

func (r *outboxGormRepository) GetPendingEvents(ctx context.Context, limit int) ([]*entity.OutboxEvent, error) {
	frameworkEvents, err := r.frameworkRepo.GetPendingEvents(ctx, limit)
	if err != nil {
		return nil, err
	}

	events := make([]*entity.OutboxEvent, len(frameworkEvents))
	for i, fe := range frameworkEvents {
		events[i] = convertFromFramework(fe)
	}
	return events, nil
}

func (r *outboxGormRepository) MarkAsProcessing(ctx context.Context, id entity.OutboxEventID) error {
	return r.frameworkRepo.MarkAsProcessing(ctx, frameworkoutbox.OutboxEventID(id))
}

func (r *outboxGormRepository) MarkAsCompleted(ctx context.Context, id entity.OutboxEventID) error {
	return r.frameworkRepo.MarkAsCompleted(ctx, frameworkoutbox.OutboxEventID(id))
}

func (r *outboxGormRepository) MarkAsFailed(ctx context.Context, id entity.OutboxEventID, errorMsg string) error {
	return r.frameworkRepo.MarkAsFailed(ctx, frameworkoutbox.OutboxEventID(id), errorMsg)
}

func (r *outboxGormRepository) IncrementRetry(ctx context.Context, id entity.OutboxEventID) error {
	return r.frameworkRepo.IncrementRetry(ctx, frameworkoutbox.OutboxEventID(id))
}
