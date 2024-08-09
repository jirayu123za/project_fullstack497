package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"primaryKey;autoIncrement"`
	GroupID   int    `gorm:"not null"`
	UserName  string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	FirstName string `gorm:"type:varchar(50)"`
	LastName  string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
