package adapters

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormAssignmentRepository struct {
	db *gorm.DB
}

func NewGormAssignmentRepository(db *gorm.DB) repositories.AssignmentRepository {
	return &GormAssignmentRepository{db: db}
}

func (r *GormAssignmentRepository) FindAssignmentsForStudent(studentID uuid.UUID) ([]*models.AssignmentWithSubmission, error) {
	var assignmentsWithSubmissions []*models.AssignmentWithSubmission

	err := r.db.Table("assignments").
		Select("assignments.*, submissions.submission_id, submissions.submitted_at").
		Joins("JOIN enrollments ON enrollments.course_id = assignments.course_id").
		Joins("LEFT JOIN submissions ON submissions.assignment_id = assignments.assignment_id AND submissions.user_id = enrollments.user_id").
		Where("enrollments.user_id = ?", studentID).
		Scan(&assignmentsWithSubmissions).Error

	return assignmentsWithSubmissions, err
}

func (r *GormAssignmentRepository) GetAssignmentByID(assignmentID string) (*models.Assignment, error) {
	var assignment models.Assignment
	err := r.db.First(&assignment, "assignment_id = ?", assignmentID).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *GormAssignmentRepository) CreateAssignment(assignment *models.Assignment) error {
	return r.db.Create(assignment).Error
}

func (r *GormAssignmentRepository) UpdateAssignment(assignment *models.Assignment) error {
	return r.db.Save(assignment).Error
}

func (r *GormAssignmentRepository) DeleteAssignment(assignmentID string) error {
	return r.db.Delete(&models.Assignment{}, "assignment_id = ?", assignmentID).Error
}
