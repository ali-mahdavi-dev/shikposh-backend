package commands

// AddToWishlist command for adding a product to wishlist
type AddToWishlist struct {
	UserID    uint64 `json:"user_id" validate:"required"`
	ProductID uint64 `json:"product_id" validate:"required"`
}

// RemoveFromWishlist command for removing a product from wishlist
type RemoveFromWishlist struct {
	UserID    uint64 `json:"user_id" validate:"required"`
	ProductID uint64 `json:"product_id" validate:"required"`
}

// ToggleWishlist command for toggling a product in wishlist
type ToggleWishlist struct {
	UserID    uint64 `json:"user_id" validate:"required"`
	ProductID uint64 `json:"product_id" validate:"required"`
}

// SyncWishlist command for syncing wishlist with server
type SyncWishlist struct {
	UserID     uint64   `json:"user_id" validate:"required"`
	ProductIDs []uint64 `json:"product_ids" validate:"required"`
}
