package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

// Course represents the structure of a course entity
type Course struct {
	ID            uuid.UUID  `json:"id,omitempty"`
	Title         string     `json:"title" binding:"required"`
	Description   string     `json:"description"`
	Duration      string     `json:"duration"`
	Version       uuid.UUID  `json:"version,omitempty"`
	Category      string     `json:"category"`
	InstructorID  uuid.UUID  `json:"instructor_id,omitempty"` // Added InstructorID field
	EnrolledCount int        `json:"enrolled_count,omitempty"`
	ContentURL    []string   `json:"content_url"` // Changed to a slice of strings
	Outline       string     `json:"outline,omitempty"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
