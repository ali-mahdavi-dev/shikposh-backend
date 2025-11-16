package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"github.com/ali-mahdavi-dev/framework/adapter"

	"gorm.io/gorm"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	adapter.BaseRepository[*productaggregate.Product]
	GetAll(ctx context.Context) ([]*productaggregate.Product, error)
	FindBySlug(ctx context.Context, slug string) (*productaggregate.Product, error)
	FindByCategoryID(ctx context.Context, categoryID entity.CategoryID) ([]*productaggregate.Product, error)
	FindByCategorySlug(ctx context.Context, categorySlug string) ([]*productaggregate.Product, error)
	FindFeatured(ctx context.Context) ([]*productaggregate.Product, error)
	Search(ctx context.Context, query string) ([]*productaggregate.Product, error)
	Filter(ctx context.Context, filters ProductFilters) ([]*productaggregate.Product, error)
	ClearFeatures(ctx context.Context, product *productaggregate.Product) error
	ClearDetails(ctx context.Context, product *productaggregate.Product) error
	ClearSpecs(ctx context.Context, product *productaggregate.Product) error
	ClearAllAssociations(ctx context.Context, product *productaggregate.Product) error
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
	adapter.BaseRepository[*productaggregate.Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productGormRepository{
		BaseRepository: adapter.NewGormRepository[*productaggregate.Product](db),
		db:             db,
	}
}

func (r *productGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&productaggregate.Product{})
}

// withPreloads applies all necessary preloads to the query
func (r *productGormRepository) withPreloads(query *gorm.DB) *gorm.DB {
	return query.
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

func (r *productGormRepository) GetAll(ctx context.Context) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) FindBySlug(ctx context.Context, slug string) (*productaggregate.Product, error) {
	var product productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).Where("slug = ?", slug).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	r.SetSeen(&product)
	return &product, nil
}

func (r *productGormRepository) FindByCategoryID(ctx context.Context, categoryID entity.CategoryID) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).Where("category_id = ?", uint64(categoryID)).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) FindByCategorySlug(ctx context.Context, categorySlug string) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).
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

func (r *productGormRepository) FindFeatured(ctx context.Context) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).Where("is_featured = ?", true).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) Search(ctx context.Context, query string) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	searchPattern := "%" + query + "%"
	err := r.withPreloads(r.Model(ctx)).
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

func (r *productGormRepository) Filter(ctx context.Context, filters ProductFilters) ([]*productaggregate.Product, error) {
	query := r.withPreloads(r.Model(ctx))

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

	var products []*productaggregate.Product
	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

func (r *productGormRepository) ClearFeatures(ctx context.Context, product *productaggregate.Product) error {
	return r.Model(ctx).Association("Features").Clear()
}

func (r *productGormRepository) ClearSpecs(ctx context.Context, product *productaggregate.Product) error {
	return r.Model(ctx).Association("Specs").Clear()
}

// clearDetailsAttachments loads details and clears their attachments
func (r *productGormRepository) clearDetailsAttachments(ctx context.Context, product *productaggregate.Product) error {
	// Load existing details to delete their attachments
	if err := r.Model(ctx).Association("Details").Find(&product.Details); err != nil {
		return err
	}

	// Delete attachments for existing details
	for i := range product.Details {
		if err := r.db.WithContext(ctx).Model(&product.Details[i]).Association("Images").Clear(); err != nil {
			return err
		}
	}

	return nil
}

func (r *productGormRepository) ClearDetails(ctx context.Context, product *productaggregate.Product) error {
	// Clear attachments for details
	if err := r.clearDetailsAttachments(ctx, product); err != nil {
		return err
	}

	// Delete existing details
	return r.Model(ctx).Association("Details").Clear()
}

func (r *productGormRepository) ClearAllAssociations(ctx context.Context, product *productaggregate.Product) error {
	// Clear attachments for details
	if err := r.clearDetailsAttachments(ctx, product); err != nil {
		return err
	}

	// Delete all associations using Select
	return r.Model(ctx).Select("Features", "Details", "Specs").Delete(product).Error
}
