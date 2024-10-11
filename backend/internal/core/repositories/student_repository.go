package repositories

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

// Secondary ports
type StudentRepository interface {
	FindCourseByUserID(UserID uuid.UUID) ([]*models.Course, error)
	FindAssignmentByUserID(UserID uuid.UUID) ([]*models.Assignment, error)
	FindAssignmentByUserIDSorted(UserID uuid.UUID) ([]*models.Assignment, error)
	FindUpcomingAssignments(UserID uuid.UUID, CourseID uuid.UUID) ([]*models.Assignment, error)

	// Using minio
	SaveSubmission(submission *models.Submission) error
}
