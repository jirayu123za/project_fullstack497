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

type SubmissionFile struct {
	SubmissionFileID string         `gorm:"primaryKey;autoIncrement" json:"submission_file_id"`
	SubmissionID     string         `gorm:"not null" json:"submission_id"`
	FileName         string         `gorm:"type:varchar(255);not null" json:"file_name"`
	FileURL          string         `gorm:"type:varchar(255);not null" json:"file_url"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type AssignmentWithSubmission struct {
	Assignment
	SubmissionID *string    `json:"submission_id,omitempty"`
	SubmittedAt  *time.Time `json:"submitted_at,omitempty"`
}

type AssignmentDashboardItem struct {
	AssignmentWithSubmission
	Status       string `json:"status"`
	DaysUntilDue int    `json:"days_until_due"`
}
