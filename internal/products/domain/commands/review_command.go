package commands

type CreateReview struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	UserID    uint64 `json:"user_id" validate:"required"`
	UserName  string `json:"user_name" validate:"required"`
	Rating    int    `json:"rating" validate:"required,min=1,max=5"`
	Comment   string `json:"comment" validate:"required,min=10"`
}

type UpdateReviewHelpful struct {
	ReviewID uint64 `json:"review_id" validate:"required"`
	Type     string `json:"type" validate:"required,oneof=helpful notHelpful"`
}
