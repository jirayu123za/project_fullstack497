package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AssignmentStudentService struct {
	repo     repositories.AssignmentStudentRepository
	fileRepo repositories.FileRepository //new repository for file metadata
}

func NewAssignmentStudentService(repo repositories.AssignmentStudentRepository) *AssignmentStudentService {
	return &AssignmentStudentService{repo: repo}
}

func (s *AssignmentStudentService) GetUserAssignments(userID uuid.UUID) ([]models.Assignment, error) {
	return s.repo.GetAssignmentsByUserID(userID)
}

func (s *AssignmentStudentService) GetAssignmentFile(assignmentID uuid.UUID) (string, error) {
	return s.repo.GetAssignmentFile(assignmentID)
}

func (s *AssignmentStudentService) SubmitAssignment(assignmentID, userID uuid.UUID, filename string) error {
	allowed, _, err := s.repo.IsSubmissionAllowed(assignmentID)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("submission is not allowed at this time")
	}

	// Save assignment file
	assignmentFile := models.AssignmentFile{
		AssignmentFileID:   uuid.New(),
		AssignmentID:       assignmentID,
		AssignmentFileName: filename,
	}
	if err := s.fileRepo.SaveAssignmentFile(assignmentFile); err != nil {
		return err
	}

	submission := models.Submission{
		SubmissionID:     uuid.New(),
		UserID:           userID,
		AssignmentID:     assignmentID,
		SubmissionFileID: assignmentFile.AssignmentFileID,
		SubmittedAt:      time.Now(),
	}

	existingSubmission, _ := s.repo.GetSubmission(assignmentID, userID)
	if existingSubmission.SubmissionID != uuid.Nil {
		submission.SubmissionID = existingSubmission.SubmissionID
		return s.repo.UpdateSubmission(submission)
	}

	return s.repo.SubmitAssignment(submission)
}
func (s *AssignmentStudentService) DeleteSubmission(assignmentID, userID uuid.UUID) error {
	allowed, dueDate, err := s.repo.IsSubmissionAllowed(assignmentID)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("deletion is not allowed after the due date")
	}

	return s.repo.DeleteSubmission(assignmentID, userID)
}
