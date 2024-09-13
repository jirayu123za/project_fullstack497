package repositories

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

// Secondary ports
type InstructorRepository interface {
	// CRUD operations for Courses
	AddCourse(Course *models.Course) error
	FindCourseByID(CourseID uuid.UUID) (*models.Course, error)
	FindCourses() ([]*models.Course, error)
	ModifyCourse(Course *models.Course) error
	RemoveCourse(Course *models.Course) error

	// CRD operations for Instructor lists
	AddInstructorList(CourseID uuid.UUID, InstructorList *models.InstructorList) error  // Create a new instructor list
	FindInstructorsList() ([]*models.InstructorList, error)                             // Find all instructor lists
	FindInstructorsListByCourseID(CourseID uuid.UUID) ([]*models.InstructorList, error) // Find an instructor list by course ID
	FindInstructorsListByListID(ListID uuid.UUID) (*models.InstructorList, error)       // Find an instructor list by list ID
	RemoveInstructorList(InstructorList *models.InstructorList) error                   // Remove an instructor list
}
