package controller

import (
	"dalabio/internal/entity"
	"dalabio/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CourseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) *CourseController {
	return &CourseController{courseService: courseService}
}

func (Cc *CourseController) CreateCourse(ctx *gin.Context) {
	var course entity.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdCourse, err := Cc.courseService.CreateCourse(course.Title, course.Description, course.Duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdCourse)

}
func (Cc *CourseController) UpdateCourse(ctx *gin.Context) {
	var course entity.Course

	courseIdParam := ctx.Param("id")
	courseID, err := uuid.FromString(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.ID = courseID

	if err := Cc.courseService.UpdateCourse(&course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})

}
func (Cc *CourseController) DeleteCourse(ctx *gin.Context) {
	courseIdParam := ctx.Param("id")
	courseID, err := uuid.FromString(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := Cc.courseService.DeleteCourse(courseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meesage": "Course deletet successfully"})

}
func (Cc *CourseController) GetCourseByID(ctx *gin.Context) {
	courseIdParam := ctx.Param("id")
	courseID, err := uuid.FromString(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := Cc.courseService.GetCourseByID(courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}
