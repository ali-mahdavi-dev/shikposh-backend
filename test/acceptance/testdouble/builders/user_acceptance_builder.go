package builders

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/account/domain/entity"
	"shikposh-backend/internal/account/service_layer/command_handler"
	"shikposh-backend/pkg/framework/adapter"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// UserAcceptanceTestBuilder helps build acceptance test scenarios
type UserAcceptanceTestBuilder struct {
	DB  *gorm.DB
	UOW unit_of_work.PGUnitOfWork
	Cfg *config.Config
}

func NewUserAcceptanceTestBuilder() *UserAcceptanceTestBuilder {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Token{},
		&entity.Profile{},
	)
	Expect(err).NotTo(HaveOccurred())

	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unit_of_work.New(db, eventCh)

	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:                    "test-secret-for-acceptance",
			AccessTokenExpireDuration: 3600,
		},
	}

	return &UserAcceptanceTestBuilder{
		DB:  db,
		UOW: uow,
		Cfg: cfg,
	}
}

func (b *UserAcceptanceTestBuilder) BuildHandler() *command_handler.UserHandler {
	return command_handler.NewUserHandler(b.UOW, b.Cfg)
}

func (b *UserAcceptanceTestBuilder) Cleanup() {
	b.DB.Exec("DELETE FROM users")
	b.DB.Exec("DELETE FROM tokens")
	b.DB.Exec("DELETE FROM profiles")
}

