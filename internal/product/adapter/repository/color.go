package repository

import (
	"context"
	"errors"
	"shikposh-backend/internal/product/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrColorNotFound = errors.New("color not found")

type ColorRepository interface {
	adapter.BaseRepository[*product_aggregate.Color]
	GetAll(ctx context.Context) ([]*product_aggregate.Color, error)
	FindByName(ctx context.Context, name string) (*product_aggregate.Color, error)
	FindByIDs(ctx context.Context, ids []uint64) ([]*product_aggregate.Color, error)
}

type colorGormRepository struct {
	adapter.BaseRepository[*product_aggregate.Color]
	db *gorm.DB
}

func NewColorRepository(db *gorm.DB) ColorRepository {
	return &colorGormRepository{
		BaseRepository: adapter.NewGormRepository[*product_aggregate.Color](db),
		db:             db,
	}
}

func (r *colorGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&product_aggregate.Color{})
}

func (r *colorGormRepository) GetAll(ctx context.Context) ([]*product_aggregate.Color, error) {
	var colors []*product_aggregate.Color
	err := r.Model(ctx).Find(&colors).Error
	if err != nil {
		return nil, err
	}
	for _, c := range colors {
		r.SetSeen(c)
	}
	return colors, nil
}

func (r *colorGormRepository) FindByName(ctx context.Context, name string) (*product_aggregate.Color, error) {
	color, err := r.FindByField(ctx, "name", name)
	if err != nil {
		if errors.Is(err, adapter.ErrEntityNotFound) {
			return nil, ErrColorNotFound
		}
		return nil, err
	}
	return color, nil
}

func (r *colorGormRepository) FindByIDs(ctx context.Context, ids []uint64) ([]*product_aggregate.Color, error) {
	if len(ids) == 0 {
		return []*product_aggregate.Color{}, nil
	}

	var colors []*product_aggregate.Color
	err := r.Model(ctx).Where("id IN ?", ids).Find(&colors).Error
	if err != nil {
		return nil, err
	}

	for _, c := range colors {
		r.SetSeen(c)
	}

	return colors, nil
}
