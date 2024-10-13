package routes

import (
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/repository"
	"dalabio/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSpacesRoutes(router *gin.Engine, spaceController *controller.SpaceController, tokenRepo repository.TokenRepository) {
	AuthMiddleware := middleware.AuthMiddleware(tokenRepo)

	spaceGroup := router.Group("/spaces")
	{
		spaceGroup.Use(AuthMiddleware)
		{
			spaceGroup.POST("", spaceController.CreateSpace)
			spaceGroup.PUT("/:id", spaceController.UpdateSpace)
			spaceGroup.DELETE("/:id", spaceController.DeleteSpace)
			spaceGroup.GET("/:id", spaceController.GetSpaceByID)
			spaceGroup.GET("", spaceController.GetAllSpaces)

		}
	}

}
