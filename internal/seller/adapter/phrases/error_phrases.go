package phrases

import (
	"github.com/ali-mahdavi-dev/shikposh-framework/errors/phrases"
)

// Seller module error phrases
const (
	SellerNotFound phrases.MessagePhrase = "Seller.NotFound"
)

// RegisterSellerPhrases registers error phrases for the seller module
func RegisterSellerPhrases() {
	phrases.GetRegistry().Register(map[phrases.Language]map[phrases.MessagePhrase]string{
		phrases.Fa: {
			SellerNotFound: "فروشنده پیدا نشد",
		},
		phrases.En: {
			SellerNotFound: "Seller not found",
		},
	})
}
