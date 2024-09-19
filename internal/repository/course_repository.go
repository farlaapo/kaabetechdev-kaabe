package repository

import (
	"dalabio/internal/entity"

	"github.com/gofrs/uuid"
)

type CourseRepository interface {
	Create(user *entity.Course) error
	Update(user *entity.Course) error
	Delete(userID uuid.UUID) error
	FindByID(userID uuid.UUID) (*entity.Course, error)
	ListAll() ([]*entity.Course, error)
}
