package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Payment struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id" binding:"required"`        // ID of the user making the payment
	OrderID        uuid.UUID `json:"order_id"`                          // Associated order ID if applicable
	Amount         float64   `json:"amount" binding:"required"`         // Payment amount
	Currency       string    `json:"currency" binding:"required"`       // Currency type (e.g., USD, EUR)
	PaymentMethod  string    `json:"payment_method" binding:"required"` // Method of payment (e.g., credit card, PayPal)
	TransactionID  string    `json:"transaction_id" gorm:"uniqueIndex"` // Unique transaction ID from the payment gateway
	Status         string    `json:"status" binding:"required"`         // Payment status (e.g., pending, completed, failed)
	PaymentGateway string    `json:"payment_gateway"`                   // Gateway used (e.g., Stripe, PayPal)
	PaymentDate    time.Time `json:"payment_date"`                      // Date when payment was made
	Notes          string    `json:"notes,omitempty"`                   // Any additional notes or metadata
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`  // Timestamp for when the payment record was created
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`  // Timestamp for when the payment record was last updated
}
