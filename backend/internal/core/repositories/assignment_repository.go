package repositories

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

type AssignmentRepository interface {
	FindAssignmentsForStudent(studentID uuid.UUID) ([]*models.AssignmentWithSubmission, error)
	GetAssignmentByID(assignmentID string) (*models.Assignment, error)
	CreateAssignment(assignment *models.Assignment) error
	UpdateAssignment(assignment *models.Assignment) error
	DeleteAssignment(assignmentID string) error
	// Add other methods as needed
}
