package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	CourseID          int    `gorm:"primaryKey;autoIncrement"`
	CourseName        string `gorm:"type:varchar(255);not null"`
	CourseDescription string `gorm:"type:varchar(255)"`
	Term              string `gorm:"type:varchar(50)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
