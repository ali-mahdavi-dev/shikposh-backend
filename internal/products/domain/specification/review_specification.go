package specification

import (
	"shikposh-backend/internal/products/domain/entity"
	"github.com/ali-mahdavi-dev/framework/specification"
)

// ReviewHasMinimumRatingSpecification checks if a review has a minimum rating
type ReviewHasMinimumRatingSpecification struct {
	minRating int
}

func NewReviewHasMinimumRatingSpecification(minRating int) specification.Specification[*entity.Review] {
	return &ReviewHasMinimumRatingSpecification{
		minRating: minRating,
	}
}

func (s *ReviewHasMinimumRatingSpecification) IsSatisfiedBy(review *entity.Review) bool {
	return review != nil && review.Rating >= s.minRating
}

// ReviewIsVerifiedSpecification checks if a review is verified
type ReviewIsVerifiedSpecification struct {
}

func NewReviewIsVerifiedSpecification() specification.Specification[*entity.Review] {
	return &ReviewIsVerifiedSpecification{}
}

func (s *ReviewIsVerifiedSpecification) IsSatisfiedBy(review *entity.Review) bool {
	return review != nil && review.Verified
}

// ReviewHasCommentSpecification checks if a review has a comment
type ReviewHasCommentSpecification struct {
}

func NewReviewHasCommentSpecification() specification.Specification[*entity.Review] {
	return &ReviewHasCommentSpecification{}
}

func (s *ReviewHasCommentSpecification) IsSatisfiedBy(review *entity.Review) bool {
	return review != nil && review.Comment != ""
}

// ReviewIsHelpfulSpecification checks if a review is helpful
// A review is helpful if helpful count > not_helpful count
type ReviewIsHelpfulSpecification struct {
}

func NewReviewIsHelpfulSpecification() specification.Specification[*entity.Review] {
	return &ReviewIsHelpfulSpecification{}
}

func (s *ReviewIsHelpfulSpecification) IsSatisfiedBy(review *entity.Review) bool {
	return review != nil && review.Helpful > review.NotHelpful
}

// ReviewCanBePublishedSpecification checks if a review can be published
// A review can be published if it has:
// - A rating (1-5)
// - A comment
// - A valid user
type ReviewCanBePublishedSpecification struct {
}

func NewReviewCanBePublishedSpecification() specification.Specification[*entity.Review] {
	return &ReviewCanBePublishedSpecification{}
}

func (s *ReviewCanBePublishedSpecification) IsSatisfiedBy(review *entity.Review) bool {
	if review == nil {
		return false
	}

	// Must have valid rating (1-5)
	if review.Rating < 1 || review.Rating > 5 {
		return false
	}

	// Must have comment
	if review.Comment == "" {
		return false
	}

	// Must have user ID
	if review.UserID == 0 {
		return false
	}

	return true
}
