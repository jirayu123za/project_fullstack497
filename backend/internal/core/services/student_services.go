package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

type StudentService interface {
	GetCourseByUserIDStd(UserID uuid.UUID) ([]*models.Course, error)
	GetAssignmentByUserIDStd(UserID uuid.UUID) ([]*models.Assignment, error)
	GetAssignmentByUserIDSortedStd(UserID uuid.UUID) ([]*models.Assignment, error)
}

type StudentServiceImpl struct {
	repo repositories.StudentRepository
}

// func instance business logic call
func NewStudentService(repo repositories.StudentRepository) StudentService {
	return &StudentServiceImpl{
		repo: repo,
	}
}

func (s *StudentServiceImpl) GetCourseByUserIDStd(UserID uuid.UUID) ([]*models.Course, error) {
	Courses, err := s.repo.FindCourseByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return Courses, nil
}

func (s *StudentServiceImpl) GetAssignmentByUserIDStd(UserID uuid.UUID) ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindAssignmentByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}

func (s *StudentServiceImpl) GetAssignmentByUserIDSortedStd(UserID uuid.UUID) ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindAssignmentByUserIDSorted(UserID)
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}
