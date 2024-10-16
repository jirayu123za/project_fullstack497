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
	RemoveCourse(CourseID uuid.UUID) error

	// CRUD operations for Courses using jwt
	FindCourseByUserID(UserID uuid.UUID) ([]*models.Course, error)
	FindNameByUserID(UserID uuid.UUID) (string, error)
	FindPersonDataByUserID(UserID uuid.UUID) (*models.User, error)
	FindUserGroupByUserID(UserID uuid.UUID) (string, error)
	FindAssignmentByUserID(UserID uuid.UUID) ([]*models.Assignment, error)
	FindAssignmentByUserIDSorted(UserID uuid.UUID) ([]*models.Assignment, error)

	// CRUD operations for Assignments
	AddAssignment(CourseID uuid.UUID, Assignment *models.Assignment) error
	FindAssignmentByAssignmentID(AssignmentID uuid.UUID) (*models.Assignment, error)
	FindAssignments() ([]*models.Assignment, error)
	FindAssignmentsByCourseID(CourseID uuid.UUID) ([]*models.Assignment, error)
	FindAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID) (*models.Assignment, error)
	ModifyAssignment(Assignment *models.Assignment) error
	ModifyAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID, assignment *models.Assignment) error
	RemoveAssignment(AssignmentID uuid.UUID) error
	RemoveAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID) error

	// Using minio
	SaveAssignmentFile(file *models.AssignmentFile) error
	SaveUpload(upload *models.Upload) error

	// R operations for Submissions
	FindSubmissionsByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID) ([]*models.User, error)

	// CRD operations for Instructor lists
	AddInstructorList(CourseID uuid.UUID, InstructorList *models.InstructorList) error  // Create a new instructor list
	FindInstructorsList() ([]*models.InstructorList, error)                             // Find all instructor lists
	FindInstructorsListByCourseID(CourseID uuid.UUID) ([]*models.InstructorList, error) // Find an instructor list by course ID
	FindInstructorsListByListID(ListID uuid.UUID) (*models.InstructorList, error)       // Find an instructor list by list ID
	RemoveInstructorList(InstructorList *models.InstructorList) error                   // Remove an instructor list

	// CRUD operations for Enrollments
	AddEnrollment(CourseID uuid.UUID, Enrollment *models.Enrollment) error
	FindEnrollments() ([]*models.Enrollment, error)
	FindEnrollmentsByCourseID(CourseID uuid.UUID) ([]*models.Enrollment, error)
	RemoveEnrollment(Enrollment *models.Enrollment) error
	FindUsersEnrollment(CourseID uuid.UUID) ([]*models.User, error)
	RemoveUserEnrollment(CourseID uuid.UUID, UserID uuid.UUID) error

	// Delete Enrollments, Assignments, InstructorLists, and Course
	RemoveEnrollmentsByCourseID(CourseID uuid.UUID) error
	RemoveAssignmentsByCourseID(CourseID uuid.UUID) error
	RemoveInstructorListsByCourseID(CourseID uuid.UUID) error
}
