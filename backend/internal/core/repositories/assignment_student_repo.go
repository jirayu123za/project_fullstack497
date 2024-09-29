package repositories

import "backend_fullstack/internal/models"

type AssignmentStudentRepository interface {
	GetAssignmentsByUserID(userID string) ([]models.Assignment, error)
	GetAssignmentFile(assignmentID string) (string, error)
}
