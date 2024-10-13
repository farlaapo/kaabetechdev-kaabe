package routes

import (
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/repository"
	"dalabio/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes sets up the routes for user-related operations.
func RegisterUserRoutes(router *gin.Engine, userController *controller.UserController, tokenRepo repository.TokenRepository) {
	// Apply middleware to protect certain routes
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	// User-related routes
	userGroup := router.Group("/users")
	{
		// Public routes
		userGroup.POST("", userController.RegisterUser)                  // Route for user registration
		userGroup.POST("/authenticate", userController.AuthenticateUser) // Route for user authentication

		// Protected routes (require valid authentication)
		userGroup.Use(authMiddleware) // Apply middleware here without additional braces
		{
			userGroup.PUT("/:id", userController.UpdateUser)    // Route for updating user information (protected)
			userGroup.DELETE("/:id", userController.DeleteUser) // Route for deactivating a user (protected)
			userGroup.GET("/:id", userController.GetUserByID)   // Route for getting a user by ID (protected)
			userGroup.GET("", userController.ListUsers)         // Route for listing all users (protected)
		}
	}
}
