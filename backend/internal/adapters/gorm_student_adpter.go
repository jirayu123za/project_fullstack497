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
		Where("enrollments.user_id = ? AND enrollments.deleted_at IS NULL AND courses.deleted_at IS NULL", UserID).
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
		Where("enrollments.user_id = ? AND enrollments.deleted_at IS NULL AND courses.deleted_at IS NULL", UserID).
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
		Where("enrollments.user_id = ? AND enrollments.deleted_at IS NULL AND courses.deleted_at IS NULL", UserID).
		Where("due_date > ?", time.Now()).
		Order("due_date ASC").
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *GormStudentRepository) FindUpcomingAssignments(UserID uuid.UUID, CourseID uuid.UUID) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if err := r.db.
		Joins("JOIN courses ON courses.course_id = assignments.course_id").
		Joins("JOIN enrollments ON enrollments.course_id = courses.course_id").
		Where("enrollments.user_id = ? AND enrollments.deleted_at IS NULL", UserID).
		Where("courses.course_id = ? AND courses.deleted_at IS NULL", CourseID).
		Where("due_date > ?", time.Now()).
		Order("due_date ASC").
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *GormStudentRepository) SaveSubmission(submission *models.Submission) error {
	if result := r.db.Create(submission); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormStudentRepository) FindSubmissionsStatus(CourseID uuid.UUID, AssignmentID uuid.UUID, UserID uuid.UUID) ([]*models.User, error) {
	var users []*models.User

	// if err := r.db.
	// 	Table("enrollments").
	// 	Select("DISTINCT ON (users.user_id) users.user_id, users.first_name, users.last_name, CASE WHEN submissions.submitted_at IS NOT NULL THEN true ELSE false END AS submission").
	// 	Joins("JOIN users ON enrollments.user_id = users.user_id").
	// 	Joins("LEFT JOIN submissions ON enrollments.user_id = submissions.user_id AND submissions.assignment_id = ? AND submissions.deleted_at IS NULL", AssignmentID).
	// 	Where("enrollments.course_id = ? AND enrollments.user_id = ? AND enrollments.deleted_at IS NULL", CourseID, UserID).
	// 	Scan(&user).Error; err != nil {
	// 	return nil, err
	// }
	if err := r.db.
		Table("enrollments").
		Select("users.user_id, users.first_name, users.last_name, COUNT(submissions.submitted_at) > 0 AS submission").
		Joins("JOIN users ON enrollments.user_id = users.user_id").
		Joins("LEFT JOIN submissions ON enrollments.user_id = submissions.user_id AND submissions.assignment_id = ? AND submissions.deleted_at IS NULL", AssignmentID).
		Where("enrollments.course_id = ? AND enrollments.user_id = ? AND enrollments.deleted_at IS NULL", CourseID, UserID).
		Group("users.user_id, users.first_name, users.last_name").
		Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
