package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

type CourseRepositoryImpl struct {
	db *sql.DB
}

//  factory function to create an instance of CourseRepository

func NewCourseRepository(db *sql.DB) repository.CourseRepository {
	return &CourseRepositoryImpl{db: db}
}

func (r *CourseRepositoryImpl) Create(course *entity.Course) error {
	// Log the course information before the insert
	log.Printf("Inserting course: %+v", course)

	// SQL Query to insert the course into the database
	query := `INSERT INTO courses (id, title, description, duration, version, category, instructor_id, enrolled_count, content_url, outline, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	// Use pq.Array() to pass the slice of strings as a PostgreSQL array
	result, err := r.db.Exec(query, course.ID, course.Title, course.Description, course.Duration, course.Version, course.Category, course.InstructorID, course.EnrolledCount, pq.Array(course.ContentURL), course.Outline, course.Status, course.CreatedAt, course.UpdatedAt)
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
	result, err := r.db.Exec(`
    UPDATE courses 
    SET title = $2, description = $3, duration = $4, version = $5, category = $6, enrolled_count = $7, content_url = $8, status = $9, updated_at = CURRENT_TIMESTAMP 
    WHERE id = $1`,
		course.ID, course.Title, course.Description, course.Duration, course.Version, course.Category, course.EnrolledCount, pq.Array(course.ContentURL), course.Status)
	log.Printf("ContentURL: %+v", course.ContentURL)
	// query := `UPDATE courses
	// 		  SET  title = $1, description = $2, duration = $3, version = $4, category = $5, instructor_id = $6 enrolled_count = $7, content_url = $8, outline = $9, status = $10,  updated_at =  $11
	// 		  WHERE id = $12`
	// // Execute the update query with the course data
	// result, err := r.db.Exec(query, course.Title, course.Description, course.Duration, course.Version, course.Category, course.EnrolledCount, course.InstructorID, course.ContentURL, course.Outline, course.Status, course.UpdatedAt, course.ID)
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
	//  var course = entity.Course{}
	var course entity.Course
	query := "SELECT id, title, description, duration, version, category, instructor_id, enrolled_count, content_url, outline, status, created_at, updated_at, deleted_at FROM courses WHERE id = $1"
	err := r.db.QueryRow(query, courseID).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.Duration,
		&course.Version,
		&course.Category,
		&course.InstructorID,
		&course.EnrolledCount,
		pq.Array(&course.ContentURL), // Use pq.Array for TEXT[] in PostgreSQL
		&course.Outline,
		&course.Status,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.DeletedAt,
	)

	//Check for errors in retrieving the course
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No course found with ID: %v", courseID)
			return nil, fmt.Errorf("course Not Found")
		}
		log.Printf("Error retrieving course by ID: %v", err)
	}

	// Return the Course if found
	return &course, nil

}
