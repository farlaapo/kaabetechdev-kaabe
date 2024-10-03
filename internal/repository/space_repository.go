package repository

import (
	"dalabio/internal/entity"

	"github.com/gofrs/uuid"
)

type SpaceRepository interface {
	Create(space *entity.Space) error
	Update(space *entity.Space) error
	Delete(spaceID uuid.UUID) error
	GetdByID(spaceID uuid.UUID) (*entity.Space, error)
}
