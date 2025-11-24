package query

import (
	"context"

	"shikposh-backend/internal/account/domain/entity"
)

func (h *UserQueryHandler) GetUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	var user *entity.User
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		user, err = h.uow.User(ctx).FindByID(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}
