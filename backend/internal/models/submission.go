package models

import (
	"time"

	"gorm.io/gorm"
)

type Submission struct {
	SubmissionID     int       `gorm:"primaryKey;autoIncrement"`
	UserID           string    `gorm:"not null"`
	AssignmentID     string    `gorm:"not null"`
	SubmissionFileID string    `gorm:"not null"`
	SubmittedAt      time.Time `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
