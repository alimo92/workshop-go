package model

import "time"

type User struct {
	// Primary key
	ID uint

	FirstName string
	LastName  string

	// Automatically managed by GORM for creation time and update time
	CreatedAt time.Time
	UpdatedAt time.Time
}
