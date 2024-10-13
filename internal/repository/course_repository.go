package repository

import (
	"dalabio/internal/entity"

	"github.com/gofrs/uuid"
)

// CourseRepository interface with required methods
type CourseRepository interface {
	Create(course *entity.Course) error
	Update(course *entity.Course) error
	Delete(courseID uuid.UUID) error
	GetdByID(courseID uuid.UUID) (*entity.Course, error)
	GetAll() ([]*entity.Course, error)
}
