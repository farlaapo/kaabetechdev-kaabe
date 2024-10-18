package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"errors"
	"log"

	"github.com/gofrs/uuid"
)

type PaymentRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.PaymentRepository.
func (r *PaymentRepositoryImpl) Create(payment *entity.Payment) error {
	//query  insert

	query := `INSERT INTO payments(id, user_id, order_id, amount, currency, payment_method, transaction_id, status, payment_gateway, payment_date,  notes, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	result, err := r.db.Exec(query, payment.ID, payment.UserID, payment.OrderID, payment.Amount, payment.Currency, payment.PaymentMethod, payment.TransactionID, payment.Status, payment.PaymentGateway, payment.PaymentDate, payment.Notes, payment.CreatedAt, payment.UpdatedAt)

	if err != nil {
		log.Printf("Error inserting course: %v, query: %s", err, query)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
	}

	log.Printf("Rows affected: %d\n", rowsAffected)

	return nil
}

// Delete implements repository.PaymentRepository.
func (r *PaymentRepositoryImpl) Delete(paymentID uuid.UUID) error {
	//query  delete

	query := `DELETE FROM payments WHERE id = $1`

	result, err := r.db.Exec(query, paymentID)

	if err != nil {
		log.Printf("Error deleting course with ID: %v, error: %v", paymentID, err)
		return nil
	}

	// Check how many rows were affected by the delete
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return nil
	}

	// If no rows were affected, it means the course was not found
	if rowsAffected == 0 {
		log.Printf("No course found with ID: %v", paymentID)
		return nil
	}

	return nil

}

// GetAll implements repository.PaymentRepository.
func (r *PaymentRepositoryImpl) GetAll() ([]*entity.Payment, error) {
	var payments []*entity.Payment
	//query select

	query := ` SELECT id, user_id, order_id, amount, currency, payment_method, transaction_id, status, payment_gateway, payment_date,notes, created_at, updated_at FROM payments`
	rows, err := r.db.Query(query)

	if err != nil {
		log.Printf("Error fetching courses: %v, query: %s", err, query)
	}

	defer rows.Close()

	for rows.Next() {
		var payment entity.Payment
		err := rows.Scan(
			&payment.ID,
			&payment.UserID,
			&payment.OrderID,
			&payment.Amount,
			&payment.Currency,
			&payment.PaymentMethod,
			&payment.TransactionID,
			&payment.Status,
			&payment.PaymentGateway,
			&payment.PaymentDate,
			&payment.Notes,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning course: %v", err)
		}

		payments = append(payments, &payment)
	}

	return payments, nil

}

// GetdByID implements repository.PaymentRepository.
func (r *PaymentRepositoryImpl) GetdByID(paymentID uuid.UUID) (*entity.Payment, error) {
	var payment entity.Payment
	// query select ID
	query := `SELECT id, user_id, order_id, amount, currency, payment_method, transaction_id, status, payment_gateway, payment_date,  notes, created_at, updated_at FROM payments WHERE id = $1`
	err := r.db.QueryRow(query, paymentID).Scan(
		&payment.ID,
		&payment.UserID,
		&payment.OrderID,
		&payment.Amount,
		&payment.Currency,
		&payment.PaymentMethod,
		&payment.TransactionID,
		&payment.Status,
		&payment.PaymentGateway,
		&payment.PaymentDate,
		&payment.Notes,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Error fetching course with ID: %v, error: %v", paymentID, err)
			return nil, errors.New("payment not found")
		}

		log.Printf("Error fetching course with ID: %v, error: %v", paymentID, err)
	}

	return &payment, nil
}

// Update implements repository.PaymentRepository.
func (r *PaymentRepositoryImpl) Update(payment *entity.Payment) error {
	//query update
	query := `UPDATE
		payments
		SET user_id = $2, order_id = $3, amount = $4, currency = $5, payment_method = $6, transaction_id = $7, status = $8, payment_gateway = $9, payment_date = $10,  notes = $11, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 
		`
	result, err := r.db.Exec(query,
		payment.ID,
		payment.UserID,
		payment.OrderID,
		payment.Amount,
		payment.Currency,
		payment.PaymentMethod,
		payment.TransactionID,
		payment.Status,
		payment.PaymentGateway,
		payment.PaymentDate,
		payment.Notes,
	)

	if err != nil {
		log.Printf("Error updating payment with ID: %v, error: %v", payment.ID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
	}

	log.Printf("Rows affected: %d\n", rowsAffected)

	return nil

}

func NewPaymentRepository(db *sql.DB) repository.PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}
