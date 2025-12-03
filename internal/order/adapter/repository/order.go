package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/order/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrOrderNotFound = errors.New("order not found")

type OrderRepository interface {
	adapter.BaseRepository[*entity.Order]
	GetByUserID(ctx context.Context, userID uint64, filters OrderFilters) ([]*entity.Order, error)
	GetByOrderNumber(ctx context.Context, orderNumber string) (*entity.Order, error)
	GetByIDWithItems(ctx context.Context, id uint64) (*entity.Order, error)
}

type OrderFilters struct {
	Status *entity.OrderStatus
	Limit  *int
	Offset *int
}

type orderGormRepository struct {
	adapter.BaseRepository[*entity.Order]
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.Order](db),
		db:             db,
	}
}

func (r *orderGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.Order{})
}

func (r *orderGormRepository) FindByID(ctx context.Context, id uint64) (*entity.Order, error) {
	var order entity.Order
	err := r.withPreloads(r.Model(ctx)).Where("id = ?", id).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, adapter.ErrEntityNotFound
		}
		return nil, err
	}
	r.SetSeen(&order)
	return &order, nil
}

func (r *orderGormRepository) GetByIDWithItems(ctx context.Context, id uint64) (*entity.Order, error) {
	var order entity.Order
	err := r.withPreloads(r.Model(ctx)).Where("id = ?", id).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return &order, nil
}

func (r *orderGormRepository) GetByUserID(ctx context.Context, userID uint64, filters OrderFilters) ([]*entity.Order, error) {
	var orders []*entity.Order
	query := r.withPreloads(r.Model(ctx)).Where("user_id = ?", userID)

	if filters.Status != nil {
		query = query.Where("status = ?", *filters.Status)
	}

	query = query.Order("created_at DESC")

	if filters.Limit != nil {
		query = query.Limit(*filters.Limit)
	}
	if filters.Offset != nil {
		query = query.Offset(*filters.Offset)
	}

	err := query.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		r.SetSeen(order)
	}

	return orders, nil
}

func (r *orderGormRepository) GetByOrderNumber(ctx context.Context, orderNumber string) (*entity.Order, error) {
	var order entity.Order
	err := r.withPreloads(r.Model(ctx)).Where("order_number = ?", orderNumber).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	r.SetSeen(&order)
	return &order, nil
}

func (r *orderGormRepository) withPreloads(query *gorm.DB) *gorm.DB {
	return query.
		Preload("Items").
		Preload("ShippingAddress")
}

