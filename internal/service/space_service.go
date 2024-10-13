package service

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type SpaceService interface {

	// CreateSpace creates a new space
	CreateSpace(Name, Description string, CoachID uuid.UUID, MemberCount uuid.UUID, SessionCount uuid.UUID, ourseCount uuid.UUID, Active bool) (*entity.Space, error)

	// UpdateSpace updates an existing space
	UpdateSpace(space *entity.Space) error

	// DeleteSpace deletes a space by its ID
	DeleteSpace(spaceID uuid.UUID) error

	// GetSpaceByID retrieves a space by its ID
	GetSpaceByID(spaceID uuid.UUID) (*entity.Space, error)

	GetAllSpaces() ([]*entity.Space, error)
}

// SpaceServiceImpl struct implementing CourseService

type spaceServiceImpl struct {
	repo      repository.SpaceRepository
	tokenRepo repository.TokenRepository
}

// GetAll implements SpaceService.
func (s *spaceServiceImpl) GetAllSpaces() ([]*entity.Space, error) {

	spaces, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all spaces: %v", err)
	}
	return spaces, nil

}

// NewSpaceService creates a new instance of SpaceService

func NewSpaceService(spaceRepo repository.SpaceRepository, tokenRepo repository.TokenRepository) SpaceService {
	return &spaceServiceImpl{
		repo:      spaceRepo,
		tokenRepo: tokenRepo,
	}
}

// CreateSpace implements SpaceService.
func (s *spaceServiceImpl) CreateSpace(Name, Description string, CoachID uuid.UUID, MemberCount uuid.UUID, SessionCount uuid.UUID, ourseCount uuid.UUID, Active bool) (*entity.Space, error) {
	// Generate uuid for new space
	neoSpace, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	// Create a new space instance
	newSpace := &entity.Space{
		ID:           neoSpace,
		Name:         Name,
		Description:  Description,
		CoachID:      CoachID,
		MemberCount:  MemberCount,
		SessionCount: SessionCount,
		CourseCount:  ourseCount,
		Active:       Active,
	}

	log.Printf("Creatinf Space: %+v", newSpace)

	err = s.repo.Create(newSpace)
	if err != nil {
		return nil, fmt.Errorf("failed to create course: %v", err)
	}

	return newSpace, nil

}

// DeleteSpace implements SpaceService.
func (s *spaceServiceImpl) DeleteSpace(spaceID uuid.UUID) error {

	_, err := s.repo.GetdByID(spaceID)
	if err != nil {
		return fmt.Errorf("could not find space with ID %s", spaceID)
	}
	if err := s.repo.Delete(spaceID); err != nil {
		return fmt.Errorf("failed to delete space with ID %s: %v", spaceID, err)
	}

	log.Printf("Successfully deleted space with ID %s", spaceID)

	return nil
}

// GetSpaceByID implements SpaceService.
func (s *spaceServiceImpl) GetSpaceByID(spaceID uuid.UUID) (*entity.Space, error) {

	space, err := s.repo.GetdByID(spaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get space with ID %s: %v", spaceID, err)
	}
	return space, nil

}

// UpdateSpace implements SpaceService.
func (s *spaceServiceImpl) UpdateSpace(space *entity.Space) error {
	_, err := s.repo.GetdByID(space.ID)
	if err != nil {
		return fmt.Errorf("could not find space with ID %s", space.ID)
	}

	if err := s.repo.Update(space); err != nil {
		return fmt.Errorf("failed to update space with ID %s: %v", space.ID, err)
	}

	return nil

}
