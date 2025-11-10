package events

// ProductCreatedEvent is raised when a new product is created
type ProductCreatedEvent struct {
	ProductID   *uint64 `json:"product_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Brand       string `json:"brand"`
	CategoryID  uint64 `json:"category_id"`
	Description string `json:"description,omitempty"`
}
