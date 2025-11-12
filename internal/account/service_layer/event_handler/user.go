package event_handler

import (
	"context"
	"fmt"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/account/domain/events"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type UserEventHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewUserEventHandler(uow unit_of_work.PGUnitOfWork) *UserEventHandler {
	return &UserEventHandler{uow: uow}
}

// RegisterEvent handles the RegisterUserEvent and creates a profile for the newly registered user
func (h *UserEventHandler) RegisterEvent(ctx context.Context, event *events.RegisterUserEvent) error {
	if event.UserID == nil {
		return fmt.Errorf("UserEventHandler.RegisterEvent: UserID is nil")
	}

	userID := entity.UserID(*event.UserID)
	logging.Info("Processing RegisterUserEvent").
		WithInt64("user_id", int64(userID)).
		WithString("username", event.UserName).
		Log()

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Check if profile already exists (idempotency check)
		_, err := h.uow.Profile(ctx).FindByUserID(ctx, userID)
		if err != nil {
			if err != repository.ErrProfileNotFound {
				return fmt.Errorf("UserEventHandler.RegisterEvent error checking existing profile: %w", err)
			}
			// Profile doesn't exist, create it
		} else {
			// Profile already exists, skip creation
			logging.Info("Profile already exists, skipping creation").
				WithInt64("user_id", int64(userID)).
				Log()
			return nil
		}

		// Create new profile for the user
		profile := entity.NewProfile(userID)

		err = h.uow.Profile(ctx).Save(ctx, profile)
		if err != nil {
			return fmt.Errorf("UserEventHandler.RegisterEvent error saving profile: %w", err)
		}

		logging.Info("Profile created successfully").
			WithInt64("user_id", int64(userID)).
			WithInt64("profile_id", int64(profile.ID)).
			Log()

		return nil
	})

	if err != nil {
		logging.Error("Failed to process RegisterUserEvent").
			WithInt64("user_id", int64(userID)).
			WithError(err).
			Log()
		return fmt.Errorf("UserEventHandler.RegisterEvent fail transaction: %w", err)
	}

	return nil
}
