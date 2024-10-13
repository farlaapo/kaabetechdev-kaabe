package routes

import (
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/repository"
	"dalabio/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMeetingRoutes(router *gin.Engine, meetingController *controller.MeetingController, tokenRepository repository.TokenRepository) {
	authMiddleware := middleware.AuthMiddleware(tokenRepository)

	meetingGroup := router.Group("/meetings")
	{
		meetingGroup.Use(authMiddleware)
		{
			meetingGroup.POST("", meetingController.CreateMeeting)
			meetingGroup.PUT("/:id", meetingController.UpdateMeeting)
			meetingGroup.DELETE("/:id", meetingController.DeleteMeeting)
			meetingGroup.GET("/:id", meetingController.GetMeetingByID)
			meetingGroup.GET("", meetingController.GetAllMeetings)
		}
	}

}
