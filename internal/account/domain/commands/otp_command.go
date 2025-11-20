package commands

// SendOtp command for sending OTP
type SendOtp struct {
	Phone string `json:"phone" validate:"required"`
	Type  string `json:"type" validate:"required,oneof=login register"`
}

// VerifyOtp command for verifying OTP
type VerifyOtp struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required,len=6"`
	Type  string `json:"type" validate:"required,oneof=login register"`
}
