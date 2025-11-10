package commands

type CreateProduct struct {
	Name        string                `json:"name" validate:"required,min=3"`
	Slug        string                `json:"slug" validate:"required,min=3"`
	Brand       string                `json:"brand" validate:"required,min=2"`
	Description *string               `json:"description,omitempty" validate:"omitempty,min=10"`
	CategoryID  uint64                `json:"category_id" validate:"required"`
	Tags        []string              `json:"tags,omitempty"`
	Sizes       []string              `json:"sizes"`
	Image       string                `json:"image"`
	IsNew       bool                  `json:"is_new"`
	IsFeatured  bool                  `json:"is_featured"`
	Features    []ProductFeatureInput `json:"features"`
	Details     []ProductDetailInput  `json:"details"`
	Specs       []ProductSpecInput    `json:"specs"`
}

type ProductFeatureInput struct {
	Feature string `json:"feature" validate:"required"`
	Order   int    `json:"order"`
}

type ProductDetailInput struct {
	ColorKey      *string  `json:"color_key,omitempty"`
	ColorName     *string  `json:"color_name,omitempty"`
	SizeKey       *string  `json:"size_key,omitempty"`
	Price         float64  `json:"price" validate:"required,min=0"`
	OriginalPrice *float64 `json:"original_price,omitempty"`
	Stock         int      `json:"stock" validate:"min=0"`
	Discount      int      `json:"discount" validate:"min=0,max=100"`
	Images        []string `json:"images"`
}

type ProductSpecInput struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Order int    `json:"order"`
}

type UpdateProduct struct {
	ID          uint64                `json:"id" validate:"required"`
	Name        string                `json:"name" validate:"required,min=3"`
	Slug        string                `json:"slug" validate:"required,min=3"`
	Brand       string                `json:"brand" validate:"required,min=2"`
	Description *string               `json:"description,omitempty" validate:"omitempty,min=10"`
	CategoryID  uint64                `json:"category_id" validate:"required"`
	Tags        []string              `json:"tags,omitempty"`
	Sizes       []string              `json:"sizes,omitempty"`
	Image       *string               `json:"image,omitempty"`
	IsNew       *bool                 `json:"is_new,omitempty"`
	IsFeatured  *bool                 `json:"is_featured,omitempty"`
	Features    []ProductFeatureInput `json:"features,omitempty"`
	Details     []ProductDetailInput  `json:"details,omitempty"`
	Specs       []ProductSpecInput    `json:"specs,omitempty"`
}

type DeleteProduct struct {
	ID         uint64 `json:"id" validate:"required"`
	SoftDelete bool   `json:"soft_delete"` // If true, soft delete; if false, hard delete
}
