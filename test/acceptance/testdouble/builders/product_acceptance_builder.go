package builders

import (
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/service_layer/command_handler"
	appadapter "shikposh-backend/pkg/framework/adapter"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProductAcceptanceTestBuilder helps build acceptance test scenarios for products
type ProductAcceptanceTestBuilder struct {
	DB  *gorm.DB
	UOW unit_of_work.PGUnitOfWork
}

func NewProductAcceptanceTestBuilder() *ProductAcceptanceTestBuilder {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	err = db.AutoMigrate(
		&entity.Category{},
		&productaggregate.Product{},
		&productaggregate.ProductFeature{},
		&productaggregate.ProductDetail{},
		&productaggregate.ProductSpec{},
	)
	Expect(err).NotTo(HaveOccurred())

	eventCh := make(chan appadapter.EventWithWaitGroup, 100)
	uow := unit_of_work.New(db, eventCh)

	return &ProductAcceptanceTestBuilder{
		DB:  db,
		UOW: uow,
	}
}

func (b *ProductAcceptanceTestBuilder) BuildHandler() *command_handler.ProductCommandHandler {
	return command_handler.NewProductCommandHandler(b.UOW)
}

func (b *ProductAcceptanceTestBuilder) Cleanup() {
	b.DB.Exec("DELETE FROM products")
	b.DB.Exec("DELETE FROM categories")
	b.DB.Exec("DELETE FROM product_features")
	b.DB.Exec("DELETE FROM product_details")
	b.DB.Exec("DELETE FROM product_specs")
}
