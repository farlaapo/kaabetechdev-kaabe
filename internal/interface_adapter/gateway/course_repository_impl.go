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

// Delete implements repository.CourseRepository.

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

	// SQL Query to insert the course into the database
	query := `
		INSERT INTO courses (id, title, description, duration, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6)
		`
	result, err := r.db.Exec(query, course.ID, course.Title, course.Description, course.Duration, course.CreatedAt, course.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting course: %v", err)
		return err
	}

	// Check rows affected
	rowsAAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	log.Printf("Rows affected: %d\n", rowsAAffected)
	return nil
}

// Update implements repository.CourseRepository.
func (r *CourseRepositoryImpl) Update(course *entity.Course) error {
	// Define the SQL update query
	query := `UPDATE courses
			  SET title = $1, description = $2, duration = $3, updated_at = 4,
			  WHERE id = $5`
	// Execute the update query with the course data
	result, err := r.db.Exec(query, course.Title, course.Description, course.Duration, course.UpdatedAt, course.ID)
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
	err := r.db.QueryRow(`SELECT id, title, description, duration, created_at, updated_at
	FROM courses WHERE id = $1`, courseID).Scan(
		&course.ID, &course.Title, &course.Description, &course.Duration, &course.CreatedAt, &course.UpdatedAt)

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
