package model

import "time"

type YaraRule struct {
	// Primary key
	ID uint

	Rule string

	// Automatically managed by GORM for creation time and update time
	CreatedAt time.Time
	UpdatedAt time.Time
}
