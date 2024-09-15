package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assignment struct {
	AssignmentID          string           `gorm:"primaryKey;autoIncrement"`
	CourseID              uuid.UUID        `gorm:"not null"`
	AssignmentName        string           `gorm:"type:varchar(255);not null"`
	AssignmentDescription string           `gorm:"type:varchar(255)"`
	DueDate               time.Time        `gorm:"not null"`
	AssignmentFiles       []AssignmentFile `gorm:"foreignKey:AssignmentID"`
	Submissions           []Submission     `gorm:"foreignKey:AssignmentID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
