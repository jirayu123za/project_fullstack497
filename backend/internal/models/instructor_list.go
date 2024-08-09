package models

import (
	"time"

	"gorm.io/gorm"
)

type InstructorList struct {
	ListID    int    `gorm:"primaryKey;autoIncrement"`
	UserID    string `gorm:"not null"`
	CourseID  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
