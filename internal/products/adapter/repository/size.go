package repository

import (
	"context"
	"errors"
	"strings"

	"shikposh-backend/internal/products/domain/entity/shared"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrSizeNotFound = errors.New("size not found")

type SizeRepository interface {
	adapter.BaseRepository[*shared.Size]
	FindByName(ctx context.Context, name string) (*shared.Size, error)
	FindOrCreateByName(ctx context.Context, name string) (*shared.Size, error)
	FindByNames(ctx context.Context, names []string) ([]*shared.Size, error)
}

type sizeGormRepository struct {
	adapter.BaseRepository[*shared.Size]
	db *gorm.DB
}

func NewSizeRepository(db *gorm.DB) SizeRepository {
	return &sizeGormRepository{
		BaseRepository: adapter.NewGormRepository[*shared.Size](db),
		db:             db,
	}
}

func (r *sizeGormRepository) Model(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&shared.Size{})
}

func (r *sizeGormRepository) FindByName(ctx context.Context, name string) (*shared.Size, error) {
	size, err := r.FindByField(ctx, "name", name)
	if err != nil {
		if errors.Is(err, adapter.ErrEntityNotFound) {
			return nil, ErrSizeNotFound
		}
		return nil, err
	}
	return size, nil
}

// generateSlug creates a simple slug from size name
func generateSizeSlug(name string) string {
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

func (r *sizeGormRepository) FindOrCreateByName(ctx context.Context, name string) (*shared.Size, error) {
	// Try to find existing size by name
	size, err := r.FindByName(ctx, name)
	if err == nil {
		return size, nil
	}
	if !errors.Is(err, ErrSizeNotFound) {
		return nil, err
	}

	// Size doesn't exist, create it
	slug := generateSizeSlug(name)
	size = &shared.Size{
		Name: name,
		Slug: slug,
	}

	if err := r.Save(ctx, size); err != nil {
		// If slug conflict, try to find by slug
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			existingSize, findErr := r.FindByField(ctx, "slug", slug)
			if findErr == nil {
				return existingSize, nil
			}
		}
		return nil, err
	}

	return size, nil
}

func (r *sizeGormRepository) FindByNames(ctx context.Context, names []string) ([]*shared.Size, error) {
	if len(names) == 0 {
		return []*shared.Size{}, nil
	}

	var sizes []*shared.Size
	err := r.Model(ctx).Where("name IN ?", names).Find(&sizes).Error
	if err != nil {
		return nil, err
	}

	for _, s := range sizes {
		r.SetSeen(s)
	}

	return sizes, nil
}
