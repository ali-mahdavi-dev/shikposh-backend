package specification

import (
	"shikposh-backend/internal/products/domain/entity"
	productaggregate "shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/pkg/framework/specification"
)

// ProductIsFeaturedSpecification checks if a product is featured
type ProductIsFeaturedSpecification struct {
}

func NewProductIsFeaturedSpecification() specification.Specification[*productaggregate.Product] {
	return &ProductIsFeaturedSpecification{}
}

func (s *ProductIsFeaturedSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	return product != nil && product.IsFeatured
}

// ProductIsNewSpecification checks if a product is new
type ProductIsNewSpecification struct {
}

func NewProductIsNewSpecification() specification.Specification[*productaggregate.Product] {
	return &ProductIsNewSpecification{}
}

func (s *ProductIsNewSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	return product != nil && product.IsNew
}

// ProductHasMinimumRatingSpecification checks if a product has a minimum rating
type ProductHasMinimumRatingSpecification struct {
	minRating float64
}

func NewProductHasMinimumRatingSpecification(minRating float64) specification.Specification[*productaggregate.Product] {
	return &ProductHasMinimumRatingSpecification{
		minRating: minRating,
	}
}

func (s *ProductHasMinimumRatingSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	return product != nil && product.Rating >= s.minRating
}

// ProductHasMinimumReviewCountSpecification checks if a product has a minimum number of reviews
type ProductHasMinimumReviewCountSpecification struct {
	minReviewCount int
}

func NewProductHasMinimumReviewCountSpecification(minReviewCount int) specification.Specification[*productaggregate.Product] {
	return &ProductHasMinimumReviewCountSpecification{
		minReviewCount: minReviewCount,
	}
}

func (s *ProductHasMinimumReviewCountSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	return product != nil && product.ReviewCount >= s.minReviewCount
}

// ProductInCategorySpecification checks if a product belongs to a specific category
type ProductInCategorySpecification struct {
	categoryID uint64
}

func NewProductInCategorySpecification(categoryID entity.CategoryID) specification.Specification[*productaggregate.Product] {
	return &ProductInCategorySpecification{
		categoryID: uint64(categoryID),
	}
}

func (s *ProductInCategorySpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	return product != nil && product.CategoryID == s.categoryID
}

// ProductHasTagSpecification checks if a product has a specific tag
type ProductHasTagSpecification struct {
	tag string
}

func NewProductHasTagSpecification(tag string) specification.Specification[*productaggregate.Product] {
	return &ProductHasTagSpecification{
		tag: tag,
	}
}

func (s *ProductHasTagSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	if product == nil || len(product.Tags) == 0 {
		return false
	}
	for _, t := range product.Tags {
		if t == s.tag {
			return true
		}
	}
	return false
}

// ProductHasAnyTagSpecification checks if a product has any of the specified tags
type ProductHasAnyTagSpecification struct {
	tags []string
}

func NewProductHasAnyTagSpecification(tags []string) specification.Specification[*productaggregate.Product] {
	return &ProductHasAnyTagSpecification{
		tags: tags,
	}
}

func (s *ProductHasAnyTagSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	if product == nil || len(product.Tags) == 0 || len(s.tags) == 0 {
		return false
	}
	for _, productTag := range product.Tags {
		for _, requiredTag := range s.tags {
			if productTag == requiredTag {
				return true
			}
		}
	}
	return false
}

// ProductInPriceRangeSpecification checks if a product's price is within a range
// Note: This checks the first detail's price as the default price
type ProductInPriceRangeSpecification struct {
	minPrice *float64
	maxPrice *float64
}

func NewProductInPriceRangeSpecification(minPrice, maxPrice *float64) specification.Specification[*productaggregate.Product] {
	return &ProductInPriceRangeSpecification{
		minPrice: minPrice,
		maxPrice: maxPrice,
	}
}

func (s *ProductInPriceRangeSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	if product == nil || len(product.Details) == 0 {
		return false
	}

	// Get the first detail's price as default
	var price float64
	for i := range product.Details {
		if product.Details[i].Price > 0 {
			price = product.Details[i].Price
			break
		}
	}

	if s.minPrice != nil && price < *s.minPrice {
		return false
	}
	if s.maxPrice != nil && price > *s.maxPrice {
		return false
	}

	return true
}

// ProductCanBePublishedSpecification checks if a product can be published
// A product can be published if it has:
// - A name
// - A slug
// - At least one detail (with price)
// - A category
type ProductCanBePublishedSpecification struct {
}

func NewProductCanBePublishedSpecification() specification.Specification[*productaggregate.Product] {
	return &ProductCanBePublishedSpecification{}
}

func (s *ProductCanBePublishedSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	if product == nil {
		return false
	}

	// Must have name
	if product.Name == "" {
		return false
	}

	// Must have slug
	if product.Slug == "" {
		return false
	}

	// Must have category
	if product.CategoryID == 0 {
		return false
	}

	// Must have at least one detail with price
	if len(product.Details) == 0 {
		return false
	}

	hasValidPrice := false
	for i := range product.Details {
		if product.Details[i].Price > 0 {
			hasValidPrice = true
			break
		}
	}

	return hasValidPrice
}

// ProductIsVerifiedSpecification checks if a product is verified
// A product is verified if it has minimum rating and review count
type ProductIsVerifiedSpecification struct {
	minRating      float64
	minReviewCount int
}

func NewProductIsVerifiedSpecification(minRating float64, minReviewCount int) specification.Specification[*productaggregate.Product] {
	return &ProductIsVerifiedSpecification{
		minRating:      minRating,
		minReviewCount: minReviewCount,
	}
}

func (s *ProductIsVerifiedSpecification) IsSatisfiedBy(product *productaggregate.Product) bool {
	if product == nil {
		return false
	}

	ratingSpec := NewProductHasMinimumRatingSpecification(s.minRating)
	reviewCountSpec := NewProductHasMinimumReviewCountSpecification(s.minReviewCount)

	combinedSpec := specification.NewBuilder(ratingSpec).And(reviewCountSpec)
	return combinedSpec.IsSatisfiedBy(product)
}
