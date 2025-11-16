package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/products/domain/entity"
	"github.com/shikposh/framework/adapter"

	"gorm.io/gorm"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepository interface {
	adapter.BaseRepository[*entity.Category]
	GetAll(ctx context.Context) ([]*entity.Category, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Category, error)
}

type categoryGormRepository struct {
	adapter.BaseRepository[*entity.Category]
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.Category](db),
		db:             db,
	}
}

func (r *categoryGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.Category{})
}

func (r *categoryGormRepository) GetAll(ctx context.Context) ([]*entity.Category, error) {
	var categories []*entity.Category
	err := r.Model(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	for _, c := range categories {
		r.SetSeen(c)
	}
	return categories, nil
}

func (r *categoryGormRepository) FindBySlug(ctx context.Context, slug string) (*entity.Category, error) {
	category, err := r.FindByField(ctx, "slug", slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return category, nil
}
