package controller

import (
	"dalabio/internal/entity"
	"dalabio/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// CourseController struct that defines the course controller with its service
type CourseController struct {
	courseService service.CourseService
}

// NewCourseController creates a new CourseController instance
func NewCourseController(courseService service.CourseService) *CourseController {
	return &CourseController{courseService: courseService}
}

// CreateCourse handles the creation of a new course
func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course entity.Course
	// Bind JSON to the course struct, which now includes ContentURL as []string
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the request context set by the AuthMiddleware
	instructorID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Instructor ID is required"})
		return
	}

	// Call the service with the bound struct's ContentURL slice and instructorID
	createdCourse, err := c.courseService.CreateCourse(
		course.Title,
		course.Description,
		course.Duration,
		course.Category,
		course.Outline,
		course.ContentURL, // Pass the ContentURL slice directly
		course.Status,
		course.EnrolledCount,
		course.Version,
		instructorID.(uuid.UUID), // Cast to uuid.UUID, assuming instructorID is stored as a UUID
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdCourse)
}

// UpdateCourse handles the update of an existing course
func (cc *CourseController) UpdateCourse(ctx *gin.Context) {
	var course entity.Course

	// Parse and validate course ID from URL
	courseIdParam := ctx.Param("id")
	courseID, err := uuid.FromString(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Bind JSON input to course struct
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.ID = courseID

	// Call service to update course
	if err := cc.courseService.UpdateCourse(&course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// DeleteCourse handles the deletion of a course by ID
func (cc *CourseController) DeleteCourse(ctx *gin.Context) {
	// Parse and validate course ID from URL
	courseIdParam := ctx.Param("id")
	courseID, err := uuid.FromString(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Call service to delete course
	if err := cc.courseService.DeleteCourse(courseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// GetCourseByID handles retrieving a course by its ID
func (cc *CourseController) GetCourseByID(ctx *gin.Context) {
	// Parse and validate course ID from URL
	courseIdParam := ctx.Param("id")
	courseID, err := uuid.FromString(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Call service to get course
	course, err := cc.courseService.GetCourseByID(courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func (cc *CourseController) GetAllCourses(ctx *gin.Context) {

	// Call service to get courses
	courses, err := cc.courseService.GetAllCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// respond success
	ctx.JSON(http.StatusOK, courses)
}
