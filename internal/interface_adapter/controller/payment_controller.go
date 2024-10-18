package controller

import (
	"dalabio/internal/entity"
	"dalabio/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// PaymentController struct that defines the payment controller with its service
type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService}
}

func (pc *PaymentController) CreatePayment(ctx *gin.Context) {

	var payment entity.Payment

	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentRequest, err := pc.paymentService.CreatePayment(payment.UserID, payment.OrderID, payment.Amount, payment.Currency, payment.PaymentMethod, payment.TransactionID, payment.Status, payment.PaymentGateway, payment.Notes)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, paymentRequest)

}

func (pc *PaymentController) GetPaymentByID(ctx *gin.Context) {

	// Parse and validate payment ID from URL
	paymentIdParam := ctx.Param("id")
	paymentID, err := uuid.FromString(paymentIdParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	// Call service to get payment
	payment, err := pc.paymentService.GetPaymentByID(paymentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

func (pc *PaymentController) GetAllPayments(ctx *gin.Context) {

	payments, err := pc.paymentService.GetAllPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

func (pc *PaymentController) UpdatePayment(ctx *gin.Context) {

	var payment entity.Payment
	// Parse and validate payment ID from URL
	paymentIdParam := ctx.Param("id")
	paymentID, err := uuid.FromString(paymentIdParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	// Bind JSON input to payment struct

	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.ID = paymentID

	// Call service to update payment
	if err := pc.paymentService.UpdatePayment(&payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment Updated successfully"})

}

func (pc *PaymentController) DeletePayment(ctx *gin.Context) {

	// Parse and validate payment ID from URL
	paymentIdParam := ctx.Param("id")
	paymentID, err := uuid.FromString(paymentIdParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	// Call service to delete payment
	if err := pc.paymentService.DeletePayment(paymentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment Deleted Sucessfull"})

}
