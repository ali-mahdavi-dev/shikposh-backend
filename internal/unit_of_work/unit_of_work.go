package unitofwork

import (
	"context"

	"gorm.io/gorm"

	accountrepository "shikposh-backend/internal/account/adapter/repository"
	productrepository "shikposh-backend/internal/products/adapter/repository"
	sellerrepository "shikposh-backend/internal/seller/adapter/repository"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
)

// PGUnitOfWork extends the base UnitOfWork with PostgreSQL-specific functionality.
type PGUnitOfWork interface {
	adapter.UnitOfWork
	// account repositories
	User(ctx context.Context) accountrepository.UserRepository
	Token(ctx context.Context) accountrepository.TokenRepository
	Profile(ctx context.Context) accountrepository.ProfileRepository

	// product repositories
	Product(ctx context.Context) productrepository.ProductRepository
	Category(ctx context.Context) productrepository.CategoryRepository
	Tag(ctx context.Context) productrepository.TagRepository
	Size(ctx context.Context) productrepository.SizeRepository
	Outbox(ctx context.Context) productrepository.OutboxRepository
	Wishlist(ctx context.Context) productrepository.WishlistRepository

	// seller repositories
	Seller(ctx context.Context) sellerrepository.SellerRepository
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
func (uow *pgUnitOfWork) User(ctx context.Context) accountrepository.UserRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "user", func(session *gorm.DB) adapter.SeenedRepository {
		return accountrepository.NewUserRepository(session)
	}).(accountrepository.UserRepository)
}

// Token returns the TokenRepository instance for the current transaction.
func (uow *pgUnitOfWork) Token(ctx context.Context) accountrepository.TokenRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "token", func(session *gorm.DB) adapter.SeenedRepository {
		return accountrepository.NewTokenRepository(session)
	}).(accountrepository.TokenRepository)
}

// Profile returns the ProfileRepository instance for the current transaction.
func (uow *pgUnitOfWork) Profile(ctx context.Context) accountrepository.ProfileRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "profile", func(session *gorm.DB) adapter.SeenedRepository {
		return accountrepository.NewProfileRepository(session)
	}).(accountrepository.ProfileRepository)
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

// Tag returns the TagRepository instance for the current transaction.
func (uow *pgUnitOfWork) Tag(ctx context.Context) productrepository.TagRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "tag", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewTagRepository(session)
	}).(productrepository.TagRepository)
}

// Size returns the SizeRepository instance for the current transaction.
func (uow *pgUnitOfWork) Size(ctx context.Context) productrepository.SizeRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "size", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewSizeRepository(session)
	}).(productrepository.SizeRepository)
}

// Outbox returns the OutboxRepository instance for the current transaction.
func (uow *pgUnitOfWork) Outbox(ctx context.Context) productrepository.OutboxRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "outbox", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewOutboxRepository(session)
	}).(productrepository.OutboxRepository)
}

// Wishlist returns the WishlistRepository instance for the current transaction.
func (uow *pgUnitOfWork) Wishlist(ctx context.Context) productrepository.WishlistRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "wishlist", func(session *gorm.DB) adapter.SeenedRepository {
		return productrepository.NewWishlistRepository(session)
	}).(productrepository.WishlistRepository)
}

// Seller returns the SellerRepository instance for the current transaction.
func (uow *pgUnitOfWork) Seller(ctx context.Context) sellerrepository.SellerRepository {
	return uow.BaseUnitOfWork.GetOrCreateRepository(ctx, "seller", func(session *gorm.DB) adapter.SeenedRepository {
		return sellerrepository.NewSellerRepository(session)
	}).(sellerrepository.SellerRepository)
}
