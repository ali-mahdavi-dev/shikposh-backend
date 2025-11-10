package entity

import (
	"strconv"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity/product_aggregate"
	"shikposh-backend/internal/products/domain/events"
	"shikposh-backend/pkg/framework/adapter"

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
type Product struct {
	adapter.BaseEntity
	Name        string                             `json:"name" gorm:"name"`
	Slug        string                             `json:"slug" gorm:"slug;uniqueIndex"`
	Brand       string                             `json:"brand" gorm:"brand"`
	Rating      float64                            `json:"rating" gorm:"rating;default:0"`
	ReviewCount int                                `json:"review_count" gorm:"review_count;default:0"`
	Description *string                            `json:"description,omitempty" gorm:"description;type:text"`
	Features    []product_aggregate.ProductFeature `json:"-" gorm:"foreignKey:ProductID"` // Aggregate Entity - Not in JSON, will be converted to array
	Details     []product_aggregate.ProductDetail  `json:"-" gorm:"foreignKey:ProductID"` // Aggregate Entity - Not in JSON, will be converted to colors and variants maps
	Specs       []product_aggregate.ProductSpec    `json:"-" gorm:"foreignKey:ProductID"` // Aggregate Entity - Not in JSON, will be converted to map
	CategoryID  uint64                             `json:"category_id" gorm:"category_id"`
	Category    *Category                          `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Tags        []string                           `json:"tags,omitempty" gorm:"type:jsonb"`
	Image       string                             `json:"image" gorm:"image"` // Main image (for backward compatibility)
	IsNew       bool                               `json:"is_new" gorm:"is_new;default:false"`
	IsFeatured  bool                               `json:"is_featured" gorm:"is_featured;default:false"`
	Sizes       []string                           `json:"sizes" gorm:"type:jsonb"`
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
		Tags:        cmd.Tags,
		Sizes:       cmd.Sizes,
		Image:       cmd.Image,
		IsNew:       cmd.IsNew,
		IsFeatured:  cmd.IsFeatured,
		Rating:      0,
		ReviewCount: 0,
	}
	product.AddEvent(&events.ProductCreatedEvent{
		ProductID:   &product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		Brand:       product.Brand,
		CategoryID:  product.CategoryID,
		Description: *product.Description,
	})

	return product
}

// BeforeCreate hook to ensure JSON fields are properly initialized
// This will be called by GORM automatically
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.Features == nil {
		p.Features = []product_aggregate.ProductFeature{}
	}
	if p.Details == nil {
		p.Details = []product_aggregate.ProductDetail{}
	}
	if p.Specs == nil {
		p.Specs = []product_aggregate.ProductSpec{}
	}
	if p.Tags == nil {
		p.Tags = []string{}
	}
	if p.Sizes == nil {
		p.Sizes = []string{}
	}
	return nil
}

// ToMap converts Colors and Variants to map format for JSON response
func (p *Product) ToMap() map[string]interface{} {
	// Get default price from first detail if exists, otherwise use 0
	defaultPrice := 0.0
	defaultDiscount := 0
	var defaultOriginalPrice *float64

	if len(p.Details) > 0 {
		for i := range p.Details {
			if p.Details[i].Price > 0 {
				defaultPrice = p.Details[i].Price
				defaultDiscount = p.Details[i].Discount
				if p.Details[i].OriginalPrice != nil {
					defaultOriginalPrice = p.Details[i].OriginalPrice
				}
				break
			}
		}
	}

	result := map[string]interface{}{
		"id":           strconv.FormatUint(p.ID, 10),
		"name":         p.Name,
		"slug":         p.Slug,
		"brand":        p.Brand,
		"rating":       p.Rating,
		"review_count": p.ReviewCount,
		"description":  p.Description,
		"category_id":  p.CategoryID,
		"tags":         p.Tags,
		"image":        p.Image,
		"price":        defaultPrice, // Default price from first detail
		"discount":     defaultDiscount,
		"is_new":       p.IsNew,
		"is_featured":  p.IsFeatured,
		"sizes":        p.Sizes,
	}

	if defaultOriginalPrice != nil {
		result["original_price"] = *defaultOriginalPrice
	}

	if p.Category != nil {
		result["category"] = p.Category.Slug
	} else {
		result["category"] = ""
	}

	// Convert Details to colors and variants maps
	colorsMap := make(map[string]map[string]interface{})
	variantsMap := make(map[string]map[string]map[string]interface{})

	for i := range p.Details {
		detail := &p.Details[i]

		// Convert images from attachments
		images := make([]string, 0)
		if detail.Images != nil {
			for j := range detail.Images {
				img := &detail.Images[j]
				images = append(images, img.FilePath)
			}
		}

		// If has color_key but no size_key, it's a color definition
		if detail.ColorKey != nil && detail.SizeKey == nil {
			colorKey := *detail.ColorKey
			colorMap := map[string]interface{}{}
			if detail.ColorName != nil {
				colorMap["name"] = *detail.ColorName
			}
			if detail.Stock > 0 {
				colorMap["stock"] = detail.Stock
			}
			if detail.Discount > 0 {
				colorMap["discount"] = detail.Discount
			}
			colorsMap[colorKey] = colorMap
		}

		// If has both color_key and size_key, it's a variant
		if detail.ColorKey != nil && detail.SizeKey != nil {
			colorKey := *detail.ColorKey
			sizeKey := *detail.SizeKey

			if variantsMap[colorKey] == nil {
				variantsMap[colorKey] = make(map[string]map[string]interface{})
			}

			variantData := map[string]interface{}{
				"price":    detail.Price,
				"stock":    detail.Stock,
				"discount": detail.Discount,
				"images":   images,
			}

			if detail.OriginalPrice != nil {
				variantData["original_price"] = *detail.OriginalPrice
			}

			variantsMap[colorKey][sizeKey] = variantData
		}
	}

	result["colors"] = colorsMap
	result["variants"] = variantsMap

	// Convert Features to array (ordered by order field)
	featuresArray := make([]string, 0, len(p.Features))
	for i := range p.Features {
		feature := &p.Features[i]
		featuresArray = append(featuresArray, feature.Feature)
	}
	result["features"] = featuresArray

	// Convert Specs to map (ordered by order field)
	specsMap := make(map[string]string)
	for i := range p.Specs {
		spec := &p.Specs[i]
		specsMap[spec.Key] = spec.Value
	}
	result["specs"] = specsMap

	return result
}
