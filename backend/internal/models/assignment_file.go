package models

import (
	"time"

	"gorm.io/gorm"
)

type AssignmentFile struct {
	AssignmentFileID   int    `gorm:"primaryKey;autoIncrement"`
	AssignmentID       int    `gorm:"not null"`
	AssignmentFileName string `gorm:"type:varchar(255);not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
