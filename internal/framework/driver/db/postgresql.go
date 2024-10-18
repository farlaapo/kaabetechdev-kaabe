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
    instructor_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Added instructor_id with foreign key constraint
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration VARCHAR(100) NOT NULL,
    version UUID NOT NULL,
    category VARCHAR(100) NOT NULL,
    enrolled_count INT NOT NULL DEFAULT 0,  -- INT with a default value of 0
    content_url TEXT[],  -- Change to TEXT[] for array of URLs
    outline TEXT,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);`

	spaceTable := `CREATE TABLE IF NOT EXISTS spaces (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),     -- Unique identifier for the space
    name VARCHAR(255) NOT NULL,                         -- The name of the space
    description TEXT,                                   -- A brief description of the space
    coach_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to the users table (coach)
    member_count UUID,                                  -- UUID for tracking the number of members (this can be adjusted if needed)
    session_count UUID,                                 -- UUID for tracking the number of sessions (this can be adjusted if needed)
    course_count UUID,                                  -- UUID for tracking the number of courses (this can be adjusted if needed)
    active BOOLEAN DEFAULT TRUE,                        -- Indicates if the space is active or disabled
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Time of creation
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Time of last update
    deleted_at TIMESTAMP NULL                           -- Soft deletion timestamp
);`

	meetingTable := `CREATE TABLE IF NOT EXISTS meetings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration VARCHAR(50),  -- To handle various duration formats like "1h30m"
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    location VARCHAR(255),
    attendee_ids UUID[],  -- UUID arrays are supported in PostgreSQL
    attendee_names TEXT[],
    attendee_emails TEXT[],
    attendee_status TEXT[],
    meeting_type VARCHAR(50),
    status VARCHAR(50),
    join_url TEXT[],
    maximum_capacity INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
`
	paymentTable := `CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    order_id UUID, 
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    transaction_id VARCHAR(100) UNIQUE,
    status VARCHAR(20) NOT NULL,
    payment_gateway VARCHAR(50),
    payment_date TIMESTAMP,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

	// Create tokens table
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
	queries := []string{userTable, tokenTable, roleTable, permissionTable, userRoleTable, userPermissionTable, courseTable, spaceTable, meetingTable, paymentTable}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}

	log.Println("Successfully created all tables")
	return nil
}
