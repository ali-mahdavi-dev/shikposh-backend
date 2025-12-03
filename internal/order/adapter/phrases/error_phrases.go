package phrases

import (
	"github.com/ali-mahdavi-dev/shikposh-framework/errors/phrases"
)

// Orders module error phrases
const (
	OrderNotFound          phrases.MessagePhrase = "Orders.Order.NotFound"
	OrderCannotBeCancelled phrases.MessagePhrase = "Orders.Order.CannotBeCancelled"
	OrderAccessDenied      phrases.MessagePhrase = "Orders.Order.AccessDenied"
)

// RegisterOrdersPhrases registers error phrases for the orders module
func RegisterOrdersPhrases() {
	phrases.GetRegistry().Register(map[phrases.Language]map[phrases.MessagePhrase]string{
		phrases.Fa: {
			OrderNotFound:          "سفارش یافت نشد",
			OrderCannotBeCancelled: "سفارش در وضعیت فعلی قابل لغو نیست",
			OrderAccessDenied:      "شما دسترسی به این سفارش ندارید",
		},
		phrases.En: {
			OrderNotFound:          "Order not found",
			OrderCannotBeCancelled: "Order cannot be cancelled in current status",
			OrderAccessDenied:      "Access denied to this order",
		},
	})
}
