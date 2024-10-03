package controller

import (
	"dalabio/internal/entity"
	"dalabio/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// SpaceController struct that defines the space controller with its service
type SpaceController struct {
	spaceService service.SpaceService
}

func NewSpaceController(spaceService service.SpaceService) *SpaceController {
	return &SpaceController{spaceService: spaceService}
}

// CreateSpace handles the creation of a new space
func (sc *SpaceController) CreateSpace(ctx *gin.Context) {
	var space entity.Space

	// Bind JSON to the spcae struct,
	if err := ctx.ShouldBindJSON(&space); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createSpace, err := sc.spaceService.CreateSpace(
		space.Name,
		space.Description,
		space.CoachID,
		space.MemberCount,
		space.SessionCount,
		space.CourseCount,
		space.Active,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createSpace)
}

// UpdateCourse handles the update of an existing space
func (sc *SpaceController) UpdateSpace(ctx *gin.Context) {
	var space entity.Space

	// Parse and validate space ID from URL
	spaceIdParms := ctx.Param("id")
	spaceID, err := uuid.FromString(spaceIdParms)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid space ID"})
		return
	}
	// Bind JSON input to space struct
	if err := ctx.ShouldBindJSON(&space); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	space.ID = spaceID
	// Call service to update space
	if err := sc.spaceService.UpdateSpace(&space); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Space Updated successfully"})

}
func (sc *SpaceController) DeleteSpace(ctx *gin.Context) {

	// Parse and validate space ID from URL
	spaceIdParms := ctx.Param("id")
	spaceID, err := uuid.FromString(spaceIdParms)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid space ID"})
		return
	}

	// Call service to update space
	if err := sc.spaceService.DeleteSpace(spaceID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Space Deleted Sucessfull"})

}

func (sc *SpaceController) GetSpaceByID(ctx *gin.Context) {
	// Parse and validate space ID from URL
	spaceIdParms := ctx.Param("id")
	spaceID, err := uuid.FromString(spaceIdParms)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid space ID"})
		return
	}
	// Call service to update space
	space, err := sc.spaceService.GetSpaceByID(spaceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, space)

}
