package service

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type PaymentService interface {

	// CreatePayment creates a new payment
	CreatePayment(UserID uuid.UUID, OrderID uuid.UUID, Amount float64, Currency string, PaymentMethod string, TransactionID string, Status string, PaymentGateway string, Notes string) (*entity.Payment, error)

	// GetPaymentByID gets a payment by ID
	GetPaymentByID(paymentID uuid.UUID) (*entity.Payment, error)

	// GetAllPayments gets all payments
	GetAllPayments() ([]*entity.Payment, error)

	// UpdatePayment updates a payment
	UpdatePayment(payment *entity.Payment) error

	// DeletePayment deletes a payment
	DeletePayment(paymentID uuid.UUID) error
}

type paymentServiceImpl struct {
	repo      repository.PaymentRepository
	repotoken repository.TokenRepository
}

// DeletePayment implements PaymentService.
func (s *paymentServiceImpl) DeletePayment(paymentID uuid.UUID) error {
	_, err := s.repo.GetdByID(paymentID)

	if err != nil {
		return fmt.Errorf("could not find payment with ID %s: %v", paymentID, err)
	}

	if err := s.repo.Delete(paymentID); err != nil {
		return fmt.Errorf("failed to delete payment with ID %s: %v", paymentID, err)
	}

	log.Printf("Successfully deleted payment with ID %s", paymentID)
	return nil
}

// CreatePayment implements PaymentService.
func (s *paymentServiceImpl) CreatePayment(UserID uuid.UUID, OrderID uuid.UUID, Amount float64, Currency string, PaymentMethod string, TransactionID string, Status string, PaymentGateway string, Notes string) (*entity.Payment, error) {
	{
		neoPayment, err := uuid.NewV4()

		if err != nil {
			return nil, err
		}

		newPayment := &entity.Payment{
			ID:             neoPayment,
			UserID:         UserID,
			OrderID:        OrderID,
			Amount:         Amount,
			Currency:       Currency,
			PaymentMethod:  PaymentMethod,
			TransactionID:  TransactionID,
			PaymentGateway: PaymentGateway,
			Notes:          Notes,
		}
		log.Printf("insering payment: %v", neoPayment)

		err = s.repo.Create(newPayment)

		if err != nil {
			return nil, err
		}

		return newPayment, nil

	}
}

// GetAllPayments implements PaymentService.
func (s *paymentServiceImpl) GetAllPayments() ([]*entity.Payment, error) {
	payment, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// GetPaymentByID implements PaymentService.
func (s *paymentServiceImpl) GetPaymentByID(paymentID uuid.UUID) (*entity.Payment, error) {

	payment, err := s.repo.GetdByID(paymentID)

	if err != nil {
		return nil, fmt.Errorf("could not find payment with ID %s: %v", paymentID, err)
	}

	return payment, nil
}

// UpdatePayment implements PaymentService.
func (s *paymentServiceImpl) UpdatePayment(payment *entity.Payment) error {
	_, err := s.repo.GetdByID(payment.ID)

	if err != nil {
		return fmt.Errorf("could not find payment with ID %s: %v", payment.ID, err)
	}

	if err := s.repo.Update(payment); err != nil {
		return fmt.Errorf("failed to update payment with ID %s: %v", payment.ID, err)
	}

	return nil
}

func NewPaymentService(paymentRepo repository.PaymentRepository, repotoken repository.TokenRepository) PaymentService {
	return &paymentServiceImpl{
		repo:      paymentRepo,
		repotoken: repotoken,
	}
}
