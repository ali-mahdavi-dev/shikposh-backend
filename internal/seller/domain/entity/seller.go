package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"

	"gorm.io/gorm"
)

type SellerID uint64

// JSONBMap is a custom type for JSONB map fields
type JSONBMap map[string]interface{}

// Value implements driver.Valuer interface
func (j JSONBMap) Value() (driver.Value, error) {
	if j == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(j)
}

// Scan implements sql.Scanner interface
func (j *JSONBMap) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONBMap)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*j = JSONBMap(result)
	return nil
}

// JSONBStringArray is a custom type for JSONB string array fields
type JSONBStringArray []string

// Value implements driver.Valuer interface
func (j JSONBStringArray) Value() (driver.Value, error) {
	if j == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(j)
}

// Scan implements sql.Scanner interface
func (j *JSONBStringArray) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONBStringArray, 0)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value")
	}

	var result []string
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*j = JSONBStringArray(result)
	return nil
}

// Seller represents a seller entity
type Seller struct {
	adapter.BaseEntity
	ID           SellerID `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `json:"name" gorm:"name;not null"`
	Avatar       string         `json:"avatar" gorm:"avatar"`
	Description  string         `json:"description" gorm:"description;type:text"`
	Rating       float64        `json:"rating" gorm:"rating;default:0"`
	Verified     bool           `json:"verified" gorm:"verified;default:false"`
	Location     string         `json:"location" gorm:"location"`
	ResponseTime string         `json:"response_time" gorm:"response_time"`
	// Social media links stored as JSONB
	SocialMedia JSONBMap `json:"social_media" gorm:"social_media;type:jsonb;default:'{}'::jsonb"`
	// Categories stored as JSONB array
	Categories JSONBStringArray `json:"categories" gorm:"categories;type:jsonb;default:'[]'::jsonb"`
}

func (s *Seller) TableName() string {
	return "sellers"
}

// ToMap converts Seller to map for JSON response
func (s *Seller) ToMap() map[string]interface{} {
	// Calculate stats (these would come from aggregations in real scenario)
	stats := map[string]interface{}{
		"total_reviews":   0, // TODO: calculate from reviews
		"average_rating":  s.Rating,
		"total_orders":    0, // TODO: calculate from orders
		"completion_rate": 0, // TODO: calculate from orders
	}

	// Extract social media
	socialMedia := map[string]interface{}{}
	if s.SocialMedia != nil {
		if instagram, ok := s.SocialMedia["instagram"].(string); ok && instagram != "" {
			socialMedia["instagram"] = instagram
		}
		if telegram, ok := s.SocialMedia["telegram"].(string); ok && telegram != "" {
			socialMedia["telegram"] = telegram
		}
		if website, ok := s.SocialMedia["website"].(string); ok && website != "" {
			socialMedia["website"] = website
		}
	}

	result := map[string]interface{}{
		"id":             s.ID,
		"name":           s.Name,
		"avatar":         s.Avatar,
		"description":    s.Description,
		"rating":         s.Rating,
		"total_products": 0, // TODO: calculate from products
		"join_date":      s.CreatedAt.Format(time.RFC3339),
		"verified":       s.Verified,
		"total_sales":    0, // TODO: calculate from orders
		"response_time":  s.ResponseTime,
		"location":       s.Location,
		"categories":     []string(s.Categories),
		"social_media":   socialMedia,
		"stats":          stats,
	}

	return result
}

// ToSummary converts Seller to summary format
func (s *Seller) ToSummary() map[string]interface{} {
	return map[string]interface{}{
		"id":             s.ID,
		"name":           s.Name,
		"avatar":         s.Avatar,
		"rating":         s.Rating,
		"total_products": 0, // TODO: calculate from products
		"verified":       s.Verified,
	}
}
