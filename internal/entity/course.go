package entity

import (
	// "database/sql/driver"
	// "encoding/json"
	// "fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Add custom marshal/unmarshal methods to the ContentURL field
type Course struct {
	ID            uuid.UUID  `json:"id,omitempty"`
	Title         string     `json:"title" binding:"required"` // Still required in the request
	Description   string     `json:"description"`              // Not required in the request
	Duration      string     `json:"duration"`
	Version       uuid.UUID  `json:"version,omitempty"`
	Category      string     `json:"category"`
	EnrolledCount int        `json:"enrolled_count,omitempty"`
	ContentURL    string     `json:"content_url"`
	Outline       string     `json:"outline,omitempty"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
