package models

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	EnrollmentID int `gorm:"primaryKey;autoIncrement"`
	UserID       int `gorm:"not null"`
	CourseID     int `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
