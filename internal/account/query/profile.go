package query

import (
	"context"

	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type ProfileQueryHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewProfileQueryHandler(uow unit_of_work.PGUnitOfWork) *ProfileQueryHandler {
	return &ProfileQueryHandler{uow: uow}
}

func (h *ProfileQueryHandler) GetProfileByID(ctx context.Context, id uint64) (*entity.Profile, error) {
	var profile *entity.Profile
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		profile, err = h.uow.Profile(ctx).FindByID(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return profile, err
}

func (h *ProfileQueryHandler) GetProfileByUserID(ctx context.Context, userID entity.UserID) (*entity.Profile, error) {
	var profile *entity.Profile
	err := h.uow.Do(ctx, func(ctx context.Context) error {
		var err error
		profile, err = h.uow.Profile(ctx).FindByUserID(ctx, userID)
		if err != nil {
			return err
		}
		return nil
	})
	return profile, err
}
