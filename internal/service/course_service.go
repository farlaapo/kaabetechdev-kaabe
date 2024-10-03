package service

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// CourseService interface
type CourseService interface {
	CreateCourse(Title, Description, Duration, Category, Outline string, ContentURLs []string, Status string, EnrolledCount int, version uuid.UUID, instructorID uuid.UUID) (*entity.Course, error)
	UpdateCourse(course *entity.Course) error
	DeleteCourse(courseID uuid.UUID) error
	GetCourseByID(courseID uuid.UUID) (*entity.Course, error)
}

// courseServiceImpl struct implementing CourseService
type courseServiceImpl struct {
	repo      repository.CourseRepository
	tokenRepo repository.TokenRepository
}

// NewCourseService creates a new instance of CourseService
func NewCourseService(coureRepo repository.CourseRepository, tokenRepo repository.TokenRepository) CourseService {
	return &courseServiceImpl{
		repo:      coureRepo,
		tokenRepo: tokenRepo,
	}
}

func (s *courseServiceImpl) CreateCourse(Title, Description, Duration, Category, Outline string, ContentURLs []string, Status string, EnrolledCount int, version uuid.UUID, instructorID uuid.UUID) (*entity.Course, error) {
	// Generate a new UUID for the course ID
	neoCourse, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Create a new course instance
	newCourse := &entity.Course{
		ID:            neoCourse,
		Title:         Title,
		Description:   Description,
		Duration:      Duration,
		Version:       version,
		Category:      Category,
		InstructorID:  instructorID, // Make sure you pass the instructorID
		EnrolledCount: EnrolledCount,
		ContentURL:    ContentURLs,
		Outline:       Outline,
		Status:        Status,
		CreatedAt:     time.Now(), // Set created_at
		UpdatedAt:     time.Now(), // Set updated_at
	}

	// Log the new course creation attempt
	log.Printf("Creating course: %+v", newCourse)

	// Save the new course to the repository
	err = s.repo.Create(newCourse)
	if err != nil {
		return nil, fmt.Errorf("failed to create course: %v", err)
	}

	return newCourse, nil
}

// UpdateCourse updates an existing course
func (s *courseServiceImpl) UpdateCourse(course *entity.Course) error {
	_, err := s.repo.GetdByID(course.ID)
	if err != nil {
		return fmt.Errorf("could not find course with ID %s", course.ID)
	}

	if err := s.repo.Update(course); err != nil {
		return fmt.Errorf("failed to update course with ID %s: %v", course.ID, err)
	}

	return nil
}

// DeleteCourse deletes a course by its ID
func (s *courseServiceImpl) DeleteCourse(courseID uuid.UUID) error {
	_, err := s.repo.GetdByID(courseID)
	if err != nil {
		return fmt.Errorf("could not find course with ID %s: %v", courseID, err)
	}

	if err := s.repo.Delete(courseID); err != nil {
		return fmt.Errorf("failed to delete course with ID %s: %v", courseID, err)
	}

	log.Printf("Successfully deleted course with ID %s", courseID)
	return nil
}

// GetCourseByID retrieves a course by its ID
func (s *courseServiceImpl) GetCourseByID(courseID uuid.UUID) (*entity.Course, error) {
	course, err := s.repo.GetdByID(courseID)
	if err != nil {
		return nil, fmt.Errorf("could not find course with ID %s: %v", courseID, err)
	}

	return course, nil
}
