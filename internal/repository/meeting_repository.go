package repository

import (
	"dalabio/internal/entity"

	"github.com/gofrs/uuid"
)

type MeetingRepository interface {
	// GetMeeting returns a meeting by its ID
	GetdByID(meetingID uuid.UUID) (*entity.Meeting, error)

	// CreateMeeting creates a new meeting
	Create(meeting *entity.Meeting) error

	// UpdateMeeting updates an existing meeting
	Update(meeting *entity.Meeting) error

	// DeleteMeeting deletes a meeting by its ID
	Delete(meetingID uuid.UUID) error

	// GetMeetings returns all meetings
	GetAll() ([]*entity.Meeting, error)
}
