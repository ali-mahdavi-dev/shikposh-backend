package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/pkg/framework/adapter"

	"gorm.io/gorm"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	adapter.BaseRepository[*entity.Product]
	GetAll(ctx context.Context) ([]*entity.Product, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Product, error)
	FindByCategoryID(ctx context.Context, categoryID uint64) ([]*entity.Product, error)
	FindByCategorySlug(ctx context.Context, categorySlug string) ([]*entity.Product, error)
	FindFeatured(ctx context.Context) ([]*entity.Product, error)
	Search(ctx context.Context, query string) ([]*entity.Product, error)
	Filter(ctx context.Context, filters ProductFilters) ([]*entity.Product, error)
}

type ProductFilters struct {
	Query    *string
	Category *string
	MinPrice *float64
	MaxPrice *float64
	Rating   *float64
	Featured *bool
	Tags     []string
	Sort     *string
}

type productGormRepository struct {
	adapter.BaseRepository[*entity.Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.Product](db),
		db:             db,
	}
}

func (r *productGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.Product{}).
		Preload("Category").
		Preload("Details").
		Preload("Details.Images").
		Preload("Features", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Specs", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Images")
}

func (r *productGormRepository) GetAll(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := r.Model(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) FindBySlug(ctx context.Context, slug string) (*entity.Product, error) {
	var product entity.Product
	err := r.Model(ctx).Where("slug = ?", slug).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	r.SetSeen(&product)
	return &product, nil
}

func (r *productGormRepository) FindByCategoryID(ctx context.Context, categoryID uint64) ([]*entity.Product, error) {
	var products []*entity.Product
	err := r.Model(ctx).Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) FindByCategorySlug(ctx context.Context, categorySlug string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := r.Model(ctx).
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("categories.slug = ?", categorySlug).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) FindFeatured(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := r.Model(ctx).Where("is_featured = ?", true).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) Search(ctx context.Context, query string) ([]*entity.Product, error) {
	var products []*entity.Product
	searchPattern := "%" + query + "%"
	err := r.Model(ctx).
		Where("name ILIKE ? OR description ILIKE ? OR brand ILIKE ?", searchPattern, searchPattern, searchPattern).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) Filter(ctx context.Context, filters ProductFilters) ([]*entity.Product, error) {
	query := r.Model(ctx)

	if filters.Query != nil && *filters.Query != "" {
		searchPattern := "%" + *filters.Query + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ? OR brand ILIKE ?", searchPattern, searchPattern, searchPattern)
	}

	if filters.Category != nil && *filters.Category != "" {
		query = query.Joins("JOIN categories ON products.category_id = categories.id").
			Where("categories.slug = ?", *filters.Category)
	}

	if filters.MinPrice != nil {
		query = query.Where("price >= ?", *filters.MinPrice)
	}

	if filters.MaxPrice != nil {
		query = query.Where("price <= ?", *filters.MaxPrice)
	}

	if filters.Rating != nil {
		query = query.Where("rating >= ?", *filters.Rating)
	}

	if filters.Featured != nil && *filters.Featured {
		query = query.Where("is_featured = ?", true)
	}

	if len(filters.Tags) > 0 {
		for _, tag := range filters.Tags {
			query = query.Where("tags @> ?", `["`+tag+`"]`)
		}
	}

	// Apply sorting
	if filters.Sort != nil {
		switch *filters.Sort {
		case "price_asc":
			query = query.Order("price ASC")
		case "price_desc":
			query = query.Order("price DESC")
		case "rating":
			query = query.Order("rating DESC")
		case "newest":
			query = query.Order("created_at DESC")
		default:
			query = query.Order("created_at DESC")
		}
	} else {
		query = query.Order("created_at DESC")
	}

	var products []*entity.Product
	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}
