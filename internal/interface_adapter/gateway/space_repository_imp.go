package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type spaceRepositoryImpl struct {
	db *sql.DB
}

func NewSpaceRepository(db *sql.DB) repository.SpaceRepository {
	return &spaceRepositoryImpl{db: db}
}

// Create implements repository.SpaceRepository.
func (r *spaceRepositoryImpl) Create(space *entity.Space) error {

	// Prepare the SQL statement

	query := `INSERT INTO spaces (id, name, description, coach_id, member_count, session_count, course_count, active, created_at, updated_at)
	   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	result, err := r.db.Exec(query, space.ID, space.Name, space.Description, space.CoachID, space.MemberCount, space.SessionCount, space.CourseCount, space.Active, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error inserting course: %v, query: %s", err, query)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}
	log.Printf("Rows affected: %d\n", rowsAffected)

	return nil

}

// Delete implements repository.SpaceRepository.
func (r *spaceRepositoryImpl) Delete(spaceID uuid.UUID) error {
	// Prepare the SQL statement

	query := `DELETE FROM spaces WHERE id = $1`
	result, err := r.db.Exec(query, spaceID)
	if err != nil {
		log.Printf("Error deleting course with ID: %v, error: %v", spaceID, err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No course found with ID: %v", spaceID)
		return nil
	}

	log.Printf("Deleted Space ID : %d\n", spaceID)

	return nil
}

// GetdByID implements repository.SpaceRepository.
func (r *spaceRepositoryImpl) GetdByID(spaceID uuid.UUID) (*entity.Space, error) {
	var space = entity.Space{}

	// Define the Space entity to store the result
	err := r.db.QueryRow(`	SELECT id, name, description, coach_id, member_count, session_count, course_count, active, created_at, updated_at
	FROM spaces WHERE id = $1`, spaceID).Scan(&space.ID, &space.Name, &space.Description, &space.CoachID, &space.MemberCount, &space.SessionCount, &space.CourseCount, &space.Active, &space.CreatedAt, &space.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No course found with ID: %v", spaceID)
			return nil, fmt.Errorf("course Not Found")
		}
		log.Printf("Error retrieving course by ID: %v", err)
	}

	return &space, nil
}

// Update implements repository.SpaceRepository.
func (r *spaceRepositoryImpl) Update(space *entity.Space) error {
	// Prepare the SQL statement

	query := `UPDATE spaces
	   SET name = $1, description = $2, coach_id = $3, member_count = $4, session_count = $5, course_count = $6, active = $7, updated_at = $8
	   WHERE id = $9`

	result, err := r.db.Exec(query, space.Name, space.Description, space.CoachID, space.MemberCount, space.SessionCount, space.CourseCount, space.Active, time.Now(), space.ID)
	if err != nil {
		log.Printf("Error updating course with ID: %v, error: %v", space.ID, err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No course found with ID: %v", space.ID)
		return nil
	}

	log.Printf("Course Updated: %d\n", space.ID)
	return nil

}
