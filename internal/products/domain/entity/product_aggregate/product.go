package product_aggregate

import (
	"strconv"
	"time"

	"shikposh-backend/internal/products/domain/commands"
	"shikposh-backend/internal/products/domain/entity"
	"shikposh-backend/internal/products/domain/events"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type ProductID uint64

type Product struct {
	adapter.BaseEntity
	ID               ProductID `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	SellerID         *uint64        `json:"seller_id,omitempty" gorm:"seller_id"`
	Title            string         `json:"title" gorm:"column:title"`
	Slug             string         `json:"slug" gorm:"slug;uniqueIndex"`
	Brand            string         `json:"brand" gorm:"brand"`
	Description      *string        `json:"description,omitempty" gorm:"description;type:text"`
	ShortDescription *string        `json:"short_description,omitempty" gorm:"short_description"`
	Thumbnail        string         `json:"thumbnail" gorm:"column:thumbnail"`
	Discount         int            `json:"discount" gorm:"discount;default:0"`
	Stock            int            `json:"stock" gorm:"stock;default:0"`
	OriginPrice      *int64         `json:"origin_price,omitempty" gorm:"origin_price"`
	Price            int64          `json:"price" gorm:"price;default:0"`
	Rating           float64        `json:"rating" gorm:"rating;default:0"`
	IsFeatured       bool           `json:"is_featured" gorm:"is_featured;default:false"`
	IsNew            bool           `json:"is_new" gorm:"is_new;default:false"`
	ReviewCount      int            `json:"review_count" gorm:"review_count;default:0"`
	// Relations
	Categories []entity.Category `json:"-" gorm:"many2many:product_categories;"`
	Colors     []Color           `json:"-" gorm:"many2many:product_colors;"`
	Sizes      []Size            `json:"-" gorm:"many2many:product_sizes;"`
	Variants   []ProductVariant  `json:"-" gorm:"foreignKey:ProductID"`
	Tags       []Tag             `json:"-" gorm:"many2many:product_tags;"`
	Features   []ProductFeature  `json:"-" gorm:"foreignKey:ProductID"`
	Specs      []ProductSpec     `json:"-" gorm:"foreignKey:ProductID"`
	Images     []ProductImage    `json:"-" gorm:"foreignKey:ProductID"`
}

func (p *Product) TableName() string {
	return "products"
}

// NewProduct creates a new Product instance using a command
func NewProduct(cmd *commands.CreateProduct) *Product {
	product := &Product{
		SellerID:         cmd.SellerID,
		Title:            cmd.Title,
		Slug:             cmd.Slug,
		Brand:            cmd.Brand,
		Description:      cmd.Description,
		ShortDescription: cmd.ShortDescription,
		Thumbnail:        cmd.Thumbnail,
		Discount:         cmd.Discount,
		Stock:            cmd.Stock,
		IsNew:            cmd.IsNew,
		IsFeatured:       cmd.IsFeatured,
		Rating:           0,
		ReviewCount:      0,
		Price:            cmd.Price,
		OriginPrice:      cmd.OriginPrice,
	}
	productID := uint64(product.ID)
	var categoryID uint64
	if len(product.Categories) > 0 {
		categoryID = uint64(product.Categories[0].ID)
	}
	product.AddEvent(&events.ProductCreatedEvent{
		ProductID:  &productID,
		Name:       product.Title,
		Slug:       product.Slug,
		Brand:      product.Brand,
		CategoryID: categoryID,
		Description: func() string {
			if product.Description != nil {
				return *product.Description
			}
			return ""
		}(),
	})

	return product
}

// BeforeCreate hook
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	return nil
}

// ToMap converts Product to Elasticsearch document format
func (p *Product) ToMap() map[string]interface{} {
	// categories: [{ id, name, slug }]
	categoriesArray := make([]map[string]interface{}, 0, len(p.Categories))
	for i := range p.Categories {
		cat := &p.Categories[i]
		categoriesArray = append(categoriesArray, map[string]interface{}{
			"id":   uint64(cat.ID),
			"name": cat.Name,
			"slug": cat.Slug,
		})
	}
	// اگر هیچ category ای وجود نداشت، یک category پیش‌فرض اضافه نمی‌کنیم
	// categoriesArray می‌تواند خالی باشد

	// colors: [{ id, name, hex }]
	colorsArray := make([]map[string]interface{}, 0, len(p.Colors))
	for i := range p.Colors {
		c := &p.Colors[i]
		colorsArray = append(colorsArray, map[string]interface{}{
			"id":   uint64(c.ID),
			"name": c.Name,
			"hex":  c.Hex,
		})
	}

	// sizes: [{ id, name }]
	sizesArray := make([]map[string]interface{}, 0, len(p.Sizes))
	for i := range p.Sizes {
		s := &p.Sizes[i]
		sizesArray = append(sizesArray, map[string]interface{}{
			"id":   uint64(s.ID),
			"name": s.Name,
		})
	}

	// variant: { "colorId": { "sizeId": { "stock": N } } }
	variantMap := make(map[string]map[string]map[string]interface{})
	for i := range p.Variants {
		v := &p.Variants[i]
		colorID := strconv.FormatUint(uint64(v.ColorID), 10)
		sizeID := strconv.FormatUint(uint64(v.SizeID), 10)
		if variantMap[colorID] == nil {
			variantMap[colorID] = make(map[string]map[string]interface{})
		}
		variantMap[colorID][sizeID] = map[string]interface{}{"stock": v.Stock}
	}

	// tags: ["tag1", "tag2"]
	tagsArray := make([]string, 0, len(p.Tags))
	for i := range p.Tags {
		tagsArray = append(tagsArray, p.Tags[i].Name)
	}

	// features: ["feature1", "feature2"]
	featuresArray := make([]string, 0, len(p.Features))
	for i := range p.Features {
		featuresArray = append(featuresArray, p.Features[i].Feature)
	}

	// specs: [{ key, value }]
	specsArray := make([]map[string]interface{}, 0, len(p.Specs))
	for i := range p.Specs {
		s := &p.Specs[i]
		specsArray = append(specsArray, map[string]interface{}{
			"key":   s.Key,
			"value": s.Value,
		})
	}

	// images: { "colorId": ["url1", "url2"] }
	imagesMap := make(map[string][]string)
	for i := range p.Images {
		img := &p.Images[i]
		colorID := strconv.FormatUint(uint64(img.ColorID), 10)
		imagesMap[colorID] = append(imagesMap[colorID], img.URL)
	}
	if len(imagesMap) == 0 && p.Thumbnail != "" {
		if len(p.Colors) > 0 {
			colorID := strconv.FormatUint(uint64(p.Colors[0].ID), 10)
			imagesMap[colorID] = []string{p.Thumbnail}
		} else {
			imagesMap["1"] = []string{p.Thumbnail}
		}
	}

	// seller_id
	sellerID := uint64(1)
	if p.SellerID != nil {
		sellerID = *p.SellerID
	}

	result := map[string]interface{}{
		"id":          uint64(p.ID),
		"seller_id":   sellerID,
		"brand":       p.Brand,
		"title":       p.Title,
		"slug":        p.Slug,
		"description": p.Description,
		"thumbnail":   p.Thumbnail,
		"categories":  categoriesArray,
		"discount":    p.Discount,
		"stock":       p.Stock,
		"price":       p.Price,
		"rating":      p.Rating,
		"is_featured": p.IsFeatured,
		"is_new":      p.IsNew,
		"created_at":  p.CreatedAt,
		"colors":      colorsArray,
		"sizes":       sizesArray,
		"variant":     variantMap,
		"tags":        tagsArray,
		"features":    featuresArray,
		"specs":       specsArray,
		"images":      imagesMap,
	}

	if p.ShortDescription != nil {
		result["short_description"] = *p.ShortDescription
	}

	if p.OriginPrice != nil {
		result["orgin_price"] = *p.OriginPrice
	}

	return result
}
