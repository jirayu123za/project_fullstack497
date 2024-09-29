package adapters

import (
	"backend_fullstack/internal/models"

	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormAssignmentStudentRepository struct {
	db *gorm.DB
}

func NewGormAssignmentRepository(db *gorm.DB) *GormAssignmentStudentRepository {
	return &GormAssignmentStudentRepository{db: db}
}

func (r *GormAssignmentStudentRepository) GetAssignmentsByUserID(userID uuid.UUID) ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.db.Joins("JOIN enrollments ON assignments.course_id = enrollments.course_id").
		Where("enrollments.user_id = ?", userID).
		Find(&assignments).Error
	return assignments, err
}

func (r *GormAssignmentStudentRepository) GetAssignmentFile(assignmentID uuid.UUID) (string, error) {
	var assignment models.AssignmentStudent
	err := r.db.First(&assignment, assignmentID).Error
	if err != nil {
		return "", err
	}
	return assignment.FileURL, nil
}
func (r *GormAssignmentStudentRepository) SubmitAssignment(submission models.Submission) error {
	return r.db.Create(&submission).Error
}
func (r *GormAssignmentStudentRepository) GetSubmission(assignmentID, userID uuid.UUID) (models.Submission, error) {
	var submission models.Submission
	err := r.db.Where("assignment_id = ? AND user_id = ?", assignmentID, userID).First(&submission).Error
	return submission, err
}
func (r *GormAssignmentStudentRepository) UpdateSubmission(submission models.Submission) error {
	return r.db.Save(&submission).Error
}
func (r *GormAssignmentStudentRepository) DeleteSubmission(assignmentID, userID uuid.UUID) error {
	return r.db.Where("assignment_id = ? AND user_id = ?", assignmentID, userID).Delete(&models.Submission{}).Error
}
func (r *GormAssignmentStudentRepository) IsSubmissionAllowed(assignmentID uuid.UUID) (bool, time.Time, error) {
	var assignment models.Assignment
	err := r.db.First(&assignment, assignmentID).Error
	if err != nil {
		return false, time.Time{}, err
	}

	now := time.Now()
	if now.After(assignment.DueDate) {
		return false, assignment.DueDate, errors.New("submission deadline has passed")
	}

	return true, assignment.DueDate, nil
}
