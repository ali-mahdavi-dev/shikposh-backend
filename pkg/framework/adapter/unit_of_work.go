package adapter

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"shikposh-backend/pkg/framework/service_layer/types"
)

type txKey struct{}

type UnitOfWork interface {
	Do(ctx context.Context, fc types.UowUseCase) error
	GetSession(ctx context.Context) *gorm.DB
	Commit() error
	Rollback() error
}

type EventWithWaitGroup struct {
	Event interface{}
	Ctx   context.Context
	Wg    *sync.WaitGroup
}

type BaseUnitOfWork struct {
	db           *gorm.DB
	repositories map[string]SeenedRepository
	ctxMap       map[context.Context]context.Context
	eventCh      chan<- EventWithWaitGroup
	mu           sync.RWMutex
}

func NewBaseUnitOfWork(db *gorm.DB, eventCh chan<- EventWithWaitGroup) UnitOfWork {
	return &BaseUnitOfWork{
		db:           db,
		repositories: make(map[string]SeenedRepository),
		ctxMap:       make(map[context.Context]context.Context),
		eventCh:      eventCh,
	}
}

func (uow *BaseUnitOfWork) GetSession(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return uow.db
}

func (uow *BaseUnitOfWork) Do(ctx context.Context, fc types.UowUseCase) error {
	uow.clearRepositories()

	if ctx.Value(txKey{}) != nil {
		return fc(ctx)
	}

	return uow.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Store transaction in context so GetSession can retrieve it
		txCtx := context.WithValue(ctx, txKey{}, tx)
		err := fc(txCtx)
		if err != nil {
			return err
		}

		uow.mu.RLock()
		repos := make([]SeenedRepository, 0, len(uow.repositories))
		for _, repo := range uow.repositories {
			repos = append(repos, repo)
		}
		uow.mu.RUnlock()

		if len(repos) == 0 {
			return nil
		}

		var wg sync.WaitGroup
		for _, repo := range repos {
			entities := repo.Seen()
			for _, entity := range entities {
				events := entity.Event()
				for _, event := range events {
					wg.Add(1)
					select {
					case uow.eventCh <- EventWithWaitGroup{Event: event, Ctx: txCtx, Wg: &wg}:
						// Event sent with WaitGroup and transaction context, will be done when handled
					case <-txCtx.Done():
						wg.Done()
						return txCtx.Err()
					}
				}
			}
		}
		wg.Wait()
		uow.clearRepositories()
		return nil
	})
}

func (uow *BaseUnitOfWork) clearRepositories() {
	uow.mu.Lock()
	defer uow.mu.Unlock()
	uow.repositories = make(map[string]SeenedRepository)
}

func (uow *BaseUnitOfWork) GetOrCreateRepository(
	ctx context.Context,
	key string,
	factory func(*gorm.DB) SeenedRepository,
) SeenedRepository {
	uow.mu.RLock()
	if repo, ok := uow.repositories[key]; ok {
		uow.mu.RUnlock()
		return repo
	}
	uow.mu.RUnlock()

	// Create new repository instance
	session := uow.GetSession(ctx)
	repo := factory(session)
	uow.repositories[key] = repo

	return repo
}

func (uow *BaseUnitOfWork) Commit() error {
	return uow.db.Commit().Error
}

func (uow *BaseUnitOfWork) Rollback() error {
	return uow.db.Rollback().Error
}
