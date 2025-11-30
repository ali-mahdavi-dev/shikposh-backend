package phrases

import (
	"github.com/ali-mahdavi-dev/shikposh-framework/errors/phrases"
)

// Admin module error phrases
const (
	AdminAccessDenied     phrases.MessagePhrase = "Admin.Access.Denied"
	SuperuserAccessDenied phrases.MessagePhrase = "Admin.Superuser.Access.Denied"
)

// RegisterAdminPhrases registers error phrases for the admin module
func RegisterAdminPhrases() {
	phrases.GetRegistry().Register(map[phrases.Language]map[phrases.MessagePhrase]string{
		phrases.Fa: {
			AdminAccessDenied:     "دسترسی ادمین مورد نیاز است",
			SuperuserAccessDenied: "دسترسی سوپر یوزر مورد نیاز است",
		},
		phrases.En: {
			AdminAccessDenied:     "Admin access required",
			SuperuserAccessDenied: "Superuser access required",
		},
	})
}
