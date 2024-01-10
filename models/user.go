package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint         `gorm:"primaryKey"` // Standard field for the primary key
	Name        string       // A regular string field
	Email       *string      // A pointer to a string, allowing for null values
	Birthday    *time.Time   // A pointer to time.Time, can be null
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time    // Automatically managed by GORM for update time
}
