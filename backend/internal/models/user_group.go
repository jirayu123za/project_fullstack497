package models

import (
	"time"

	"gorm.io/gorm"
)

type UserGroup struct {
	GroupID   int    `gorm:"primaryKey;autoIncrement"`
	GroupName string `gorm:"type:varchar(50);not null"`
	Users     []User `gorm:"foreignKey:GroupID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
