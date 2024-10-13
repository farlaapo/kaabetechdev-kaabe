package gateway

import (
	"dalabio/internal/entity"
	"dalabio/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

type MeetingRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.MeetingRepository.
func (r *MeetingRepositoryImpl) Create(meeting *entity.Meeting) error {
	query := ` INSERT INTO meetings (id, title, description, duration, start_time, end_time, location,  attendee_ids, attendee_names, attendee_emails, attendee_status, meeting_type, status, join_url, maximum_capacity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`
	result, err := r.db.Exec(query, meeting.ID, meeting.Title, meeting.Description, meeting.Duration, meeting.StartTime, meeting.EndTime, meeting.Location, pq.Array(meeting.AttendeeIDs), pq.Array(meeting.AttendeeNames), pq.Array(meeting.AttendeeEmails), pq.Array(meeting.AttendeeStatus), meeting.MeetingType, meeting.Status, pq.Array(meeting.JoinURL), meeting.MaximumCapacity, meeting.CreatedAt, meeting.UpdatedAt)

	if err != nil {
		log.Printf("Error inserting course: %v, query: %s", err, query)
		return err
	}

	rowsaffected, err := result.RowsAffected()

	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
	}

	log.Printf("Rows affected: %d\n", rowsaffected)

	return nil

}

// Update implements repository.MeetingRepository.

func (r *MeetingRepositoryImpl) Update(meeting *entity.Meeting) error {
	// Define the SQL update query
	query :=
		`UPDATE meetings
    SET title = $2, 
        description = $3, 
        duration = $4, 
        start_time = $5, 
        end_time = $6, 
        location = $7, 
        attendee_ids = $8, 
        attendee_names = $9, 
        attendee_emails = $10, 
        attendee_status = $11, 
        meeting_type = $12, 
        status = $13, 
        join_url = $14, 
        maximum_capacity = $15, 
        updated_at = CURRENT_TIMESTAMP
    WHERE id = $1;`

	// Concatenate JoinURL into a single string if necessary
	joinURL := strings.Join(meeting.JoinURL, ",") // Adjust this as needed

	// Execute the SQL update query
	result, err := r.db.Exec(query,
		meeting.ID,                       // $1
		meeting.Title,                    // $2
		meeting.Description,              // $3
		meeting.Duration,                 // $4
		meeting.StartTime,                // $5
		meeting.EndTime,                  // $6
		meeting.Location,                 // $7
		pq.Array(meeting.AttendeeIDs),    // $8
		pq.Array(meeting.AttendeeNames),  // $9
		pq.Array(meeting.AttendeeEmails), // $10
		pq.Array(meeting.AttendeeStatus), // $11
		meeting.MeetingType,              // $12
		meeting.Status,                   // $13
		joinURL,                          // $14
		meeting.MaximumCapacity,          // $15
	)

	if err != nil {
		log.Printf("Error updating meeting with ID: %v, error: %v", meeting.ID, err)
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	// If no rows were affected, it means the meeting was not found
	if rowsAffected == 0 {
		log.Printf("No meeting found with ID: %v", meeting.ID)
		return nil
	}

	log.Printf("Meeting Updated: %d\n", meeting.ID)
	return nil
}

// Get implements repository.MeetingRepository.
func (r *MeetingRepositoryImpl) GetdByID(meetingID uuid.UUID) (*entity.Meeting, error) {

	var meeting entity.Meeting
	// SQL Query to insert the course into the database
	query := `SELECT id, title, description, duration, start_time, end_time, location,  attendee_ids, attendee_names, attendee_emails, attendee_status, meeting_type, status, join_url, maximum_capacity, created_at, updated_at  FROM meetings WHERE id = $1`

	err := r.db.QueryRow(query, meetingID).Scan(
		&meeting.ID,
		&meeting.Title,
		&meeting.Description,
		&meeting.Duration,
		&meeting.StartTime,
		&meeting.EndTime,
		&meeting.Location,

		pq.Array(&meeting.AttendeeIDs),
		pq.Array(&meeting.AttendeeNames),
		pq.Array(&meeting.AttendeeEmails),
		pq.Array(&meeting.AttendeeStatus),
		&meeting.MeetingType,
		&meeting.Status,
		pq.Array(&meeting.JoinURL),
		&meeting.MaximumCapacity,
		&meeting.CreatedAt,
		&meeting.UpdatedAt,
	)

	// If no rows were returned, it means the course was not found
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No course found with ID: %v", meetingID)
			return nil, fmt.Errorf("course Not Found")
		}
		log.Printf("Error retrieving course by ID: %v", err)
	}

	return &meeting, nil
}

// GetAll implements repository.MeetingRepository.
func (r *MeetingRepositoryImpl) GetAll() ([]*entity.Meeting, error) {
	var meetings []*entity.Meeting

	query := `SELECT id, title, description, duration, start_time, end_time, location, attendee_ids, attendee_names, attendee_emails, attendee_status, meeting_type, status, join_url, maximum_capacity, created_at, updated_at FROM meetings`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error retrieving meetings: %v", err)
		return nil, err // Return nil on error
	}
	defer rows.Close()

	for rows.Next() {
		var meeting entity.Meeting
		err := rows.Scan(
			&meeting.ID,
			&meeting.Title,
			&meeting.Description,
			&meeting.Duration,
			&meeting.StartTime,
			&meeting.EndTime,
			&meeting.Location,
			pq.Array(&meeting.AttendeeIDs),
			pq.Array(&meeting.AttendeeNames),
			pq.Array(&meeting.AttendeeEmails),
			pq.Array(&meeting.AttendeeStatus),
			&meeting.MeetingType,
			&meeting.Status,
			pq.Array(&meeting.JoinURL),
			&meeting.MaximumCapacity,
			&meeting.CreatedAt,
			&meeting.UpdatedAt,
		)

		if err != nil {
			log.Printf("Error scanning meeting: %v", err)
			return nil, err // Return nil on error
		}
		meetings = append(meetings, &meeting)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating meetings: %v", err)
		return nil, err
	}

	return meetings, nil
}

// Delete implements repository.MeetingRepository.
func (r *MeetingRepositoryImpl) Delete(meetingID uuid.UUID) error {

	// SQL Query to delete the course from the database
	query := `DELETE FROM meetings WHERE id = $1`

	result, err := r.db.Exec(query, meetingID)
	if err != nil {
		log.Printf("Error deleting course with ID: %v, error: %v", meetingID, err)
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
		log.Printf("No course found with ID: %v", meetingID)
		return nil
	}

	return nil

}

func NewMeetingRepository(db *sql.DB) repository.MeetingRepository {
	return &MeetingRepositoryImpl{db: db}

}
