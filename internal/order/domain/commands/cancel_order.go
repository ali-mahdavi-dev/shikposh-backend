package commands

// CancelOrder command for cancelling an order
type CancelOrder struct {
	OrderID uint64 `json:"order_id" validate:"required"`
}


