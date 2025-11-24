package phrases

import (
	"github.com/ali-mahdavi-dev/shikposh-framework/errors/phrases"
)

// Products module error phrases
const (
	CategoryNotFound  phrases.MessagePhrase = "Products.Category.NotFound"
	ProductNotFound   phrases.MessagePhrase = "Products.Product.NotFound"
	ProductSlugExists phrases.MessagePhrase = "Products.Product.SlugExists"
)

// RegisterProductsPhrases registers error phrases for the products module
func RegisterProductsPhrases() {
	registry := phrases.GetRegistry()
	registry.Register(map[phrases.Language]map[phrases.MessagePhrase]string{
		phrases.Fa: {
			CategoryNotFound:  "دسته‌بندی پیدا نشد",
			ProductNotFound:   "محصول پیدا نشد",
			ProductSlugExists: "محصولی با این slug از قبل وجود دارد",
		},
		phrases.En: {
			CategoryNotFound:  "Category not found",
			ProductNotFound:   "Product not found",
			ProductSlugExists: "Product with this slug already exists",
		},
	})
}
