package product_aggregate

import (
	"strconv"
	"time"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/events"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

// Product is the Aggregate Root for the Product Aggregate.
// The Product Aggregate consists of:
//   - Product (Aggregate Root)
//   - ProductFeature (Aggregate Entity)
//   - ProductDetail (Aggregate Entity)
//   - ProductSpec (Aggregate Entity)
//
// All operations on aggregate entities must go through the Product aggregate root.
type ProductID uint64

type Product struct {
	adapter.BaseEntity
	ID          ProductID `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt   `gorm:"index"`
	Name        string           `json:"name" gorm:"name"`
	Slug        string           `json:"slug" gorm:"slug;uniqueIndex"`
	Brand       string           `json:"brand" gorm:"brand"`
	Rating      float64          `json:"rating" gorm:"rating;default:0"`
	ReviewCount int              `json:"review_count" gorm:"review_count;default:0"`
	Description *string          `json:"description,omitempty" gorm:"description;type:text"`
	Features    []ProductFeature `json:"-" gorm:"foreignKey:ProductID"` // Aggregate Entity - Not in JSON, will be converted to array
	Details     []ProductDetail  `json:"-" gorm:"foreignKey:ProductID"` // Aggregate Entity - Not in JSON, will be converted to colors and variants maps
	Specs       []ProductSpec    `json:"-" gorm:"foreignKey:ProductID"` // Aggregate Entity - Not in JSON, will be converted to map
	CategoryID  uint64           `json:"category_id" gorm:"category_id"`
	Tags        []Tag            `json:"-" gorm:"many2many:product_tags;"`
	Image       string           `json:"image" gorm:"image"` // Main image (for backward compatibility)
	IsNew       bool             `json:"is_new" gorm:"is_new;default:false"`
	IsFeatured  bool             `json:"is_featured" gorm:"is_featured;default:false"`
	Sizes       []Size           `json:"-" gorm:"many2many:product_sizes;"`
	Colors      []Color          `json:"-" gorm:"many2many:product_colors;"`
	Price       float64          `json:"price" gorm:"price;default:0"`
	OriginPrice *float64         `json:"origin_price,omitempty" gorm:"origin_price"`
}

func (p *Product) TableName() string {
	return "products"
}

// NewProduct creates a new Product instance using a command
func NewProduct(cmd *commands.CreateProduct) *Product {
	product := &Product{
		Name:        cmd.Name,
		Slug:        cmd.Slug,
		Brand:       cmd.Brand,
		Description: cmd.Description,
		CategoryID:  cmd.CategoryID,
		Tags:        []Tag{},   // Tags will be set separately in command handler
		Sizes:       []Size{},  // Sizes will be set separately in command handler
		Colors:      []Color{}, // Colors will be set separately in command handler
		Image:       cmd.Image,
		IsNew:       cmd.IsNew,
		IsFeatured:  cmd.IsFeatured,
		Rating:      0,
		ReviewCount: 0,
		Price:       cmd.Price,
		OriginPrice: cmd.OriginPrice,
	}
	productID := uint64(product.ID)
	categoryID := uint64(product.CategoryID)
	product.AddEvent(&events.ProductCreatedEvent{
		ProductID:   &productID,
		Name:        product.Name,
		Slug:        product.Slug,
		Brand:       product.Brand,
		CategoryID:  categoryID,
		Description: *product.Description,
	})

	return product
}

// BeforeCreate hook to ensure JSON fields are properly initialized
// This will be called by GORM automatically
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.Features == nil {
		p.Features = []ProductFeature{}
	}
	if p.Details == nil {
		p.Details = []ProductDetail{}
	}
	if p.Specs == nil {
		p.Specs = []ProductSpec{}
	}
	if p.Tags == nil {
		p.Tags = []Tag{}
	}
	if p.Sizes == nil {
		p.Sizes = []Size{}
	}
	if p.Colors == nil {
		p.Colors = []Color{}
	}
	return nil
}

// convertTagsToStringArray converts []Tag to []string for JSON response
func convertTagsToStringArray(tags []Tag) []string {
	if len(tags) == 0 {
		return []string{}
	}
	result := make([]string, len(tags))
	for i := range tags {
		result[i] = tags[i].Name
	}
	return result
}

// convertSizesToStringArray converts []Size to []string for JSON response
func convertSizesToStringArray(sizes []Size) []string {
	if len(sizes) == 0 {
		return []string{}
	}
	result := make([]string, len(sizes))
	for i := range sizes {
		result[i] = sizes[i].Name
	}
	return result
}

// convertSizesToArray converts []Size to array of size objects for JSON response
func convertSizesToArray(sizes []Size) []map[string]interface{} {
	if len(sizes) == 0 {
		return []map[string]interface{}{}
	}
	result := make([]map[string]interface{}, len(sizes))
	for i := range sizes {
		result[i] = map[string]interface{}{
			"id":   sizes[i].ID,
			"name": sizes[i].Name,
		}
	}
	return result
}

// convertColorsToArray converts []Color to array of color objects for JSON response
func convertColorsToArray(colors []Color) []map[string]interface{} {
	if len(colors) == 0 {
		return []map[string]interface{}{}
	}
	result := make([]map[string]interface{}, len(colors))
	for i := range colors {
		result[i] = map[string]interface{}{
			"id":   colors[i].ID,
			"name": colors[i].Name,
			"hex":  colors[i].Hex,
		}
	}
	return result
}

// ToMap converts Product to map format for JSON response (new structure)
func (p *Product) ToMap() map[string]interface{} {
	// Get price, stock, discount from first detail if exists, otherwise use defaults
	defaultPrice := 0.0
	defaultStock := 0
	defaultDiscount := 0

	if len(p.Details) > 0 {
		// Use first detail with price > 0
		for i := range p.Details {
			if p.Details[i].Price > 0 {
				defaultPrice = p.Details[i].Price
				defaultStock = p.Details[i].Stock
				defaultDiscount = p.Details[i].Discount
				break
			}
		}
	}

	// Build images map: { "colorId": [urls] }
	imagesMap := make(map[string][]string)

	// Group images by color ID from Colors array
	// First, initialize images map for all colors
	for i := range p.Colors {
		colorID := strconv.FormatUint(uint64(p.Colors[i].ID), 10)
		imagesMap[colorID] = []string{}
	}

	// Then, add images from Details based on color matching
	for i := range p.Details {
		detail := &p.Details[i]
		if detail.ColorKey == nil {
			continue
		}

		// Find matching color ID from Colors array
		var colorID string
		for j := range p.Colors {
			// Match by name, slug, or colorKey
			if p.Colors[j].Name == *detail.ColorKey ||
				p.Colors[j].Slug == *detail.ColorKey ||
				strconv.FormatUint(uint64(p.Colors[j].ID), 10) == *detail.ColorKey {
				colorID = strconv.FormatUint(uint64(p.Colors[j].ID), 10)
				break
			}
		}

		// If color not found, skip this detail
		if colorID == "" {
			continue
		}

		// Convert images from attachments
		images := make([]string, 0)
		if detail.Images != nil {
			for j := range detail.Images {
				img := &detail.Images[j]
				images = append(images, img.FilePath)
			}
		}

		// Add images to map (append to existing)
		if len(images) > 0 {
			imagesMap[colorID] = append(imagesMap[colorID], images...)
		}
	}

	// If no images from details, use thumbnail for first color
	if len(imagesMap) == 0 || (len(imagesMap) > 0 && p.Image != "") {
		// Use first color ID or "1" as default
		if len(p.Colors) > 0 {
			firstColorID := strconv.FormatUint(uint64(p.Colors[0].ID), 10)
			if len(imagesMap[firstColorID]) == 0 && p.Image != "" {
				imagesMap[firstColorID] = []string{p.Image}
			}
		} else if p.Image != "" {
			imagesMap["1"] = []string{p.Image}
		}
	}

	// Remove empty image arrays
	for colorID, images := range imagesMap {
		if len(images) == 0 {
			delete(imagesMap, colorID)
		}
	}

	// Build categories array: [{id, name}]
	categoriesArray := []map[string]interface{}{
		{
			"id":   p.CategoryID,
			"name": "همه", // Default name, should be loaded from Category entity if needed
		},
	}

	// Build colors array: [{id, name, hex}]
	colorsArray := convertColorsToArray(p.Colors)

	// Build specs array: [{key, value}]
	specsArray := make([]map[string]interface{}, 0, len(p.Specs))
	for i := range p.Specs {
		spec := &p.Specs[i]
		specsArray = append(specsArray, map[string]interface{}{
			"key":   spec.Key,
			"value": spec.Value,
		})
	}

	// Build sizes array: [{id, name}]
	sizesArray := convertSizesToArray(p.Sizes)

	// Build variant map: { "colorId": { "sizeId": { "stock": number } } }
	variantMap := make(map[string]map[string]map[string]interface{})
	for i := range p.Details {
		detail := &p.Details[i]
		if detail.ColorKey == nil || detail.SizeKey == nil {
			continue
		}

		// Find matching color ID from Colors array
		var colorID string
		for j := range p.Colors {
			if p.Colors[j].Name == *detail.ColorKey ||
				p.Colors[j].Slug == *detail.ColorKey ||
				strconv.FormatUint(uint64(p.Colors[j].ID), 10) == *detail.ColorKey {
				colorID = strconv.FormatUint(uint64(p.Colors[j].ID), 10)
				break
			}
		}

		// Find matching size ID from Sizes array
		var sizeID string
		for j := range p.Sizes {
			if p.Sizes[j].Name == *detail.SizeKey ||
				p.Sizes[j].Slug == *detail.SizeKey ||
				strconv.FormatUint(uint64(p.Sizes[j].ID), 10) == *detail.SizeKey {
				sizeID = strconv.FormatUint(uint64(p.Sizes[j].ID), 10)
				break
			}
		}

		// If color or size not found, skip this detail
		if colorID == "" || sizeID == "" {
			continue
		}

		// Initialize nested maps if needed
		if variantMap[colorID] == nil {
			variantMap[colorID] = make(map[string]map[string]interface{})
		}

		// Add stock for this color-size combination
		variantMap[colorID][sizeID] = map[string]interface{}{
			"stock": detail.Stock,
		}
	}

	// Use product-level price if available, otherwise use defaultPrice from details
	productPrice := defaultPrice
	if p.Price > 0 {
		productPrice = p.Price
	}

	// Build result with new structure
	result := map[string]interface{}{
		"id":          uint64(p.ID),
		"seller_id":   1, // TODO: Add seller_id to Product entity
		"brand":       p.Brand,
		"title":       p.Name, // Map Name to title
		"slug":        p.Slug,
		"description": p.Description,
		"thumbnail":   p.Image, // Map Image to thumbnail
		"categories":  categoriesArray,
		"discount":    defaultDiscount,
		"stock":       defaultStock,
		"price":       productPrice,
		"rating":      p.Rating,
		"is_featured": p.IsFeatured,
		"is_new":      p.IsNew,
		"created_at":  p.CreatedAt,
		"colors":      colorsArray,
		"sizes":       sizesArray,
		"variant":     variantMap,
		"tags":        convertTagsToStringArray(p.Tags),
		"features":    convertFeaturesToArray(p.Features),
		"specs":       specsArray,
		"images":      imagesMap,
	}

	// Add origin_price if available
	if p.OriginPrice != nil {
		result["origin_price"] = *p.OriginPrice
	}

	return result
}

// convertFeaturesToArray converts []ProductFeature to []string for JSON response
func convertFeaturesToArray(features []ProductFeature) []string {
	if len(features) == 0 {
		return []string{}
	}
	result := make([]string, len(features))
	for i := range features {
		result[i] = features[i].Feature
	}
	return result
}
