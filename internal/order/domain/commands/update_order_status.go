package commands

// UpdateOrderStatus command for updating an order's status / payment info / notes
type UpdateOrderStatus struct {
	OrderID       uint64  `json:"order_id" validate:"required"`
	Status        *string `json:"status,omitempty"`
	PaymentStatus *string `json:"payment_status,omitempty"`
	Notes         *string `json:"notes,omitempty"`
}


