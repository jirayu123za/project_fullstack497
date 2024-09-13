package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
	"time"

	"github.com/google/uuid"
)

type AssignmentService interface {
	GetAssignmentDashboard(studentID uuid.UUID) ([]*models.AssignmentDashboardItem, error)
	GetAssignmentByID(assignmentID string) (*models.Assignment, error)
	CreateAssignment(assignment *models.Assignment) error
	UpdateAssignment(assignment *models.Assignment) error
	DeleteAssignment(assignmentID string) error
}

type assignmentService struct {
	repo repositories.AssignmentRepository
}

func NewAssignmentService(repo repositories.AssignmentRepository) AssignmentService {
	return &assignmentService{repo: repo}
}

func (s *assignmentService) GetAssignmentDashboard(studentID uuid.UUID) ([]*models.AssignmentDashboardItem, error) {
	assignments, err := s.repo.FindAssignmentsForStudent(studentID)
	if err != nil {
		return nil, err
	}

	var dashboardItems []*models.AssignmentDashboardItem
	for _, assignment := range assignments {
		status := "Pending"
		if assignment.SubmissionID != nil {
			status = "Submitted"
		}

		daysUntilDue := int(time.Until(assignment.DueDate).Hours() / 24)

		item := &models.AssignmentDashboardItem{
			AssignmentWithSubmission: *assignment,
			Status:                   status,
			DaysUntilDue:             daysUntilDue,
		}
		dashboardItems = append(dashboardItems, item)
	}

	return dashboardItems, nil
}

func (s *assignmentService) GetAssignmentByID(assignmentID string) (*models.Assignment, error) {
	return s.repo.GetAssignmentByID(assignmentID)
}

func (s *assignmentService) CreateAssignment(assignment *models.Assignment) error {
	return s.repo.CreateAssignment(assignment)
}

func (s *assignmentService) UpdateAssignment(assignment *models.Assignment) error {
	return s.repo.UpdateAssignment(assignment)
}

func (s *assignmentService) DeleteAssignment(assignmentID string) error {
	return s.repo.DeleteAssignment(assignmentID)
}
