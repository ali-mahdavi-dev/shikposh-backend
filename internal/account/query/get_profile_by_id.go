package query

import (
	"context"

	"shikposh-backend/internal/account/domain/entity"
)

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
