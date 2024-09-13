package adapters

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Secondary adapters
type GormInstructorRepository struct {
	db *gorm.DB
}

func NewGormInstructorRepository(db *gorm.DB) *GormInstructorRepository {
	return &GormInstructorRepository{
		db: db,
	}
}

func (r *GormInstructorRepository) AddCourse(Course *models.Course) error {
	// Implement the logic to AddCourse to the database using GORM.
	if course := r.db.Create(Course); course.Error != nil {
		return course.Error
	}
	return nil
}

func (r *GormInstructorRepository) FindCourseByID(courseID uuid.UUID) (*models.Course, error) {
	var course *models.Course
	if result := r.db.Preload("Assignments").
		Preload("InstructorLists").
		Preload("Enrollments").
		First(&course, "course_id = ?", courseID); result.Error != nil {
		return nil, result.Error
	}
	return course, nil
}

func (r *GormInstructorRepository) FindCourses() ([]*models.Course, error) {
	var courses []*models.Course
	if result := r.db.Find(&courses); result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (r *GormInstructorRepository) ModifyCourse(Course *models.Course) error {
	var existingCourse *models.Course
	if result := r.db.First(&existingCourse, "course_id = ?", Course.CourseID); result.Error != nil {
		return result.Error
	}

	existingCourse.CourseName = Course.CourseName
	existingCourse.CourseDescription = Course.CourseDescription

	if result := r.db.Save(&existingCourse); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormInstructorRepository) RemoveCourse(Course *models.Course) error {
	var findCourse *models.Course
	if result := r.db.First(&findCourse, "course_id = ?", Course.CourseID); result.Error != nil {
		return result.Error
	}

	if result := r.db.Delete(&findCourse); result.Error != nil {
		return result.Error
	}
	return nil
}

// Under line here be GormInstructorRepository of Instructor list
func (r *GormInstructorRepository) AddInstructorList(CourseID uuid.UUID, InstructorList *models.InstructorList) error {
	// Implement the logic to AddInstructorList to the database using GORM.
	var findCourse *models.Course
	if result := r.db.First(&findCourse, "course_id = ?", CourseID); result.Error != nil {
		return result.Error
	}
	InstructorList.CourseID = CourseID
	if instructorList := r.db.Create(InstructorList); instructorList.Error != nil {
		return instructorList.Error
	}
	return nil
}

func (r *GormInstructorRepository) FindInstructorsList() ([]*models.InstructorList, error) {
	var instructorLists []*models.InstructorList
	if result := r.db.Find(&instructorLists); result.Error != nil {
		return nil, result.Error
	}
	return instructorLists, nil
}

func (r *GormInstructorRepository) FindInstructorsListByCourseID(CourseID uuid.UUID) ([]*models.InstructorList, error) {
	var instructorLists []*models.InstructorList
	if result := r.db.Find(&instructorLists, "course_id = ?", CourseID); result.Error != nil {
		return nil, result.Error
	}
	return instructorLists, nil
}

func (r *GormInstructorRepository) FindInstructorsListByListID(ListID uuid.UUID) (*models.InstructorList, error) {
	var instructorList *models.InstructorList
	if result := r.db.First(&instructorList, "list_id = ?", ListID); result.Error != nil {
		return nil, result.Error
	}
	return instructorList, nil
}

func (r *GormInstructorRepository) RemoveInstructorList(InstructorList *models.InstructorList) error {
	var findInstructorList *models.InstructorList
	if result := r.db.First(&findInstructorList, "list_id = ?", InstructorList.ListID); result.Error != nil {
		return result.Error
	}

	if result := r.db.Delete(&findInstructorList); result.Error != nil {
		return result.Error
	}
	return nil
}
