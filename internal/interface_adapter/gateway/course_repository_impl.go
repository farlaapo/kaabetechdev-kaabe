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

// courseRepositoryImp  struct  of courseRepository
type courseRepositoryImp struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) repository.CourseRepository {
	return &courseRepositoryImp{db: db}
}

func (c courseRepositoryImp) Create(course *entity.Course) error {
	// Generate a new UUID for the user
	newUud, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error generating UUID: %v", err)
		return err
	}
	course.ID = newUud

	fmt.Println("Inserting course: ID=%s, Title=%s, Description=%s, Duration=%s, StartDate=%s, EndDate=%s, Instructor=%v\n",
		course.ID, course.Title, course.Description, course.Duration, course.StartDate, course.EndDate, course.Instructor)

	query := `INSERT INTO users (id, title, description, duration, end_time, last_name, is_active, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	result, err := c.db.Exec(query, course.ID, course.Title, course.Description, course.Duration, course.StartDate, course.EndDate, course.Instructor, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	} else {
		log.Printf("Rows affected: %d", rowsAffected)
	}

	var insertedUser entity.User
	err = r.db.QueryRow(`SELECT id, username, email, password, first_name, last_name, is_active, created_at, updated_at
                         FROM users WHERE email = $1`, user.Email).Scan(
		&insertedUser.ID, &insertedUser.Username, &insertedUser.Email, &insertedUser.Password,
		&insertedUser.FirstName, &insertedUser.LastName, &insertedUser.IsActive, &insertedUser.CreatedAt, &insertedUser.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving inserted user: %v", err)
		return err
	}

	log.Printf("Inserted User: %+v", insertedUser)

	return nil
}
