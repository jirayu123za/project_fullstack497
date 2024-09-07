package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstructorList struct {
	ListID    uuid.UUID `gorm:"primaryKey" json:"list_id"`
	UserID    uuid.UUID `gorm:"not null" json:"user_id"`
	CourseID  uuid.UUID `gorm:"not null" json:"course_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
