package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID          uuid.UUID `gorm:"primaryKey" json:"user_id"`
	GroupID         uint      `gorm:"not null" json:"group_id"`
	UserName        string    `gorm:"type:varchar(50);not null" json:"user_name"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password"`
	FirstName       string    `gorm:"type:varchar(50)" json:"first_name"`
	LastName        string    `gorm:"type:varchar(50)" json:"last_name"`
	Email           string    `gorm:"type:varchar(50);not null" json:"email"`
	ProfileImageURL string    `gorm:"type:varchar(255)" json:"profile_image_url"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt   `gorm:"index"`
	Enrollments     []Enrollment     `gorm:"foreignKey:UserID"`
	InstructorLists []InstructorList `gorm:"foreignKey:UserID"`
	Submissions     []Submission     `gorm:"foreignKey:UserID"`
	Uploads         []Upload         `gorm:"foreignKey:UserID"`
	Submitted       bool             `gorm:"-" json:"submitted"`
	Submission      bool             `gorm:"-" json:"submission"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.UserID == uuid.Nil {
		user.UserID = uuid.New()
	}
	return
}
