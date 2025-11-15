package builders

import (
	"fmt"
	"io/fs"
	"net/url"
	"os"

	accountMigrations "shikposh-backend/internal/account/adapter/migrations"
	productsMigrations "shikposh-backend/internal/products/adapter/migrations"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

func getTestDBConfig() dbConfig {
	return dbConfig{
		host:     getEnvOrDefault("TEST_DB_HOST", "localhost"),
		port:     getEnvOrDefault("TEST_DB_PORT", "5433"),
		user:     getEnvOrDefault("TEST_DB_USER", "postgres"),
		password: getEnvOrDefault("TEST_DB_PASSWORD", "admin"),
		name:     getEnvOrDefault("TEST_DB_NAME", "shikposh_test"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func connectToPostgres(config dbConfig) (*gorm.DB, error) {
	dsn := buildDSN(config.host, config.port, config.user, config.password, "postgres")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func buildDSN(host, port, user, password, dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tehran",
		host, port, user, password, dbName)
}

func runMigrations(host, port, user, password, dbName string) error {
	cnn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Tehran",
		user, password, host, port, dbName)
	u, err := url.Parse(cnn)
	if err != nil {
		return fmt.Errorf("invalid DB connection string: %w", err)
	}

	dbConn := dbmate.New(u)
	combinedFS := combineFS(accountMigrations.Migrations, productsMigrations.Migrations)
	dbConn.FS = combinedFS
	dbConn.MigrationsDir = []string{"./"}
	dbConn.AutoDumpSchema = false

	return dbConn.CreateAndMigrate()
}

func combineFS(fsList ...fs.FS) fs.FS {
	return &combinedFS{filesystems: fsList}
}

type combinedFS struct {
	filesystems []fs.FS
}

func (c *combinedFS) Open(name string) (fs.File, error) {
	for _, filesystem := range c.filesystems {
		if file, err := filesystem.Open(name); err == nil {
			return file, nil
		}
	}
	return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
}

func (c *combinedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	var allEntries []fs.DirEntry
	seen := make(map[string]bool)

	for _, filesystem := range c.filesystems {
		entries, err := fs.ReadDir(filesystem, name)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if !seen[entry.Name()] {
				allEntries = append(allEntries, entry)
				seen[entry.Name()] = true
			}
		}
	}

	if len(allEntries) == 0 {
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: fs.ErrNotExist}
	}

	return allEntries, nil
}

