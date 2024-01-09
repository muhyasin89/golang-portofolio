package models

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
