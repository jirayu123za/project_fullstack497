package models

import (
	"time"

	"gorm.io/gorm"
)

type Upload struct {
	UploadID         string `gorm:"primaryKey;autoIncrement"`
	UserID           string `gorm:"not null"`
	AssignmentFileID string `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
