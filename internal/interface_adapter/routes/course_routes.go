package routes

import (
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/repository"
	"dalabio/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCoursesRoutes(router *gin.Engine, courseController *controller.CourseController, tokenRepo repository.TokenRepository) {

	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	courseGroup := router.Group("/")
	{
		courseGroup.POST("/courses", courseController.CreateCourse)

		// Protected routes (require valid authentication)
		courseGroup.Use(authMiddleware)
		{
			courseGroup.PUT("/courses/:id", courseController.UpdateCourse)
			courseGroup.DELETE("/courses/:id", courseController.DeleteCourse)
			courseGroup.GET("/courses/:id", courseController.GetCourseByID)
		}
	}

}
