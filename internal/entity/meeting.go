package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Meeting struct {
	ID              uuid.UUID   `json:"id"`
	Title           string      `json:"title" binding:"required"`
	Description     string      `json:"description,omitempty"`
	Duration        string      `json:"duration" binding:"required"`
	StartTime       time.Time   `json:"start_time" binding:"required"`
	EndTime         time.Time   `json:"end_time" binding:"required"`
	Location        string      `json:"location,omitempty"`
	AttendeeIDs     []uuid.UUID `json:"attendee_ids,omitempty"`     // List of attendee user IDs
	AttendeeNames   []string    `json:"attendee_names,omitempty"`   // List of attendee names
	AttendeeEmails  []string    `json:"attendee_emails,omitempty"`  // List of attendee emails
	AttendeeStatus  []string    `json:"attendee_status,omitempty"`  // Corresponding attendee statuses (e.g., "invited", "joined")
	MeetingType     string      `json:"meeting_type"`               // e.g., "virtual", "in-person"
	Status          string      `json:"status"`                     // e.g., "scheduled", "ongoing", "completed", "cancelled"
	JoinURL         []string    `json:"join_url,omitempty"`         // Virtual meeting link if applicable
	MaximumCapacity int         `json:"maximum_capacity,omitempty"` // Maximum number of attendees allowed
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	DeletedAt       *time.Time  `json:"deleted_at,omitempty"` // For soft deletes
}
