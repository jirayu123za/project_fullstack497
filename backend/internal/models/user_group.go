package models

import (
	"time"

	"gorm.io/gorm"
)

type UserGroup struct {
	GroupID   int    `gorm:"primaryKey;"`
	GroupName string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
