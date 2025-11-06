package adapter

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"shikposh-backend/pkg/framework/infrastructure/logging"
	"shikposh-backend/pkg/framework/service_layer/types"
)

// txKey is a key type for storing transaction in context
type txKey struct{}

// UnitOfWork defines the interface for managing database transactions
type UnitOfWork interface {
	Do(ctx context.Context, fc types.UowUseCase) error
	GetSession(ctx context.Context) *gorm.DB
	CollectNewEvents(ctx context.Context, eventCh chan<- any)
	Commit() error
	Rollback() error
}

// BaseUnitOfWork implements the UnitOfWork pattern with GORM
type BaseUnitOfWork struct {
	db           *gorm.DB
	repositories map[string]SeenedRepository
	mu           sync.RWMutex
}

// NewBaseUnitOfWork creates a new instance of BaseUnitOfWork
func NewBaseUnitOfWork(db *gorm.DB) UnitOfWork {
	return &BaseUnitOfWork{
		db:           db,
		repositories: make(map[string]SeenedRepository),
	}
}

// GetSession returns the database session from context if available (transaction),
// otherwise returns the base database connection
func (uow *BaseUnitOfWork) GetSession(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return uow.db
}

// Do executes a function within a database transaction
// If the function returns an error, the transaction will be rolled back
func (uow *BaseUnitOfWork) Do(ctx context.Context, fc types.UowUseCase) error {
	// Clear repositories before starting new transaction
	uow.clearRepositories()
	
	return uow.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Store transaction in context so GetSession can retrieve it
		txCtx := context.WithValue(ctx, txKey{}, tx)
		return fc(txCtx)
	})
}

// CollectNewEvents collects domain events from repositories and sends them to the channel.
// Note: This should be called after transaction is committed, so repositories are cleared
// and entities are collected before the transaction ends.
func (uow *BaseUnitOfWork) CollectNewEvents(ctx context.Context, eventCh chan<- any) {
	uow.mu.RLock()
	repos := make([]SeenedRepository, 0, len(uow.repositories))
	for _, repo := range uow.repositories {
		repos = append(repos, repo)
	}
	uow.mu.RUnlock()

	if len(repos) == 0 {
		return
	}

	// Collect all entities before accessing their events
	// This must be done while repositories still have access to entities
	allEntities := make([]Entity, 0, 10)
	for _, repo := range repos {
		entities := repo.Seen()
		allEntities = append(allEntities, entities...)
	}

	// Clear repositories immediately after collecting entities
	// This prevents using repositories with closed transactions
	uow.clearRepositories()

	// Now process events from collected entities
	var wg sync.WaitGroup
	for _, entity := range allEntities {
		wg.Add(1)
		go func(e Entity) {
			defer wg.Done()
			events := e.Event()
			for _, event := range events {
				select {
				case eventCh <- event:
				case <-ctx.Done():
					logging.Warn("Context cancelled while collecting events").
						WithError(ctx.Err()).
						Log()
					return
				}
			}
		}(entity)
	}

	wg.Wait()
}

// clearRepositories clears all cached repositories after event collection.
func (uow *BaseUnitOfWork) clearRepositories() {
	uow.mu.Lock()
	defer uow.mu.Unlock()
	uow.repositories = make(map[string]SeenedRepository)
}

// GetOrCreateRepository returns a cached repository or creates a new one using double-check locking.
func (uow *BaseUnitOfWork) GetOrCreateRepository(
	ctx context.Context,
	key string,
	factory func(*gorm.DB) SeenedRepository,
) SeenedRepository {
	// Fast path: check with read lock
	uow.mu.RLock()
	if repo, ok := uow.repositories[key]; ok {
		uow.mu.RUnlock()
		return repo
	}
	uow.mu.RUnlock()

	// Slow path: acquire write lock and create repository
	uow.mu.Lock()
	defer uow.mu.Unlock()

	// Double-check after acquiring write lock
	if repo, ok := uow.repositories[key]; ok {
		return repo
	}

	// Create new repository instance
	session := uow.GetSession(ctx)
	repo := factory(session)
	uow.repositories[key] = repo

	return repo
}

// Commit commits the current transaction (if in manual mode)
// Note: With Transaction() method, commit is automatic on success
func (uow *BaseUnitOfWork) Commit() error {
	return uow.db.Commit().Error
}

// Rollback rolls back the current transaction (if in manual mode)
// Note: With Transaction() method, rollback is automatic on error
func (uow *BaseUnitOfWork) Rollback() error {
	return uow.db.Rollback().Error
}
