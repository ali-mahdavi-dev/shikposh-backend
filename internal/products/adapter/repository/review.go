package repository

import (
	"context"
	"errors"

	accountentity "shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

var ErrReviewNotFound = errors.New("review not found")

type ReviewRepository interface {
	adapter.BaseRepository[*entity.Review]
	FindByProductID(ctx context.Context, productID productaggregate.ProductID) ([]*entity.Review, error)
	FindByUserID(ctx context.Context, userID accountentity.UserID) ([]*entity.Review, error)
}

type reviewGormRepository struct {
	adapter.BaseRepository[*entity.Review]
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.Review](db),
		db:             db,
	}
}

func (r *reviewGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.Review{}).Preload("Product")
}

func (r *reviewGormRepository) FindByProductID(ctx context.Context, productID productaggregate.ProductID) ([]*entity.Review, error) {
	var reviews []*entity.Review
	err := r.Model(ctx).Where("product_id = ?", uint64(productID)).Order("created_at DESC").Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	for _, review := range reviews {
		r.SetSeen(review)
	}
	return reviews, nil
}

func (r *reviewGormRepository) FindByUserID(ctx context.Context, userID accountentity.UserID) ([]*entity.Review, error) {
	var reviews []*entity.Review
	err := r.Model(ctx).Where("user_id = ?", uint64(userID)).Order("created_at DESC").Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	for _, review := range reviews {
		r.SetSeen(review)
	}
	return reviews, nil
}
