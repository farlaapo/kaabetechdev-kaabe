package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type CourseRepositoryImpl struct {
	db *sql.DB
}

//  factory function to create an instance of CourseRepository

func NewCourseRepository(db *sql.DB) repository.CourseRepository {
	return &CourseRepositoryImpl{db: db}
}

func (r *CourseRepositoryImpl) Create(course *entity.Course) error {
	// Generate UUID for the course
	newUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)
		return err
	}
	course.ID = newUUID

	// Log the course information before the insert
	log.Printf("Inserting course: %+v", course)

	// SQL Query to insert the course into the database
	query := `INSERT INTO courses (id, title, description, duration, version, category, enrolled_count, content_url, outline, status, created_at, updated_at)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 )`

	result, err := r.db.Exec(query, course.ID, course.Title, course.Description, course.Duration, course.Version, course.Category, course.EnrolledCount, course.ContentURL, course.Outline, course.Status, course.CreatedAt, course.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting course: %v, query: %s", err, query)
		return err
	}

	// Check rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	log.Printf("Rows affected: %d\n", rowsAffected)
	return nil
}

// Update implements repository.CourseRepository.
func (r *CourseRepositoryImpl) Update(course *entity.Course) error {
	// Define the SQL update query
	query := `UPDATE courses
			  SET  title = $1, description = $2, duration = $3, version = $4, category = $5,  enrolled_count = $6, content_url = $7, outline = $8, status = $9,  updated_at =  $10
			  WHERE id = $11`
	// Execute the update query with the course data
	result, err := r.db.Exec(query, course.Title, course.Description, course.Duration, course.Version, course.Category, course.EnrolledCount, course.ContentURL, course.Outline, course.Status, course.UpdatedAt, course.ID)
	if err != nil {
		log.Printf("Error updating course with ID: %v, error: %v", course.ID, err)
		return err
	}
	// If no rows were affected, it means the course was not found
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}
	// If no rows were affected, it means the course was not found
	if rowsAffected == 0 {
		log.Printf("No course found with ID: %v", course.ID)
		return nil
	}

	log.Printf("Course Updated: %d\n", course.ID)
	return nil

}

func (r *CourseRepositoryImpl) Delete(courseID uuid.UUID) error {
	// SQL Query to insert the course into the database
	query := `DELETE FROM courses WHERE id = $1`
	result, err := r.db.Exec(query, courseID)
	if err != nil {
		log.Printf("Error deleting course with ID: %v, error: %v", courseID, err)
		return err
	}
	// Check how many rows were affected by the delete
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}
	// If no rows were affected, it means the course was not found
	if rowsAffected == 0 {
		log.Printf("No course found with ID: %v", courseID)
		return nil
	}

	log.Printf("Deleted Course ID : %d\n", courseID)
	return nil

}

// GetdByID implements repository.CourseRepository.
func (r *CourseRepositoryImpl) GetdByID(courseID uuid.UUID) (*entity.Course, error) {
	// Define the Course entity to store the result
	var course = entity.Course{}

	// Fetch the course from the database using the provided ID
	err := r.db.QueryRow(`SELECT id, title, description, duration, version, category,  enrolled_count, content_url, outline, status, created_at, updated_at
	FROM courses WHERE id = $1`, courseID).Scan(
		&course.ID, &course.Title, &course.Description, &course.Duration, &course.Version, &course.Category, &course.EnrolledCount, &course.ContentURL, &course.Outline, &course.Status, &course.CreatedAt, &course.UpdatedAt)

	// Check for errors in retrieving the course
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No course found with ID: %v", courseID)
			return nil, fmt.Errorf("Course Not Found")
		}
		log.Printf("Error retrieving course by ID: %v", err)
	}
	//	// Return the Course if found
	return &course, nil

}
