package repository

import (
	"dalabio/internal/entity"

	"github.com/gofrs/uuid"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(userID uuid.UUID) error
	FindByID(userID uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	ListAll() ([]*entity.User, error)
}
