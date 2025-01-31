package store

import (
	"database/sql"
	"fmt"
)

type Store struct {
	DB *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

type Connection struct {
	ID        int
	Timestamp string
	SSID      string
	Manual    bool
	Type      string // Type of entry e.g. sick, annual ...
}

func (s *Store) GetWorkDayEntries() (int, error) {
	// Query that counts unique days from this month
	query := `
        SELECT COUNT(DISTINCT DATE(timestamp)) as unique_days
        FROM connections 
        WHERE strftime('%Y-%m', timestamp) = strftime('%Y-%m', 'now');
    `

	var uniqueDays int
	err := s.DB.QueryRow(query).Scan(&uniqueDays)
	if err != nil {
		return 0, fmt.Errorf("failed to query unique days: %w", err)
	}

	return uniqueDays, nil
}
