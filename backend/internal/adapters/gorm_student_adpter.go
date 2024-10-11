package adapters

import (
	"backend_fullstack/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Secondary adapters
type GormStudentRepository struct {
	db *gorm.DB
}

func NewGormStudentRepository(db *gorm.DB) *GormStudentRepository {
	return &GormStudentRepository{
		db: db,
	}
}

// Using jwt
func (r *GormStudentRepository) FindCourseByUserID(UserID uuid.UUID) ([]*models.Course, error) {
	var courses []*models.Course
	if err := r.db.
		Joins("JOIN enrollments ON enrollments.course_id = courses.course_id").
		Where("enrollments.user_id = ?", UserID).
		Preload("Assignments").
		Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *GormStudentRepository) FindAssignmentByUserID(UserID uuid.UUID) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if err := r.db.
		Joins("JOIN courses ON courses.course_id = assignments.course_id").
		Joins("JOIN enrollments ON enrollments.course_id = courses.course_id").
		Where("enrollments.user_id = ?", UserID).
		Preload("AssignmentFiles").
		Preload("Submissions").
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *GormStudentRepository) FindAssignmentByUserIDSorted(UserID uuid.UUID) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if err := r.db.
		Joins("JOIN courses ON courses.course_id = assignments.course_id").
		Joins("JOIN enrollments ON enrollments.course_id = courses.course_id").
		Where("enrollments.user_id = ?", UserID).
		Where("due_date > ?", time.Now()).
		Order("due_date ASC").
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}
