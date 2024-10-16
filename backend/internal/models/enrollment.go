package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	EnrollmentID uuid.UUID `gorm:"primaryKey"`
	UserID       uuid.UUID `gorm:"not null" json:"user_id"`
	CourseID     uuid.UUID `gorm:"not null" json:"course_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (enrollment *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if enrollment.EnrollmentID == uuid.Nil {
		enrollment.EnrollmentID = uuid.New()
	}
	return
}
