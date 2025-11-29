package commands

// CreateOrder command for creating a new order
type CreateOrder struct {
	UserID          uint64                `json:"user_id"`
	Items           []OrderItemInput      `json:"items" validate:"required,min=1"`
	ShippingAddress *ShippingAddressInput `json:"shipping_address,omitempty"`
	PaymentMethod   string                `json:"payment_method" validate:"required"`
	TotalAmount     int64                 `json:"total_amount" validate:"required,min=0"`
	DiscountAmount  int64                 `json:"discount_amount" validate:"min=0"`
	ShippingCost    int64                 `json:"shipping_cost" validate:"min=0"`
	FinalAmount     int64                 `json:"final_amount" validate:"required,min=0"`
}

// OrderItemInput represents an item in the order
type OrderItemInput struct {
	ProductID    uint64  `json:"product_id" validate:"required"`
	ProductName  string  `json:"product_name" validate:"required"`
	ProductSlug  *string `json:"product_slug,omitempty"`
	ProductImage *string `json:"product_image,omitempty"`
	Quantity     int     `json:"quantity" validate:"required,min=1"`
	Price        int64   `json:"price" validate:"required,min=0"`
	Discount     int64   `json:"discount" validate:"min=0"`
	Color        *string `json:"color,omitempty"`
	Size         *string `json:"size,omitempty"`
	VariantID    *uint64 `json:"variant_id,omitempty"`
}

// ShippingAddressInput represents shipping address for the order
type ShippingAddressInput struct {
	FullName   string  `json:"full_name" validate:"required"`
	Phone      string  `json:"phone" validate:"required"`
	Address    string  `json:"address" validate:"required"`
	City       string  `json:"city" validate:"required"`
	Province   string  `json:"province" validate:"required"`
	PostalCode *string `json:"postal_code,omitempty"`
}
