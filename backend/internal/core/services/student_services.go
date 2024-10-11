package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type StudentService interface {
	GetCourseByUserIDStd(UserID uuid.UUID) ([]*models.Course, error)
	GetAssignmentByUserIDStd(UserID uuid.UUID) ([]*models.Assignment, error)
	GetAssignmentByUserIDSortedStd(UserID uuid.UUID) ([]*models.Assignment, error)
	GetUpcomingAssignments(UserID uuid.UUID, CourseID uuid.UUID) ([]*models.Assignment, error)

	// Using minio
	CreateSubmission(submission *models.Submission) error
	CreateAssignmentFiles(userID, assignmentID uuid.UUID, userGroupName, userName string, files []*multipart.FileHeader) ([]uuid.UUID, error)
}

type StudentServiceImpl struct {
	repo      repositories.StudentRepository
	minioRepo repositories.MinIORepository
}

// func instance business logic call
func NewStudentService(repo repositories.StudentRepository, minioRepo repositories.MinIORepository) StudentService {
	return &StudentServiceImpl{
		repo:      repo,
		minioRepo: minioRepo,
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

func (s *StudentServiceImpl) GetUpcomingAssignments(UserID uuid.UUID, CourseID uuid.UUID) ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindUpcomingAssignments(UserID, CourseID)
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}

func (s *StudentServiceImpl) CreateSubmission(submission *models.Submission) error {
	if err := s.repo.SaveSubmission(submission); err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) CreateAssignmentFiles(userID, assignmentID uuid.UUID, userGroupName, userName string, files []*multipart.FileHeader) ([]uuid.UUID, error) {
	var fileIDs []uuid.UUID

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}

		fileExtension := filepath.Ext(fileHeader.Filename)
		newFileName := uuid.New().String() + fileExtension

		submission := models.Submission{
			UserID:           userID,
			AssignmentID:     assignmentID,
			SubmissionFileID: newFileName,
			SubmittedAt:      time.Now(),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}

		if err := s.repo.SaveSubmission(&submission); err != nil {
			file.Close()
			return nil, err
		}

		if err := s.minioRepo.SaveFileToMinIO(file, userGroupName, userName, newFileName); err != nil {
			file.Close()
			return nil, err
		}

		file.Close()
	}

	return fileIDs, nil
}
