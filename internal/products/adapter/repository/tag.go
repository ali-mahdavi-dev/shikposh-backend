package repository

import (
	"context"
	"errors"
	"strings"

	"shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrTagNotFound = errors.New("tag not found")

type TagRepository interface {
	adapter.BaseRepository[*product_aggregate.Tag]
	GetAll(ctx context.Context) ([]*product_aggregate.Tag, error)
	FindByName(ctx context.Context, name string) (*product_aggregate.Tag, error)
	FindOrCreateByName(ctx context.Context, name string) (*product_aggregate.Tag, error)
	FindByNames(ctx context.Context, names []string) ([]*product_aggregate.Tag, error)
}

type tagGormRepository struct {
	adapter.BaseRepository[*product_aggregate.Tag]
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagGormRepository{
		BaseRepository: adapter.NewGormRepository[*product_aggregate.Tag](db),
		db:             db,
	}
}

func (r *tagGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&product_aggregate.Tag{})
}

func (r *tagGormRepository) GetAll(ctx context.Context) ([]*product_aggregate.Tag, error) {
	var tags []*product_aggregate.Tag
	err := r.Model(ctx).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	for _, t := range tags {
		r.SetSeen(t)
	}
	return tags, nil
}

func (r *tagGormRepository) FindByName(ctx context.Context, name string) (*product_aggregate.Tag, error) {
	tag, err := r.FindByField(ctx, "name", name)
	if err != nil {
		if errors.Is(err, adapter.ErrEntityNotFound) {
			return nil, ErrTagNotFound
		}
		return nil, err
	}
	return tag, nil
}

// generateSlug creates a simple slug from tag name
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	// Remove special characters (keep only alphanumeric and hyphens)
	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return strings.Trim(result.String(), "-")
}

func (r *tagGormRepository) FindOrCreateByName(ctx context.Context, name string) (*product_aggregate.Tag, error) {
	// Try to find existing tag by name
	tag, err := r.FindByName(ctx, name)
	if err == nil {
		return tag, nil
	}
	if !errors.Is(err, ErrTagNotFound) {
		return nil, err
	}

	// Tag doesn't exist, create it
	slug := generateSlug(name)
	tag = &product_aggregate.Tag{
		Name: name,
		Slug: slug,
	}

	if err := r.Save(ctx, tag); err != nil {
		// If slug conflict, try to find by slug
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			existingTag, findErr := r.FindByField(ctx, "slug", slug)
			if findErr == nil {
				return existingTag, nil
			}
		}
		return nil, err
	}

	return tag, nil
}

func (r *tagGormRepository) FindByNames(ctx context.Context, names []string) ([]*product_aggregate.Tag, error) {
	if len(names) == 0 {
		return []*product_aggregate.Tag{}, nil
	}

	var tags []*product_aggregate.Tag
	err := r.Model(ctx).Where("name IN ?", names).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	for _, t := range tags {
		r.SetSeen(t)
	}

	return tags, nil
}
