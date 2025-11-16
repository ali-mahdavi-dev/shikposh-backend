package builders

import (
	"fmt"
	"time"

	"shikposh-backend/internal/products/service_layer/command_handler"
	appadapter "github.com/ali-mahdavi-dev/framework/adapter"
	"shikposh-backend/internal/unit_of_work"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProductIntegrationTestBuilder helps build integration test scenarios with PostgreSQL
type ProductIntegrationTestBuilder struct {
	DB  *gorm.DB
	UOW unitofwork.PGUnitOfWork
}

func NewProductIntegrationTestBuilder() (*ProductIntegrationTestBuilder, error) {
	dbConfig := getTestDBConfig()
	testDBName := fmt.Sprintf("%s_%d", dbConfig.name, time.Now().UnixNano())

	adminDB, err := connectToPostgres(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	err = adminDB.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName)).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create test database: %w", err)
	}

	sqlDB, _ := adminDB.DB()
	sqlDB.Close()

	testDSN := buildDSN(dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, testDBName)
	db, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %w", err)
	}

	err = runMigrations(dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, testDBName)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	eventCh := make(chan appadapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)

	return &ProductIntegrationTestBuilder{
		DB:  db,
		UOW: uow,
	}, nil
}

func (b *ProductIntegrationTestBuilder) BuildHandler() *command_handler.ProductCommandHandler {
	return command_handler.NewProductCommandHandler(b.UOW)
}

func (b *ProductIntegrationTestBuilder) Cleanup() {
	var dbName string
	b.DB.Raw("SELECT current_database()").Scan(&dbName)

	sqlDB, _ := b.DB.DB()
	sqlDB.Close()

	config := getTestDBConfig()
	adminDB, err := connectToPostgres(config)
	if err != nil {
		return
	}

	adminDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
	adminSQLDB, _ := adminDB.DB()
	adminSQLDB.Close()
}
