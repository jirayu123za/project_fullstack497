package models

import (
	"time"

	"gorm.io/gorm"
)

type Submission struct {
	SubmissionID     int       `gorm:"primaryKey;autoIncrement"`
	UserID           int       `gorm:"not null"`
	AssignmentID     int       `gorm:"not null"`
	SubmissionFileID int       `gorm:"not null"`
	SubmittedAt      time.Time `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
