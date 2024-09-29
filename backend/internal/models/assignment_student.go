package models

import "time"

type AssignmentStudent struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	FileURL     string    `json:"file_url"`
}
