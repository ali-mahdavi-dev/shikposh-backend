package repository

import (
	"context"
	"errors"
	"time"

	"shikposh-backend/internal/product/domain/entity"
	productaggregate "shikposh-backend/internal/product/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	adapter.BaseRepository[*productaggregate.Product]
	GetAll(ctx context.Context) ([]*productaggregate.Product, error)
	FindBySlug(ctx context.Context, slug string) (*productaggregate.Product, error)
	FindByIDIncludingDeleted(ctx context.Context, id uint64) (*productaggregate.Product, error)
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

// FindByID overrides BaseRepository FindByID to include preloads
func (r *productGormRepository) FindByID(ctx context.Context, id uint64) (*productaggregate.Product, error) {
	var product productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, adapter.ErrEntityNotFound
		}
		return nil, err
	}
	r.SetSeen(&product)
	return &product, nil
}

// FindByIDIncludingDeleted finds a product by ID including soft deleted ones
// This is useful for delete operations where we need to find the product even if it's already soft deleted
func (r *productGormRepository) FindByIDIncludingDeleted(ctx context.Context, id uint64) (*productaggregate.Product, error) {
	var product productaggregate.Product
	err := r.withPreloads(r.Model(ctx)).Unscoped().Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	r.SetSeen(&product)
	return &product, nil
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
	// Use explicit delete to avoid GORM association issues
	return r.db.WithContext(ctx).
		Where("product_id = ?", product.ID).
		Delete(&productaggregate.ProductFeature{}).Error
}

func (r *productGormRepository) ClearSpecs(ctx context.Context, product *productaggregate.Product) error {
	// Use explicit delete to avoid GORM association issues
	return r.db.WithContext(ctx).
		Where("product_id = ?", product.ID).
		Delete(&productaggregate.ProductSpec{}).Error
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
	// Use explicit delete to avoid GORM association issues
	return r.db.WithContext(ctx).
		Where("product_id = ?", product.ID).
		Delete(&productaggregate.ProductVariant{}).Error
}

func (r *productGormRepository) ClearImages(ctx context.Context, product *productaggregate.Product) error {
	// Use explicit delete to avoid GORM association issues
	return r.db.WithContext(ctx).
		Where("product_id = ?", product.ID).
		Delete(&productaggregate.ProductImage{}).Error
}

func (r *productGormRepository) ClearCategories(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).Association("Categories").Clear()
}

func (r *productGormRepository) ClearAllAssociations(ctx context.Context, product *productaggregate.Product) error {
	return r.db.WithContext(ctx).Model(product).
		Select("Features", "Specs", "Tags", "Sizes", "Colors", "Variants", "Images", "Categories").
		Delete(product).Error
}

// Remove overrides BaseRepository Remove to properly handle soft delete with WHERE condition
func (r *productGormRepository) Remove(ctx context.Context, product *productaggregate.Product, softDelete bool) error {
	r.SetSeen(product)
	if softDelete {
		now := time.Now()
		// Use Model with Where to ensure WHERE condition is set
		return r.db.WithContext(ctx).
			Model(product).
			Where("id = ?", product.ID).
			Update("deleted_at", &now).Error
	}

	// Hard delete
	return r.db.WithContext(ctx).
		Where("id = ?", product.ID).
		Delete(product).Error
}

// Modify overrides BaseRepository Modify to properly handle product updates with associations
// We need to update the product first, then handle associations separately to avoid duplicate key errors
func (r *productGormRepository) Modify(ctx context.Context, product *productaggregate.Product) error {
	// Save associations separately to avoid issues
	variants := product.Variants
	features := product.Features
	specs := product.Specs
	images := product.Images
	categories := product.Categories
	colors := product.Colors
	sizes := product.Sizes
	tags := product.Tags

	// Clear associations from product before saving
	product.Variants = nil
	product.Features = nil
	product.Specs = nil
	product.Images = nil
	product.Categories = nil
	product.Colors = nil
	product.Sizes = nil
	product.Tags = nil

	// Update product fields first (without associations)
	err := r.db.WithContext(ctx).Model(product).Updates(map[string]interface{}{
		"title":             product.Title,
		"slug":              product.Slug,
		"brand":             product.Brand,
		"description":       product.Description,
		"short_description": product.ShortDescription,
		"thumbnail":         product.Thumbnail,
		"discount":          product.Discount,
		"stock":             product.Stock,
		"origin_price":      product.OriginPrice,
		"price":             product.Price,
		"is_new":            product.IsNew,
		"is_featured":       product.IsFeatured,
		"updated_at":        product.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}

	// Restore associations
	product.Variants = variants
	product.Features = features
	product.Specs = specs
	product.Images = images
	product.Categories = categories
	product.Colors = colors
	product.Sizes = sizes
	product.Tags = tags

	// Save associations - clear first to ensure clean state, then insert
	// We clear here even though handler cleared, because we're in the same transaction
	// and want to ensure no duplicates
	if err := r.ClearVariants(ctx, product); err != nil {
		return err
	}
	if err := r.ClearFeatures(ctx, product); err != nil {
		return err
	}
	if err := r.ClearSpecs(ctx, product); err != nil {
		return err
	}
	if err := r.ClearImages(ctx, product); err != nil {
		return err
	}

	// Now insert new associations
	// Use ON CONFLICT DO UPDATE to handle duplicates gracefully
	if len(variants) > 0 {
		for i := range variants {
			variants[i].ProductID = product.ID
		}
		// Use Clauses to add ON CONFLICT handling
		if err := r.db.WithContext(ctx).
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "product_id"}, {Name: "color_id"}, {Name: "size_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"stock"}),
			}).
			Create(&variants).Error; err != nil {
			return err
		}
	}
	if len(features) > 0 {
		for i := range features {
			features[i].ProductID = product.ID
		}
		if err := r.db.WithContext(ctx).Create(&features).Error; err != nil {
			return err
		}
	}
	if len(specs) > 0 {
		for i := range specs {
			specs[i].ProductID = product.ID
		}
		// Use ON CONFLICT DO UPDATE for specs (unique constraint on product_id, key)
		if err := r.db.WithContext(ctx).
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "product_id"}, {Name: "key"}},
				DoUpdates: clause.AssignmentColumns([]string{"value", "order"}),
			}).
			Create(&specs).Error; err != nil {
			return err
		}
	}
	if len(images) > 0 {
		for i := range images {
			images[i].ProductID = product.ID
		}
		if err := r.db.WithContext(ctx).Create(&images).Error; err != nil {
			return err
		}
	}
	// For many-to-many relationships, clear first then append
	// Clear all many-to-many associations
	if err := r.ClearCategories(ctx, product); err != nil {
		return err
	}
	if err := r.ClearColors(ctx, product); err != nil {
		return err
	}
	if err := r.ClearSizes(ctx, product); err != nil {
		return err
	}
	if err := r.ClearTags(ctx, product); err != nil {
		return err
	}

	// Now append new many-to-many associations
	if len(categories) > 0 {
		if err := r.db.WithContext(ctx).Model(product).Association("Categories").Append(categories); err != nil {
			return err
		}
	}
	if len(colors) > 0 {
		if err := r.db.WithContext(ctx).Model(product).Association("Colors").Append(colors); err != nil {
			return err
		}
	}
	if len(sizes) > 0 {
		if err := r.db.WithContext(ctx).Model(product).Association("Sizes").Append(sizes); err != nil {
			return err
		}
	}
	if len(tags) > 0 {
		if err := r.db.WithContext(ctx).Model(product).Association("Tags").Append(tags); err != nil {
			return err
		}
	}

	r.SetSeen(product)
	return nil
}
