package entity
import (
	"time"

	"github.com/gofrs/uuid"
	
)

type Space struct {
	ID              uuid.UUID      `json:"id" gorm:"primaryKey"` // Unique identifier for the space
	Name            string    `json:"name"`                 // The name of the space (e.g., "JavaScript Mastery")
	Description     string    `json:"description"`          // A brief description of the space
	CoachID         uuid.UUID       `json:"coach_id"`             // ID of the coach who owns the space
	MemberCount     uuid.UUID       `json:"member_count"`         // Number of members or clients in the space
	SessionCount    uuid.UUID       `json:"session_count"`        // Number of active sessions hosted in the space
	CourseCount     uuid.UUID       `json:"course_count"`         // Number of courses offered in the space
	Active          bool      `json:"active"`               // Indicates if the space is currently active or disabled 
	CreatedAt       time.Time `json:"created_at"`           // Time when the space was created
	UpdatedAt       time.Time `json:"updated_at"`           // Time when the space was last updated
}
