package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/product/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrWishlistNotFound = errors.New("wishlist item not found")

type WishlistRepository interface {
	adapter.BaseRepository[*entity.Wishlist]
	Add(ctx context.Context, wishlist *entity.Wishlist) error
	FindByUserID(ctx context.Context, userID uint64) ([]*entity.Wishlist, error)
	FindByUserAndProduct(ctx context.Context, userID uint64, productID uint64) (*entity.Wishlist, error)
	DeleteByUserAndProduct(ctx context.Context, userID uint64, productID uint64) error
	GetProductIDs(ctx context.Context, userID uint64) ([]uint64, error)
}

type wishlistGormRepository struct {
	adapter.BaseRepository[*entity.Wishlist]
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.Wishlist](db),
		db:             db,
	}
}

func (r *wishlistGormRepository) Add(ctx context.Context, wishlist *entity.Wishlist) error {
	return r.db.Create(wishlist).Error
}

func (r *wishlistGormRepository) FindByUserID(ctx context.Context, userID uint64) ([]*entity.Wishlist, error) {
	var items []*entity.Wishlist
	if err := r.db.Where("user_id = ?", userID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *wishlistGormRepository) FindByUserAndProduct(ctx context.Context, userID uint64, productID uint64) (*entity.Wishlist, error) {
	var item entity.Wishlist
	if err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWishlistNotFound
		}
		return nil, err
	}
	return &item, nil
}

func (r *wishlistGormRepository) DeleteByUserAndProduct(ctx context.Context, userID uint64, productID uint64) error {
	result := r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&entity.Wishlist{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrWishlistNotFound
	}
	return nil
}

func (r *wishlistGormRepository) GetProductIDs(ctx context.Context, userID uint64) ([]uint64, error) {
	var productIDs []uint64
	if err := r.db.Model(&entity.Wishlist{}).Where("user_id = ?", userID).Pluck("product_id", &productIDs).Error; err != nil {
		return nil, err
	}
	return productIDs, nil
}
