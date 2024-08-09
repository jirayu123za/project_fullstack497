package models

import (
	"time"

	"gorm.io/gorm"
)

type InstructorList struct {
	ListID    int `gorm:"primaryKey;autoIncrement"`
	UserID    int `gorm:"not null"`
	CourseID  int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
