package commands

// SendOtp command for sending OTP
type SendOtp struct {
	Phone string `json:"phone" validate:"required"`
}

// VerifyOtp command for verifying OTP
type VerifyOtp struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required,len=6"`
}
