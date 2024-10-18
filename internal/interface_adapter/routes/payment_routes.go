package routes

import (
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/repository"
	"dalabio/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.Engine, spaceController *controller.PaymentController, tokenRepo repository.TokenRepository) {
	AuthMiddleware := middleware.AuthMiddleware(tokenRepo)

	spaceGroup := router.Group("/payments")
	{
		spaceGroup.Use(AuthMiddleware)
		{
			spaceGroup.POST("", spaceController.CreatePayment)
			spaceGroup.PUT("/:id", spaceController.UpdatePayment)
			spaceGroup.DELETE("/:id", spaceController.DeletePayment)
			spaceGroup.GET("/:id", spaceController.GetPaymentByID)
			spaceGroup.GET("", spaceController.GetAllPayments)

		}
	}

}
