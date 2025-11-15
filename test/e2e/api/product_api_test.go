package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"shikposh-backend/config"
	products "shikposh-backend/internal/products"
	"shikposh-backend/internal/products/adapter/repository"
	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"

	"github.com/gofiber/fiber/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("Product API E2E", func() {
	var (
		builder *ProductE2ETestBuilder
	)

	BeforeEach(func() {
		builder = NewProductE2ETestBuilder()
	})

	AfterEach(func() {
		builder.Cleanup()
	})

	Describe("GET /api/v1/public/products", func() {
		Context("when requesting all products", func() {
			It("should return all products via HTTP API", func() {
				req := httptest.NewRequest(http.MethodGet, "/api/v1/public/products", nil)
				resp, err := builder.app.Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				var result map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&result)
				Expect(err).NotTo(HaveOccurred())
				Expect(result["data"]).NotTo(BeNil())
			})
		})
	})

	Describe("POST /api/v1/admin/products", func() {
		Context("when creating a new product", func() {
			It("should create product via HTTP API", func() {
				// First create category
				categoryRepo := repository.NewCategoryRepository(builder.db)
				category := &entity.Category{
					Name: "Clothing",
					Slug: "clothing",
				}
				err := categoryRepo.Save(context.Background(), category)
				Expect(err).NotTo(HaveOccurred())

				desc := "Product description"
				cmd := commands.CreateProduct{
					Name:        "Men's T-Shirt",
					Brand:       "Test Brand",
					Description: &desc,
					CategoryID:  uint64(category.ID),
					Details: []commands.ProductDetailInput{
						{Price: 100000.0, Stock: 10},
					},
				}

				body, err := json.Marshal(cmd)
				Expect(err).NotTo(HaveOccurred())

				req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/products", bytes.NewBuffer(body))
				req.Header.Set("Content-Type", "application/json")

				resp, err := builder.app.Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			})
		})
	})
})

// ProductE2ETestBuilder helps build E2E test scenarios for products
type ProductE2ETestBuilder struct {
	app *fiber.App
	db  *gorm.DB
}

func NewProductE2ETestBuilder() *ProductE2ETestBuilder {
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

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	cfg := &config.Config{}

	// Bootstrap products module
	err = products.Bootstrap(app, db, cfg, nil)
	Expect(err).NotTo(HaveOccurred())

	return &ProductE2ETestBuilder{
		app: app,
		db:  db,
	}
}

func (b *ProductE2ETestBuilder) Cleanup() {
	b.db.Exec("DELETE FROM products")
	b.db.Exec("DELETE FROM categories")
	b.db.Exec("DELETE FROM product_features")
	b.db.Exec("DELETE FROM product_details")
	b.db.Exec("DELETE FROM product_specs")
}
