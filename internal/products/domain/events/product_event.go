package events

// ProductCreatedEvent is raised when a new product is created
type ProductCreatedEvent struct {
	ProductID   *uint64  `json:"product_id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Brand       string   `json:"brand"`
	CategoryIDs []uint64 `json:"category_ids"`
	Description string   `json:"description,omitempty"`
}

// ProductUpdatedEvent is raised when a product is updated
type ProductUpdatedEvent struct {
	ProductID   *uint64  `json:"product_id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Brand       string   `json:"brand"`
	CategoryIDs []uint64 `json:"category_ids"`
	Description string   `json:"description,omitempty"`
}

// ProductDeletedEvent is raised when a product is deleted
type ProductDeletedEvent struct {
	ProductID  *uint64 `json:"product_id"`
	SoftDelete bool    `json:"soft_delete"`
}
