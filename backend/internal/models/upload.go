package models

import (
	"time"

	"gorm.io/gorm"
)

type Upload struct {
	UploadID         string `gorm:"primaryKey;autoIncrement"`
	UserID           int    `gorm:"not null"`
	AssignmentFileID int    `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
