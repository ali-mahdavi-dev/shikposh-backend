package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/seller/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrSellerNotFound = errors.New("seller not found")

type SellerRepository interface {
	adapter.BaseRepository[*entity.Seller]
	GetAll(ctx context.Context) ([]*entity.Seller, error)
	FindByCategory(ctx context.Context, category string) ([]*entity.Seller, error)
	Search(ctx context.Context, query string) ([]*entity.Seller, error)
}

type sellerGormRepository struct {
	adapter.BaseRepository[*entity.Seller]
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) SellerRepository {
	return &sellerGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.Seller](db),
		db:             db,
	}
}

func (r *sellerGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.Seller{})
}

func (r *sellerGormRepository) GetAll(ctx context.Context) ([]*entity.Seller, error) {
	var sellers []*entity.Seller
	err := r.Model(ctx).Find(&sellers).Error
	if err != nil {
		return nil, err
	}
	for _, s := range sellers {
		r.SetSeen(s)
	}
	return sellers, nil
}

func (r *sellerGormRepository) FindByCategory(ctx context.Context, category string) ([]*entity.Seller, error) {
	var sellers []*entity.Seller
	// Search in JSONB array using PostgreSQL JSONB operators
	err := r.Model(ctx).
		Where("categories @> ?", `["`+category+`"]`).
		Find(&sellers).Error
	if err != nil {
		return nil, err
	}
	for _, s := range sellers {
		r.SetSeen(s)
	}
	return sellers, nil
}

func (r *sellerGormRepository) Search(ctx context.Context, query string) ([]*entity.Seller, error) {
	var sellers []*entity.Seller
	searchPattern := "%" + query + "%"
	err := r.Model(ctx).
		Where("name ILIKE ? OR description ILIKE ?", searchPattern, searchPattern).
		Find(&sellers).Error
	if err != nil {
		return nil, err
	}
	for _, s := range sellers {
		r.SetSeen(s)
	}
	return sellers, nil
}
