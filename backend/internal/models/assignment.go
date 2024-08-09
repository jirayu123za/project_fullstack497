package models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	AssignmentID          int       `gorm:"primaryKey;autoIncrement"`
	CourseID              int       `gorm:"not null"`
	AssignmentName        string    `gorm:"type:varchar(255);not null"`
	AssignmentDescription string    `gorm:"type:varchar(255)"`
	DueDate               time.Time `gorm:"not null"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
