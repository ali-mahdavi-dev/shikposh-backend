package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

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
	FindFeaturedForReindex(ctx context.Context) ([]*productaggregate.Product, error)
	FindAllForReindex(ctx context.Context) ([]*productaggregate.Product, error)
	Search(ctx context.Context, query string) ([]*productaggregate.Product, error)
	Filter(ctx context.Context, filters ProductFilters) ([]*productaggregate.Product, error)
	ClearFeatures(ctx context.Context, product *productaggregate.Product) error
	ClearVariants(ctx context.Context, product *productaggregate.Product) error
	ClearSpecs(ctx context.Context, product *productaggregate.Product) error
	ClearTags(ctx context.Context, product *productaggregate.Product) error
	ClearSizes(ctx context.Context, product *productaggregate.Product) error
	ClearColors(ctx context.Context, product *productaggregate.Product) error
	ClearImages(ctx context.Context, product *productaggregate.Product) error
	ClearCategories(ctx context.Context, product *productaggregate.Product) error
	ClearAllAssociations(ctx context.Context, product *productaggregate.Product) error
}

type ProductFilters struct {
	Query        *string
	Category     *string
	CategoryName *string
	MinPrice     *int64
	MaxPrice     *int64
	Rating       *float64
	Featured     *bool
	Tags         []string
	Sort         *string
	Limit        *int
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
		Preload("Categories").
		Preload("Colors").
		Preload("Sizes").
		Preload("Variants").
		Preload("Tags").
		Preload("Features", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Specs", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		})
}

// withPreloadsWithoutImages applies preloads without Images (for reindexing)
func (r *productGormRepository) withPreloadsWithoutImages(query *gorm.DB) *gorm.DB {
	return query.
		Preload("Categories").
		Preload("Colors").
		Preload("Sizes").
		Preload("Variants").
		Preload("Tags").
		Preload("Features", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Specs", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		})
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

// FindFeaturedForReindex returns featured products without Images preload (for reindexing)
func (r *productGormRepository) FindFeaturedForReindex(ctx context.Context) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	err := r.withPreloadsWithoutImages(r.Model(ctx)).Where("is_featured = ?", true).Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		r.SetSeen(p)
	}
	return products, nil
}

// FindAllForReindex returns all products without Images preload (for reindexing)
func (r *productGormRepository) FindAllForReindex(ctx context.Context) ([]*productaggregate.Product, error) {
	var products []*productaggregate.Product
	err := r.withPreloadsWithoutImages(r.Model(ctx)).Find(&products).Error
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
		Where("title ILIKE ? OR description ILIKE ? OR brand ILIKE ?", searchPattern, searchPattern, searchPattern).
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
		query = query.Where("title ILIKE ? OR description ILIKE ? OR brand ILIKE ?", searchPattern, searchPattern, searchPattern)
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
		query = query.Joins("JOIN product_tags ON products.id = product_tags.product_id").
			Joins("JOIN tags ON product_tags.tag_id = tags.id").
			Where("tags.name IN ?", filters.Tags).
			Group("products.id")
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
		case "discount_desc":
			query = query.Order("discount DESC")
		case "newest":
			query = query.Order("created_at DESC")
		default:
			query = query.Order("created_at DESC")
		}
	} else {
		query = query.Order("created_at DESC")
	}

	if filters.Limit != nil && *filters.Limit > 0 {
		query = query.Limit(*filters.Limit)
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
	return r.db.WithContext(ctx).Model(product).Association("Features").Clear()
}

func (r *productGormRepository) ClearSpecs(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Specs").Clear()
}

func (r *productGormRepository) ClearTags(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Tags").Clear()
}

func (r *productGormRepository) ClearSizes(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Sizes").Clear()
}

func (r *productGormRepository) ClearColors(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Colors").Clear()
}

func (r *productGormRepository) ClearVariants(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Variants").Clear()
}

func (r *productGormRepository) ClearImages(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Images").Clear()
}

func (r *productGormRepository) ClearCategories(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Categories").Clear()
}

func (r *productGormRepository) ClearAllAssociations(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).
		Select("Features", "Specs", "Tags", "Sizes", "Colors", "Variants", "Images", "Categories").
		Delete(product).Error
}
