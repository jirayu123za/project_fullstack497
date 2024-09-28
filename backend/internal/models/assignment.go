package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assignment struct {
	AssignmentID          uuid.UUID        `gorm:"primaryKey"`
	CourseID              uuid.UUID        `gorm:"not null" json:"course_id"`
	AssignmentName        string           `gorm:"type:varchar(255);not null" json:"assignment_name"`
	AssignmentDescription string           `gorm:"type:varchar(255)" json:"assignment_description"`
	DueDate               time.Time        `gorm:"type:date;not null" json:"due_date"`
	AssignmentFiles       []AssignmentFile `gorm:"foreignKey:AssignmentID"`
	Submissions           []Submission     `gorm:"foreignKey:AssignmentID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}

func (assignment *Assignment) BeforeCreate(tx *gorm.DB) (err error) {
	if assignment.AssignmentID == uuid.Nil {
		assignment.AssignmentID = uuid.New()
	}
	return
}

// UnmarshalJSON to handle the date fields in "yyyy-mm-dd" format
func (assignment *Assignment) UnmarshalJSON(data []byte) error {
	type Alias Assignment
	aux := &struct {
		DueDate string `json:"due_date"`
		*Alias
	}{
		Alias: (*Alias)(assignment),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Mapping of date strings to assignment's date fields
	dateFields := []struct {
		dateStr string
		target  *time.Time
	}{
		{aux.DueDate, &assignment.DueDate},
	}

	// Parse dates and assign to respective fields
	for _, field := range dateFields {
		if field.dateStr != "" {
			parsedDate, err := time.Parse("2006-01-02", field.dateStr)
			if err != nil {
				return err
			}
			*field.target = parsedDate
		}
	}

	return nil
}
