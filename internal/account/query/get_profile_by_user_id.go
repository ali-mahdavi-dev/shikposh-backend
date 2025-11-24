package query

import (
	"context"

	"shikposh-backend/internal/account/domain/entity"
)

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

