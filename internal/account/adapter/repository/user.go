package repository

import (
	"context"
	"errors"

	"shikposh-backend/internal/account/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	adapter.BaseRepository[*entity.User]
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
}

type userGormRepository struct {
	adapter.BaseRepository[*entity.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGormRepository{
		BaseRepository: adapter.NewGormRepository[*entity.User](db),
		db:             db,
	}
}

func (u *userGormRepository) Model(ctx context.Context) *gorm.DB {
	return u.db.Model(&entity.User{})
}

func (u *userGormRepository) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	user, err := u.FindByField(ctx, "phone", phone)
	if err != nil {
		if errors.Is(err, adapter.ErrEntityNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}
