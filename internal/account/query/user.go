package query

import (
	"context"

	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type UserQueryHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewUserQueryHandler(uow unit_of_work.PGUnitOfWork) *UserQueryHandler {
	return &UserQueryHandler{uow: uow}
}

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

func (h *UserQueryHandler) GetUserByUserName(ctx context.Context, username string) (*entity.User, error) {
	var user *entity.User
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		user, err = h.uow.User(ctx).FindByUserName(ctx, username)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}
