package unit_of_work

import (
	"context"

	"gorm.io/gorm"

	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/pkg/framework/adapter"
)

// PGUnitOfWork extends the base UnitOfWork with PostgreSQL-specific functionality.
type PGUnitOfWork interface {
	adapter.UnitOfWork
	User(ctx context.Context) repository.UserRepository
	Token(ctx context.Context) repository.TokenRepository
	Profile(ctx context.Context) repository.ProfileRepository
}

type pgUnitOfWork struct {
	*adapter.BaseUnitOfWork
	db *gorm.DB
}

// New creates a new PostgreSQL UnitOfWork instance.
func New(db *gorm.DB, eventCh chan<- adapter.EventWithWaitGroup) PGUnitOfWork {
	return &pgUnitOfWork{
		BaseUnitOfWork: adapter.NewBaseUnitOfWork(db, eventCh).(*adapter.BaseUnitOfWork),
		db:             db,
	}
}

// User returns the UserRepository instance for the current transaction.
func (uow *pgUnitOfWork) User(ctx context.Context) repository.UserRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "user", func(session *gorm.DB) adapter.SeenedRepository {
		return repository.NewUserRepository(session)
	}).(repository.UserRepository)
}

// Token returns the TokenRepository instance for the current transaction.
func (uow *pgUnitOfWork) Token(ctx context.Context) repository.TokenRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "token", func(session *gorm.DB) adapter.SeenedRepository {
		return repository.NewTokenRepository(session)
	}).(repository.TokenRepository)
}

// Profile returns the ProfileRepository instance for the current transaction.
func (uow *pgUnitOfWork) Profile(ctx context.Context) repository.ProfileRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "profile", func(session *gorm.DB) adapter.SeenedRepository {
		return repository.NewProfileRepository(session)
	}).(repository.ProfileRepository)
}
