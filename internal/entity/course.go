package entity

import "time"

// Course represents a Course entity

type Course struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    string    `json:"duration"` // e.g., "3 weeks", "1 month"
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Instructor  string    `json:"instructor"`
}
