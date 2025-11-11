package repository

import (
	"context"
	"time"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

type OutboxRepository interface {
	adapter.BaseRepository[*entity.OutboxEvent]
	Create(ctx context.Context, event *entity.OutboxEvent) error
	GetPendingEvents(ctx context.Context, limit int) ([]*entity.OutboxEvent, error)
	MarkAsProcessing(ctx context.Context, id uint64) error
	MarkAsCompleted(ctx context.Context, id uint64) error
	MarkAsFailed(ctx context.Context, id uint64, errorMsg string) error
	IncrementRetry(ctx context.Context, id uint64) error
}

type outboxGormRepository struct {
	adapter.BaseRepository[*entity.OutboxEvent]
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) OutboxRepository {
	return &outboxGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.OutboxEvent](db),
		db:             db,
	}
}

func (r *outboxGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.OutboxEvent{})
}

func (r *outboxGormRepository) Create(ctx context.Context, event *entity.OutboxEvent) error {
	return r.Model(ctx).Create(event).Error
}

func (r *outboxGormRepository) GetPendingEvents(ctx context.Context, limit int) ([]*entity.OutboxEvent, error) {
	var events []*entity.OutboxEvent
	err := r.Model(ctx).
		Where("status = ?", entity.OutboxStatusPending).
		Where("retry_count < max_retries").
		Order("created_at ASC").
		Limit(limit).
		Find(&events).Error
	return events, err
}

func (r *outboxGormRepository) MarkAsProcessing(ctx context.Context, id uint64) error {
	return r.Model(ctx).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     entity.OutboxStatusProcessing,
			"updated_at": time.Now(),
		}).Error
}

func (r *outboxGormRepository) MarkAsCompleted(ctx context.Context, id uint64) error {
	now := time.Now()
	return r.Model(ctx).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":       entity.OutboxStatusCompleted,
			"processed_at": now,
			"updated_at":   now,
		}).Error
}

func (r *outboxGormRepository) MarkAsFailed(ctx context.Context, id uint64, errorMsg string) error {
	return r.Model(ctx).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":        entity.OutboxStatusFailed,
			"error_message": errorMsg,
			"updated_at":    time.Now(),
		}).Error
}

func (r *outboxGormRepository) IncrementRetry(ctx context.Context, id uint64) error {
	return r.Model(ctx).
		Where("id = ?", id).
		Update("retry_count", gorm.Expr("retry_count + 1")).Error
}
