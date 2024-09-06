package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"errors"
	"log"
	"time"
)

// tokenRepositoryImpl is the implementation of TokenRepository.
type tokenRepositoryImpl struct {
	db *sql.DB
}

// NewTokenRepository creates a new instance of TokenRepository.
func NewTokenRepository(db *sql.DB) repository.TokenRepository {
	return &tokenRepositoryImpl{db: db}
}

// FindByToken retrieves a token by its value.
func (r *tokenRepositoryImpl) FindByToken(token string) (*entity.Token, error) {
	t := &entity.Token{}
	query := `SELECT id, user_id, token, expires_at, created_at, updated_at FROM tokens WHERE token = $1`
	row := r.db.QueryRow(query, token)

	err := row.Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("token not found")
		}
		return nil, err
	}

	return t, nil
}

// Create inserts a new token into the database.
func (r *tokenRepositoryImpl) Create(token *entity.Token) error {
	query := `INSERT INTO tokens (id, user_id, token, expires_at, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6)`
	result, err := r.db.Exec(query, token.ID, token.UserID, token.Token, token.ExpiresAt, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error inserting token: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	log.Printf("Rows affected: %d", rowsAffected)
	return nil
}
