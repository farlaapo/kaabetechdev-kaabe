package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// userRepositoryImpl is the implementation of UserRepository.
type userRepositoryImpl struct {
	db *sql.DB
}

// Delete implements repository.UserRepository.
// NewUserRepository creates a new instance of UserRepositoryImpl.
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

// Create inserts a new user into the database.

func (r *userRepositoryImpl) Create(user *entity.User) error {
	// Generate a new UUID for the user
	newUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)
		return err
	}
	user.ID = newUUID

	log.Printf("Inserting User: ID=%s, Username=%s, Email=%s, Password=%s, FirstName=%s, LastName=%s, IsActive=%v\n",
		user.ID, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.IsActive)

	query := `INSERT INTO users (id, username, email, password, first_name, last_name, is_active, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	result, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.IsActive, time.Now(), time.Now())
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

// Update updates an existing user in the database.
func (r *userRepositoryImpl) Update(user *entity.User) error {
	// Define the SQL update query
	query := `UPDATE users
			  SET username = $1, email = $2, password = $3, first_name = $4, last_name = $5, is_active = $6, updated_at = $7
			  WHERE id = $8`

	// Execute the update query with the user data
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.IsActive, time.Now(), user.ID)
	if err != nil {
		log.Printf("Error updating user with ID: %v, error: %v", user.ID, err)
		return err
	}

	// Check how many rows were affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	// If no rows were affected, it means the user was not found
	if rowsAffected == 0 {
		log.Printf("No user found with ID: %v", user.ID)
		return fmt.Errorf("user not found")
	}

	log.Printf("Updated user with ID: %v", user.ID)
	return nil
}

// Delete deletes a user from the database.
func (r *userRepositoryImpl) Delete(userID uuid.UUID) error {
	// Define the SQL delete query
	query := `DELETE FROM users WHERE id = $1`

	// Execute the delete query with the user ID
	result, err := r.db.Exec(query, userID)
	if err != nil {
		log.Printf("Error deleting user with ID: %v, error: %v", userID, err)
		return err
	}

	// Check how many rows were affected by the delete
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	// If no rows were affected, it means the user was not found
	if rowsAffected == 0 {
		log.Printf("No user found with ID: %v", userID)
		return fmt.Errorf("user not found")
	}

	log.Printf("Deleted user with ID: %v", userID)
	return nil
}

// FindByID finds a user by their ID.
func (r *userRepositoryImpl) FindByID(userID uuid.UUID) (*entity.User, error) {
	// Define the user entity to store the result
	var user entity.User

	// Fetch the user from the database using the provided ID
	err := r.db.QueryRow(`SELECT id, username, email, password, first_name, last_name, is_active, created_at, updated_at
						 FROM users WHERE id = $1`, userID).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.FirstName, &user.LastName, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	// Check for errors in retrieving the user
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with ID: %v", userID)
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error retrieving user by ID: %v", err)
		return nil, err
	}

	// Return the user if found
	return &user, nil
}

// FindByEmail finds a user by their email.
func (r *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, email, password, first_name, last_name, is_active, created_at, updated_at FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// ListAll lists all users in the database.
func (r *userRepositoryImpl) ListAll() ([]*entity.User, error) {
	panic("unimplemented")
}
