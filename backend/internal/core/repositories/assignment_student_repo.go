package repositories

import (
	"backend_fullstack/internal/models"
	"time"

	"github.com/google/uuid"
)

type AssignmentStudentRepository interface {
	GetAssignmentsByUserID(userID uuid.UUID) ([]models.Assignment, error)
	GetAssignmentFile(assignmentID uuid.UUID) (string, error)
	SubmitAssignment(submission models.Submission) error
	GetSubmission(assignmentID, userID uuid.UUID) (models.Submission, error)
	UpdateSubmission(submission models.Submission) error
	DeleteSubmission(assignmentID, userID uuid.UUID) error
	IsSubmissionAllowed(assignmentID uuid.UUID) (bool, time.Time, error)
}
