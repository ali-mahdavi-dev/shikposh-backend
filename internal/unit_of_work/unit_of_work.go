package unitofwork

import (
	"context"

	"gorm.io/gorm"

	"shikposh-backend/internal/account/adapter/repository"
	productrepository "shikposh-backend/internal/products/adapter/repository"
	"github.com/shikposh/framework/adapter"
)

// PGUnitOfWork extends the base UnitOfWork with PostgreSQL-specific functionality.
type PGUnitOfWork interface {
	adapter.UnitOfWork
	User(ctx context.Context) repository.UserRepository
	Token(ctx context.Context) repository.TokenRepository
	Profile(ctx context.Context) repository.ProfileRepository
	Product(ctx context.Context) productrepository.ProductRepository
	Category(ctx context.Context) productrepository.CategoryRepository
	Review(ctx context.Context) productrepository.ReviewRepository
	Outbox(ctx context.Context) productrepository.OutboxRepository
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

// Product returns the ProductRepository instance for the current transaction.
func (uow *pgUnitOfWork) Product(ctx context.Context) productrepository.ProductRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "product", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewProductRepository(session)
	}).(productrepository.ProductRepository)
}

// Category returns the CategoryRepository instance for the current transaction.
func (uow *pgUnitOfWork) Category(ctx context.Context) productrepository.CategoryRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "category", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewCategoryRepository(session)
	}).(productrepository.CategoryRepository)
}

// Review returns the ReviewRepository instance for the current transaction.
func (uow *pgUnitOfWork) Review(ctx context.Context) productrepository.ReviewRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "review", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewReviewRepository(session)
	}).(productrepository.ReviewRepository)
}

// Outbox returns the OutboxRepository instance for the current transaction.
func (uow *pgUnitOfWork) Outbox(ctx context.Context) productrepository.OutboxRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "outbox", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewOutboxRepository(session)
	}).(productrepository.OutboxRepository)
}
