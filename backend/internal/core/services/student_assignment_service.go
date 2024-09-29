package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
)

type AssignmentStudentService struct {
	repo repositories.AssignmentStudentRepository
}

func NewAssignmentStudentService(repo repositories.AssignmentStudentRepository) *AssignmentStudentService {
	return &AssignmentStudentService{repo: repo}
}

func (s *AssignmentStudentService) GetUserAssignments(userID string) ([]models.Assignment, error) {
	return s.repo.GetAssignmentsByUserID(userID)
}

func (s *AssignmentStudentService) GetAssignmentFile(assignmentID string) (string, error) {
	return s.repo.GetAssignmentFile(assignmentID)
}
