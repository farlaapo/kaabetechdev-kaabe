package routes

import (
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/repository"
	"dalabio/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCoursesRoutes(router *gin.Engine, courseController *controller.CourseController, tokenRepo repository.TokenRepository) {

	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	courseGroup := router.Group("/courses")
	{

		// Protected routes (require valid authentication)
		courseGroup.Use(authMiddleware)
		{
			courseGroup.POST("", courseController.CreateCourse)
			courseGroup.PUT("/:id", courseController.UpdateCourse)
			courseGroup.DELETE("/:id", courseController.DeleteCourse)
			courseGroup.GET("/:id", courseController.GetCourseByID)
			courseGroup.GET("", courseController.GetAllCourses)
		}
	}

}
