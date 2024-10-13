package service

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type MeetingService interface {
	// GetMeeting returns a meeting by its ID
	GetMeetingByID(meetingID uuid.UUID) (*entity.Meeting, error)

	// GetMeetings returns all meetings
	GetAllMeetings() ([]*entity.Meeting, error)

	// CreateMeeting creates a new meeting
	CreateMeeting(Title, Description, Duration, Location, MeetingType, Status string, AttendeeIDs []uuid.UUID, AttendeeNames []string, AttendeeEmails []string, AttendeeStatus []string, JoinURL []string, MaximumCapacity int) (*entity.Meeting, error)

	// UpdateMeeting updates an existing meeting
	UpdateMeeting(meeting *entity.Meeting) error

	// DeleteMeeting deletes a meeting by its ID
	DeleteMeeting(meetingID uuid.UUID) error
}

type meetingService struct {
	repo     repository.MeetingRepository
	tokenRep repository.TokenRepository
}

// GetAllMeetings implements MeetingService.
func (s *meetingService) GetAllMeetings() ([]*entity.Meeting, error) {
	meeting, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return meeting, nil

}

// CreateMeeting implements MeetingService.
func (s *meetingService) CreateMeeting(Title, Description, Duration, Location, MeetingType, Status string, AttendeeIDs []uuid.UUID, AttendeeNames []string, AttendeeEmails []string, AttendeeStatus []string, JoinURL []string, MaximumCapacity int) (*entity.Meeting, error) {
	neoMeeting, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	newMeeting := &entity.Meeting{
		ID:              neoMeeting,
		Title:           Title,
		Description:     Description,
		Duration:        Duration,
		MeetingType:     MeetingType,
		Status:          Status,
		Location:        Location,
		AttendeeIDs:     AttendeeIDs,
		AttendeeNames:   AttendeeNames,
		AttendeeEmails:  AttendeeEmails,
		AttendeeStatus:  AttendeeStatus,
		JoinURL:         JoinURL,
		MaximumCapacity: MaximumCapacity,
	}

	log.Printf("Creating course: %+v", neoMeeting)

	err = s.repo.Create(newMeeting)

	if err != nil {
		return nil, err
	}

	return newMeeting, nil

}

// UpdateMeeting implements MeetingService.
func (s *meetingService) UpdateMeeting(meeting *entity.Meeting) error {
	_, err := s.repo.GetdByID(meeting.ID)

	if err != nil {
		return fmt.Errorf("could not find course with ID %s", meeting.ID)
	}

	if err := s.repo.Update(meeting); err != nil {
		return fmt.Errorf("failed to update course with ID %s: %v", meeting.ID, err)
	}

	return nil

}

// DeleteMeeting implements MeetingService.
func (s *meetingService) DeleteMeeting(meetingID uuid.UUID) error {

	_, err := s.repo.GetdByID(meetingID)
	if err != nil {
		return fmt.Errorf("could not find course with ID %s: %v", meetingID, err)
	}

	if err := s.repo.Delete(meetingID); err != nil {
		return fmt.Errorf("failed to delete course with ID %s: %v", meetingID, err)
	}

	log.Printf("Successfully deleted course with ID %s", meetingID)
	return nil
}

// GetMeetingByID implements MeetingService.
func (s *meetingService) GetMeetingByID(meetingID uuid.UUID) (*entity.Meeting, error) {
	meeting, err := s.repo.GetdByID(meetingID)

	if err != nil {
		return nil, fmt.Errorf("could not find course with ID %s: %v", meetingID, err)
	}

	return meeting, nil

}

func NewMeetingService(meetingRepo repository.MeetingRepository, tokenRep repository.TokenRepository) MeetingService {
	return &meetingService{
		repo:     meetingRepo,
		tokenRep: tokenRep,
	}
}
