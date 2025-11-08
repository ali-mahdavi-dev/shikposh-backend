package commands

type CreateReview struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	UserID    uint64 `json:"user_id" validate:"required"`
	UserName  string `json:"user_name" validate:"required"`
	Rating    int    `json:"rating" validate:"required,min=1,max=5"`
	Comment   string `json:"comment" validate:"required,min=10"`
}

// ReviewHelpfulType represents the type of helpful feedback for a review
type ReviewHelpfulType string

const (
	ReviewHelpfulTypeHelpful    ReviewHelpfulType = "helpful"
	ReviewHelpfulTypeNotHelpful ReviewHelpfulType = "notHelpful"
)

// IsValid checks if the ReviewHelpfulType is valid
func (t ReviewHelpfulType) IsValid() bool {
	return t == ReviewHelpfulTypeHelpful || t == ReviewHelpfulTypeNotHelpful
}

// String returns the string representation of the type
func (t ReviewHelpfulType) String() string {
	return string(t)
}

type UpdateReviewHelpful struct {
	ReviewID uint64            `json:"review_id" validate:"required"`
	Type     ReviewHelpfulType `json:"type" validate:"required"`
}
