package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

// Course represents a Course entity
// think about carefully the course and make sure making more discriptive

type Course struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Duration    string     `json:"duration"` // e.g., "3 weeks", "1 month"
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
