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
	CreateCourse(Title, Description, Duration, Category, ContentUR, Outline, Status string, EnrolledCount int, version uuid.UUID) (*entity.Course, error)
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

// func (s *courseServiceImpl) CreateCourse(Title string, Description string, Duration string) (*entity.Course, error) {

// }
func (s *courseServiceImpl) CreateCourse(Title, Description, Duration, Category, ContentUR, Outline, Status string, EnrolledCount int, version uuid.UUID) (*entity.Course, error) {

	// Generate a new UUID for the course ID
	neoCourse, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Create a new course instance
	newCourse := &entity.Course{
		ID:          neoCourse,
		Title:       Title,
		Description: Description,
		Duration:    Duration,
		Version:     version, // Use the provided version, not neoCourse.
		Category:    Category,
		// Use the passed instructor ID.
		EnrolledCount: EnrolledCount,
		ContentURL:    ContentUR,
		Outline:       Outline,
		Status:        Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

// GetCourseByID implements CourseService.

// UpdateCourse updates an existing course
func (s *courseServiceImpl) UpdateCourse(course *entity.Course) error {
	// Add your update logic here using s.repo
	_, err := s.repo.GetdByID(course.ID)
	if err != nil {
		// If the course does not exist, return the error
		return fmt.Errorf("could not find user with ID %s", course.ID)
	}

	// Call the repository to update the user
	if err := s.repo.Update(course); err != nil {
		return fmt.Errorf("failed to update user with ID %s: %v", course.ID, err)
	}

	return nil
}

// DeleteCourse deletes a course by its ID
func (s *courseServiceImpl) DeleteCourse(courseID uuid.UUID) error {
	_, err := s.repo.GetdByID(courseID)
	if err != nil {
		// If the user does not exist, return the error
		fmt.Printf("Could not find user with ID %s: %v", courseID, err)
		return fmt.Errorf("could not find user with ID %s", courseID)
	}

	// Call the repository to delete the user
	if err := s.repo.Delete(courseID); err != nil {
		fmt.Printf("Failed to delete user with ID %s: %v", courseID, err)
		return fmt.Errorf("failed to delete user with ID %s: %v", courseID, err)
	}

	fmt.Printf("Successfully deleted user with ID %s", courseID)
	return nil
}

func (s *courseServiceImpl) GetCourseByID(courseID uuid.UUID) (*entity.Course, error) {
	// Call the repository to find the course by its ID
	course, err := s.repo.GetdByID(courseID)
	if err != nil {
		return nil, fmt.Errorf("could not find user with ID %d: %v", courseID, err)
	}

	return course, nil

}
