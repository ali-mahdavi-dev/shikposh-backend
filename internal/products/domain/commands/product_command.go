package commands

// CreateProduct command for creating a new product
type CreateProduct struct {
	SellerID         *uint64               `json:"seller_id,omitempty"`
	Title            string                `json:"title" validate:"required,min=3"`
	Slug             string                `json:"slug" validate:"required,min=3"`
	Brand            string                `json:"brand" validate:"required,min=2"`
	Description      *string               `json:"description,omitempty" validate:"omitempty,min=10"`
	ShortDescription *string               `json:"short_description,omitempty" validate:"omitempty,max=500"`
	Thumbnail        string                `json:"thumbnail"`
	Categories       []uint64              `json:"categories" validate:"required,min=1"` // At least one category is required
	Discount         int                   `json:"discount" validate:"min=0,max=100"`
	Stock            int                   `json:"stock" validate:"min=0"`
	OriginPrice      *int64                `json:"origin_price,omitempty" validate:"omitempty,min=0"`
	Price            int64                 `json:"price" validate:"min=0"`
	IsNew            bool                  `json:"is_new"`
	IsFeatured       bool                  `json:"is_featured"`
	Colors           []uint64              `json:"colors,omitempty"`   // Color IDs
	Sizes            []uint64              `json:"sizes,omitempty"`    // Size IDs
	Variants         []ProductVariantInput `json:"variants,omitempty"` // color_id + size_id + stock
	Tags             []string              `json:"tags,omitempty"`
	Features         []string              `json:"features,omitempty"`
	Specs            []ProductSpecInput    `json:"specs,omitempty"`
	Images           []ProductImageInput   `json:"images,omitempty"` // color_id + urls
}

// ProductVariantInput - ترکیب رنگ + سایز با موجودی
type ProductVariantInput struct {
	ColorID uint64 `json:"color_id" validate:"required"`
	SizeID  uint64 `json:"size_id" validate:"required"`
	Stock   int    `json:"stock" validate:"min=0"`
}

// ProductSpecInput - مشخصات محصول
type ProductSpecInput struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Order int    `json:"order"`
}

// ProductImageInput - عکس‌ها بر اساس رنگ
type ProductImageInput struct {
	ColorID uint64   `json:"color_id" validate:"required"`
	URLs    []string `json:"urls" validate:"required,min=1"`
}

// UpdateProduct command for updating an existing product
type UpdateProduct struct {
	ID               uint64                `json:"id" validate:"required"`
	SellerID         *uint64               `json:"seller_id,omitempty"`
	Title            string                `json:"title" validate:"required,min=3"`
	Slug             string                `json:"slug" validate:"required,min=3"`
	Brand            string                `json:"brand" validate:"required,min=2"`
	Description      *string               `json:"description,omitempty" validate:"omitempty,min=10"`
	ShortDescription *string               `json:"short_description,omitempty" validate:"omitempty,max=500"`
	Thumbnail        *string               `json:"thumbnail,omitempty"`
	Categories       []uint64              `json:"categories,omitempty"`
	Discount         *int                  `json:"discount,omitempty" validate:"omitempty,min=0,max=100"`
	Stock            *int                  `json:"stock,omitempty" validate:"omitempty,min=0"`
	OriginPrice      *int64                `json:"origin_price,omitempty" validate:"omitempty,min=0"`
	Price            *int64                `json:"price,omitempty" validate:"omitempty,min=0"`
	IsNew            *bool                 `json:"is_new,omitempty"`
	IsFeatured       *bool                 `json:"is_featured,omitempty"`
	Colors           []uint64              `json:"colors,omitempty"`
	Sizes            []uint64              `json:"sizes,omitempty"`
	Variants         []ProductVariantInput `json:"variants,omitempty"`
	Tags             []string              `json:"tags,omitempty"`
	Features         []string              `json:"features,omitempty"`
	Specs            []ProductSpecInput    `json:"specs,omitempty"`
	Images           []ProductImageInput   `json:"images,omitempty"`
}

// DeleteProduct command for deleting a product
type DeleteProduct struct {
	ID         uint64 `json:"id" validate:"required"`
	SoftDelete bool   `json:"soft_delete"`
}
