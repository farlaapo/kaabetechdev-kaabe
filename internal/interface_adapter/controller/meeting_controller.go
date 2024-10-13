package controller

import (
	"dalabio/internal/entity"
	"dalabio/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// MeetingController struct that defines the meeting controller with its service
type MeetingController struct {
	meetingService service.MeetingService
}

// NewMeetingController returns a new MeetingController
func NewMeetingController(meetingService service.MeetingService) *MeetingController {
	return &MeetingController{
		meetingService: meetingService,
	}
}

// GetAllMeetings returns all meetings
func (mc *MeetingController) GetAllMeetings(ctx *gin.Context) {
	meetings, err := mc.meetingService.GetAllMeetings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, meetings)
}

// GetMeetingByID returns a meeting by its ID
func (mc *MeetingController) GetMeetingByID(ctx *gin.Context) {

	// Parse and validate meeting ID from URL
	meetingIdParam := ctx.Param("id")
	meetingID, err := uuid.FromString(meetingIdParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	//call services
	meeting, err := mc.meetingService.GetMeetingByID(meetingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, meeting)

}

// CreateMeeting creates a new meeting
func (mc *MeetingController) CreateMeeting(ctx *gin.Context) {

	// Bind JSON input to meeting struct
	var meeting entity.Meeting

	if err := ctx.ShouldBindJSON(&meeting); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create meeting
	createMeeting, err := mc.meetingService.CreateMeeting(meeting.Title, meeting.Description, meeting.Duration, meeting.Location, meeting.MeetingType, meeting.Status, meeting.AttendeeIDs, meeting.AttendeeNames, meeting.AttendeeEmails, meeting.AttendeeStatus, meeting.JoinURL, meeting.MaximumCapacity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// respon with created meeting
	ctx.JSON(http.StatusOK, createMeeting)
}

// UpdateMeeting handles the update of an existing meeting
func (mc *MeetingController) UpdateMeeting(ctx *gin.Context) {
	var Meeting entity.Meeting

	// Parse and validate meeting ID from URL
	meetingIdParam := ctx.Param("id")
	meetingID, err := uuid.FromString(meetingIdParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Bind JSON input to meeting struct
	if err := ctx.ShouldBindJSON(&Meeting); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Meeting.ID = meetingID
	// Call service to update meeting
	if err := mc.meetingService.UpdateMeeting(&Meeting); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meeting Updated successfully"})

}

// DeleteMeeting handles the deletion of a meeting by ID
func (mc *MeetingController) DeleteMeeting(ctx *gin.Context) {

	// Parse and validate meeting ID from URL
	meetingIdParam := ctx.Param("id")
	meetingID, err := uuid.FromString(meetingIdParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	// Call service to delete meeting
	if err := mc.meetingService.DeleteMeeting(meetingID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meeting Deleted Sucessfull"})
}
