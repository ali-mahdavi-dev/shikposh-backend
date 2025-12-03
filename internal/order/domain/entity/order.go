package entity

import (
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type OrderID uint64

type OrderStatus string

const (
	OrderStatusPending          OrderStatus = "pending"
	OrderStatusPaymentConfirmed OrderStatus = "payment_confirmed"
	OrderStatusProcessing       OrderStatus = "processing"
	OrderStatusConfirmed        OrderStatus = "confirmed"
	OrderStatusShipped          OrderStatus = "shipped"
	OrderStatusDelivered        OrderStatus = "delivered"
	OrderStatusCancelled        OrderStatus = "cancelled"
	OrderStatusRefunded         OrderStatus = "refunded"
)

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type Order struct {
	adapter.BaseEntity
	ID             OrderID        `gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	OrderNumber    string         `json:"order_number" gorm:"column:order_number;uniqueIndex;not null"`
	UserID         uint64         `json:"user_id" gorm:"column:user_id;not null;index"`
	Status         OrderStatus    `json:"status" gorm:"column:status;type:varchar(20);default:'payment_confirmed'"`
	TotalAmount    int64          `json:"total_amount" gorm:"column:total_amount;not null"`
	DiscountAmount int64          `json:"discount_amount" gorm:"column:discount_amount;default:0"`
	ShippingCost   int64          `json:"shipping_cost" gorm:"column:shipping_cost;default:0"`
	FinalAmount    int64          `json:"final_amount" gorm:"column:final_amount;not null"`
	PaymentMethod  *string        `json:"payment_method,omitempty" gorm:"column:payment_method"`
	PaymentStatus  PaymentStatus  `json:"payment_status" gorm:"column:payment_status;type:varchar(20);default:'pending'"`
	TrackingNumber *string        `json:"tracking_number,omitempty" gorm:"column:tracking_number"`
	ShippedAt      *time.Time     `json:"shipped_at,omitempty" gorm:"column:shipped_at"`
	DeliveredAt    *time.Time     `json:"delivered_at,omitempty" gorm:"column:delivered_at"`
	Notes          *string        `json:"notes,omitempty" gorm:"column:notes;type:text"`
	// Relations
	Items           []OrderItem   `json:"items" gorm:"foreignKey:OrderID"`
	ShippingAddress *OrderAddress `json:"shipping_address,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID           uint64    `gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	OrderID      OrderID   `json:"order_id" gorm:"column:order_id;not null;index"`
	ProductID    uint64    `json:"product_id" gorm:"column:product_id;not null;index"`
	ProductName  string    `json:"product_name" gorm:"column:product_name;not null"`
	ProductSlug  *string   `json:"product_slug,omitempty" gorm:"column:product_slug"`
	ProductImage *string   `json:"product_image,omitempty" gorm:"column:product_image"`
	Quantity     int       `json:"quantity" gorm:"column:quantity;not null"`
	Price        int64     `json:"price" gorm:"column:price;not null"`
	Discount     int64     `json:"discount" gorm:"column:discount;default:0"`
	Color        *string   `json:"color,omitempty" gorm:"column:color"`
	Size         *string   `json:"size,omitempty" gorm:"column:size"`
	VariantID    *uint64   `json:"variant_id,omitempty" gorm:"column:variant_id"`
}

type OrderAddress struct {
	ID         uint64    `gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	OrderID    OrderID   `json:"order_id" gorm:"column:order_id;not null;index"`
	FullName   string    `json:"full_name" gorm:"column:full_name;not null"`
	Phone      string    `json:"phone" gorm:"column:phone;not null"`
	Address    string    `json:"address" gorm:"column:address;not null;type:text"`
	City       string    `json:"city" gorm:"column:city;not null"`
	Province   string    `json:"province" gorm:"column:province;not null"`
	PostalCode *string   `json:"postal_code,omitempty" gorm:"column:postal_code"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *OrderItem) TableName() string {
	return "order_items"
}

func (o *OrderAddress) TableName() string {
	return "order_addresses"
}

// GenerateOrderNumber generates a unique order number
func GenerateOrderNumber() string {
	return "ORD-" + time.Now().Format("20060102") + "-" + time.Now().Format("150405") + "-" + time.Now().Format("000")
}
