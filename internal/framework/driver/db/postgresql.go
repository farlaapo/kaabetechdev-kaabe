package db

import (
	"dalabio/pkg/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB initializes and returns a PostgreSQL database connection.
func ConnectDB(cfg *config.DBConfig) (*sql.DB, error) {
	connStr := cfg.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Ping the database to ensure the connection is successful
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return db, nil
}

// CreateTables ensures that the required tables are created.
func CreateTables(db *sql.DB) error {
	userTable := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		is_active BOOLEAN DEFAULT TRUE,
		last_login TIMESTAMP,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL
	);`

	courseTable := `CREATE TABLE IF NOT EXISTS courses (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	duration VARCHAR(100) NOT NULL,
	version UUID NOT NULL,
	category VARCHAR(100) NOT NULL,
	enrolled_count INT NOT NULL DEFAULT 0,  -- Changed to INT with a default value of 0
	content_url TEXT,
	outline TEXT,
	status VARCHAR(50) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);`

	tokenTable := `CREATE TABLE IF NOT EXISTS tokens (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		token VARCHAR(255) UNIQUE NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL
	);`

	// Create roles table
	roleTable := `CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL
	);`

	// Create permissions table
	permissionTable := `CREATE TABLE IF NOT EXISTS permissions (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL
	);`

	// Create user_roles table
	userRoleTable := `CREATE TABLE IF NOT EXISTS user_roles (
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		role_id INT REFERENCES roles(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, role_id)
	);`

	// Create user_permissions table
	userPermissionTable := `CREATE TABLE IF NOT EXISTS user_permissions (
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, permission_id)
	);`

	// Execute the table creation queries
	queries := []string{userTable, tokenTable, roleTable, permissionTable, userRoleTable, userPermissionTable, courseTable}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}

	log.Println("Successfully created all tables")
	return nil
}
