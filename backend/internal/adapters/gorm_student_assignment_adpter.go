package adapters

import (
	"backend_fullstack/internal/models"

	"gorm.io/gorm"
)

type GormAssignmentStudentRepository struct {
	db *gorm.DB
}

func NewGormAssignmentRepository(db *gorm.DB) *GormAssignmentStudentRepository {
	return &GormAssignmentStudentRepository{db: db}
}

func (r *GormAssignmentStudentRepository) GetAssignmentsByUserID(userID string) ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.db.Table("assignments").
		Joins("JOIN enrollments ON assignments.course_id = enrollments.course_id").
		Where("enrollments.user_id = ?", userID).
		Find(&assignments).Error
	return assignments, err
}

func (r *GormAssignmentStudentRepository) GetAssignmentFile(assignmentID string) (string, error) {
	var assignment models.AssignmentStudent
	err := r.db.Where("id = ?", assignmentID).First(&assignment).Error
	return assignment.FileURL, err
}
