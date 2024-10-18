package repository

import (
	"dalabio/internal/entity"

	"github.com/gofrs/uuid"
)

type PaymentRepository interface {
	Create(payment *entity.Payment) error
	GetdByID(paymentID uuid.UUID) (*entity.Payment, error)
	GetAll() ([]*entity.Payment, error)
	Update(payment *entity.Payment) error
	Delete(paymentID uuid.UUID) error
}
