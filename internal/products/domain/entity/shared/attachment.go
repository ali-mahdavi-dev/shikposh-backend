package shared

import (
	"shikposh-backend/pkg/framework/adapter"
)

type Attachment struct {
	adapter.BaseEntity
	AttachableType string `json:"attachable_type" gorm:"attachable_type"` // e.g., "ProductDetail", "Product"
	AttachableID   string `json:"attachable_id" gorm:"attachable_id"`
	FileType       string `json:"file_type" gorm:"file_type"` // e.g., "image", "document"
	FileName       string `json:"file_name" gorm:"file_name"`
	FilePath       string `json:"file_path" gorm:"file_path"`
	FileSize       int64  `json:"file_size" gorm:"file_size"`
	MimeType       string `json:"mime_type" gorm:"mime_type"`
	Order          int    `json:"order" gorm:"order;default:0"` // For ordering multiple attachments
}

func (a *Attachment) TableName() string {
	return "attachments"
}
