package commands

// user
type RegisterUser struct {
	Phone            string `json:"phone" validate:"required"`
	AvatarIdentifier string `json:"avatar_identifier"`
	UserName         string `json:"user_name" validate:"required,min=3"`
	FirstName        string `json:"first_name" validate:"required,min=3"`
	LastName         string `json:"last_name" validate:"required,min=3"`
	Email            string `json:"email" validate:"omitempty,email"`
	Password         string `json:"password"` // Optional for OTP-based registration
}

type LoginUser struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Logout struct {
	UserID uint64 `json:"user_id" validate:"required"`
}
