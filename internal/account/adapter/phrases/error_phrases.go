package phrases

import (
	"github.com/ali-mahdavi-dev/shikposh-framework/errors/phrases"
)

// Account module error phrases
const (
	UserNotFound       phrases.MessagePhrase = "Account.User.NotFound"
	UserAlreadyExists  phrases.MessagePhrase = "Account.User.AlreadyExists"
	UserAgeInvalid     phrases.MessagePhrase = "Account.User.AgeInvalid"
	UserInvalid        phrases.MessagePhrase = "Account.User.Invalid"
	PhoneAlreadyExists phrases.MessagePhrase = "Account.User.PhoneAlreadyExists"
	OtpInvalid         phrases.MessagePhrase = "Account.Otp.Invalid"
	OtpExpired         phrases.MessagePhrase = "Account.Otp.Expired"
	OtpRateLimited     phrases.MessagePhrase = "Account.Otp.RateLimited"
)

// RegisterAccountPhrases registers error phrases for the account module
func RegisterAccountPhrases() {
	phrases.GetRegistry().Register(map[phrases.Language]map[phrases.MessagePhrase]string{
		phrases.Fa: {
			UserNotFound:       "کاربر پیدا نشد",
			UserAlreadyExists:  "کاربر از قبل وجود دارد",
			UserAgeInvalid:     "سن کاربر کمتر از ۱۸ است",
			UserInvalid:        "اطلاعات کاربر درست نمیباشد",
			PhoneAlreadyExists: "این شماره تلفن قبلاً ثبت شده است",
			OtpInvalid:         "کد OTP نامعتبر است",
			OtpExpired:         "کد OTP منقضی شده است",
			OtpRateLimited:     "شما درخواست‌های زیادی ارسال کرده‌اید. لطفاً کمی صبر کنید",
		},
		phrases.En: {
			UserNotFound:       "User not found",
			UserAlreadyExists:  "User already exists",
			UserAgeInvalid:     "User age is less than 18",
			UserInvalid:        "User information is not valid",
			PhoneAlreadyExists: "This phone number is already registered",
			OtpInvalid:         "Invalid OTP code",
			OtpExpired:         "OTP code has expired",
			OtpRateLimited:     "Too many requests. Please wait a moment",
		},
	})
}
