package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

// Primary port
type InstructorService interface {
	// CRUD operations for Courses
	CreateCourse(Course *models.Course) error
	GetCourseByID(CourseID uuid.UUID) (*models.Course, error)
	GetCourses() ([]*models.Course, error)
	UpdateCourse(Course *models.Course) error
	DeleteCourse(CourseID uuid.UUID) error

	// using jwt
	GetCourseByUserID(UserID uuid.UUID) ([]*models.Course, error)
	GetNameByUserID(UserID uuid.UUID) (string, error)
	GetPersonDataByUserID(UserID uuid.UUID) (*models.User, error)
	GetUserGroupByUserID(UserID uuid.UUID) (string, error)
	GetAssignmentByUserID(UserID uuid.UUID) ([]*models.Assignment, error)
	GetAssignmentByUserIDSorted(UserID uuid.UUID) ([]*models.Assignment, error)

	// CRUD operations for Assignments
	CreateAssignment(CourseID uuid.UUID, Assignment *models.Assignment) error
	GetAssignmentByAssignmentID(AssignmentID uuid.UUID) (*models.Assignment, error)
	GetAssignments() ([]*models.Assignment, error)
	GetAssignmentsByCourseID(CourseID uuid.UUID) ([]*models.Assignment, error)
	GetAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID) (*models.Assignment, error)
	// v1 using assignment_id
	UpdateAssignment(Assignment *models.Assignment) error
	UpdateAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID, Assignment *models.Assignment) error
	DeleteAssignment(AssignmentID uuid.UUID) error

	// CRD operations for Instructor lists
	CreateInstructorList(CourseID uuid.UUID, InstructorList *models.InstructorList) error
	GetInstructorsList() ([]*models.InstructorList, error)
	GetInstructorsListByCourseID(CourseID uuid.UUID) ([]*models.InstructorList, error)
	GetInstructorsListByListID(ListID uuid.UUID) (*models.InstructorList, error)
	DeleteInstructorList(InstructorList *models.InstructorList) error

	// CRUD operations for Enrollments
	CreateEnrollment(CourseID uuid.UUID, Enrollment *models.Enrollment) error
	GetEnrollments() ([]*models.Enrollment, error)
	GetEnrollmentsByCourseID(CourseID uuid.UUID) ([]*models.Enrollment, error)
	DeleteEnrollment(Enrollment *models.Enrollment) error
	GetUsersEnrollment(CourseID uuid.UUID) ([]*models.User, error)
	DeleteUserEnrollment(CourseID uuid.UUID, UserID uuid.UUID) error

	// Delete Enrollments, Assignments, InstructorLists, and Course
	DeleteEnrollmentsByCourseID(CourseID uuid.UUID) error
	DeleteAssignmentsByCourseID(CourseID uuid.UUID) error
	DeleteInstructorListsByCourseID(CourseID uuid.UUID) error
}

type InstructorServiceImpl struct {
	repo repositories.InstructorRepository
}

// func instance business logic call
func NewInstructorService(repo repositories.InstructorRepository) InstructorService {
	return &InstructorServiceImpl{
		repo: repo,
	}
}

func (s *InstructorServiceImpl) CreateCourse(Course *models.Course) error {
	if err := s.repo.AddCourse(Course); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) GetCourseByID(CourseID uuid.UUID) (*models.Course, error) {
	Course, err := s.repo.FindCourseByID(CourseID)
	if err != nil {
		return nil, err
	}
	return Course, nil
}

func (s *InstructorServiceImpl) GetCourses() ([]*models.Course, error) {
	Courses, err := s.repo.FindCourses()
	if err != nil {
		return nil, err
	}
	return Courses, nil
}

func (s *InstructorServiceImpl) UpdateCourse(Course *models.Course) error {
	existingCourses, err := s.repo.FindCourseByID(Course.CourseID)
	if err != nil {
		return err
	}

	existingCourses.CourseName = Course.CourseName
	//existingCourses.CourseDescription = Course.CourseDescription

	if err := s.repo.ModifyCourse(existingCourses); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) DeleteCourse(CourseID uuid.UUID) error {
	// 1. Remove Enrollments associated with the course
	if err := s.repo.RemoveEnrollmentsByCourseID(CourseID); err != nil {
		return err
	}

	// 2. Remove Assignments associated with the course
	if err := s.repo.RemoveAssignmentsByCourseID(CourseID); err != nil {
		return err
	}

	// 3. Remove Instructor Lists associated with the course
	if err := s.repo.RemoveInstructorListsByCourseID(CourseID); err != nil {
		return err
	}

	// 4. Finally, remove the course itself
	if err := s.repo.RemoveCourse(CourseID); err != nil {
		return err
	}
	return nil
}

// using JWT
func (s *InstructorServiceImpl) GetCourseByUserID(UserID uuid.UUID) ([]*models.Course, error) {
	Courses, err := s.repo.FindCourseByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return Courses, nil
}

func (s *InstructorServiceImpl) GetNameByUserID(UserID uuid.UUID) (string, error) {
	Name, err := s.repo.FindNameByUserID(UserID)
	if err != nil {
		return "", err
	}
	return Name, nil
}

func (s *InstructorServiceImpl) GetPersonDataByUserID(UserID uuid.UUID) (*models.User, error) {
	PersonData, err := s.repo.FindPersonDataByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return PersonData, nil
}

func (s *InstructorServiceImpl) GetUserGroupByUserID(UserID uuid.UUID) (string, error) {
	UserGroup, err := s.repo.FindUserGroupByUserID(UserID)
	if err != nil {
		return "", err
	}
	return UserGroup, nil
}

func (s *InstructorServiceImpl) GetAssignmentByUserID(UserID uuid.UUID) ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindAssignmentByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}

func (s *InstructorServiceImpl) GetAssignmentByUserIDSorted(UserID uuid.UUID) ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindAssignmentByUserIDSorted(UserID)
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}

// Under line here be InstructorServiceImpl of Instructor assignment
func (s *InstructorServiceImpl) CreateAssignment(CourseID uuid.UUID, Assignment *models.Assignment) error {
	existingCourse, err := s.repo.FindCourseByID(CourseID)
	if err != nil {
		return err
	}

	if err := s.repo.AddAssignment(existingCourse.CourseID, Assignment); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) GetAssignmentByAssignmentID(AssignmentID uuid.UUID) (*models.Assignment, error) {
	assignment, err := s.repo.FindAssignmentByAssignmentID(AssignmentID)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

func (s *InstructorServiceImpl) GetAssignments() ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindAssignments()
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}

func (s *InstructorServiceImpl) GetAssignmentsByCourseID(CourseID uuid.UUID) ([]*models.Assignment, error) {
	Assignments, err := s.repo.FindAssignmentsByCourseID(CourseID)
	if err != nil {
		return nil, err
	}
	return Assignments, nil
}

func (s *InstructorServiceImpl) GetAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID) (*models.Assignment, error) {
	assignment, err := s.repo.FindAssignmentByCourseIDAndAssignmentID(CourseID, AssignmentID)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// v1 using assignment_id
func (s *InstructorServiceImpl) UpdateAssignment(Assignment *models.Assignment) error {
	existingAssignment, err := s.repo.FindAssignmentByAssignmentID(Assignment.AssignmentID)
	if err != nil {
		return err
	}

	existingAssignment.AssignmentName = Assignment.AssignmentName
	existingAssignment.AssignmentDescription = Assignment.AssignmentDescription
	existingAssignment.DueDate = Assignment.DueDate

	if err := s.repo.ModifyAssignment(existingAssignment); err != nil {
		return err
	}

	return nil
}

func (s *InstructorServiceImpl) UpdateAssignmentByCourseIDAndAssignmentID(CourseID uuid.UUID, AssignmentID uuid.UUID, Assignment *models.Assignment) error {
	existingAssignment, err := s.repo.FindAssignmentByCourseIDAndAssignmentID(CourseID, AssignmentID)
	if err != nil {
		return err
	}

	existingAssignment.AssignmentName = Assignment.AssignmentName
	existingAssignment.AssignmentDescription = Assignment.AssignmentDescription
	existingAssignment.DueDate = Assignment.DueDate

	if err := s.repo.ModifyAssignmentByCourseIDAndAssignmentID(CourseID, AssignmentID, existingAssignment); err != nil {
		return err
	}

	return nil
}

func (s *InstructorServiceImpl) DeleteAssignment(AssignmentID uuid.UUID) error {
	existingAssignment, err := s.repo.FindAssignmentByAssignmentID(AssignmentID)
	if err != nil {
		return err
	}

	if err := s.repo.RemoveAssignment(existingAssignment.AssignmentID); err != nil {
		return err
	}
	return nil
}

// Under line here be InstructorServiceImpl of Instructor list
func (s *InstructorServiceImpl) CreateInstructorList(CourseID uuid.UUID, InstructorList *models.InstructorList) error {
	if err := s.repo.AddInstructorList(CourseID, InstructorList); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) GetInstructorsList() ([]*models.InstructorList, error) {
	InstructorLists, err := s.repo.FindInstructorsList()
	if err != nil {
		return nil, err
	}
	return InstructorLists, nil
}

func (s *InstructorServiceImpl) GetInstructorsListByCourseID(CourseID uuid.UUID) ([]*models.InstructorList, error) {
	InstructorLists, err := s.repo.FindInstructorsListByCourseID(CourseID)
	if err != nil {
		return nil, err
	}
	return InstructorLists, nil
}

func (s *InstructorServiceImpl) GetInstructorsListByListID(ListID uuid.UUID) (*models.InstructorList, error) {
	InstructorLists, err := s.repo.FindInstructorsListByListID(ListID)
	if err != nil {
		return nil, err
	}
	return InstructorLists, nil
}

func (s *InstructorServiceImpl) DeleteInstructorList(InstructorList *models.InstructorList) error {
	deleteInstructorList, err := s.repo.FindInstructorsListByListID(InstructorList.ListID)
	if err != nil {
		return err
	}

	if err := s.repo.RemoveInstructorList(deleteInstructorList); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) CreateEnrollment(CourseID uuid.UUID, Enrollment *models.Enrollment) error {
	if err := s.repo.AddEnrollment(CourseID, Enrollment); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) GetEnrollments() ([]*models.Enrollment, error) {
	Enrollments, err := s.repo.FindEnrollments()
	if err != nil {
		return nil, err
	}
	return Enrollments, nil
}

func (s *InstructorServiceImpl) GetEnrollmentsByCourseID(CourseID uuid.UUID) ([]*models.Enrollment, error) {
	Enrollments, err := s.repo.FindEnrollmentsByCourseID(CourseID)
	if err != nil {
		return nil, err
	}
	return Enrollments, nil
}

func (s *InstructorServiceImpl) DeleteEnrollment(Enrollment *models.Enrollment) error {
	deleteEnrollment, err := s.repo.FindEnrollmentsByCourseID(Enrollment.CourseID)
	if err != nil {
		return err
	}

	for _, enrollment := range deleteEnrollment {
		if err := s.repo.RemoveEnrollment(enrollment); err != nil {
			return err
		}
	}
	return nil
}

func (s *InstructorServiceImpl) GetUsersEnrollment(CourseID uuid.UUID) ([]*models.User, error) {
	Users, err := s.repo.FindUsersEnrollment(CourseID)
	if err != nil {
		return nil, err
	}
	return Users, nil
}

func (s *InstructorServiceImpl) DeleteUserEnrollment(CourseID uuid.UUID, UserID uuid.UUID) error {
	if err := s.repo.RemoveUserEnrollment(CourseID, UserID); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) DeleteEnrollmentsByCourseID(courseID uuid.UUID) error {
	if err := s.repo.RemoveEnrollmentsByCourseID(courseID); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) DeleteAssignmentsByCourseID(courseID uuid.UUID) error {
	if err := s.repo.RemoveAssignmentsByCourseID(courseID); err != nil {
		return err
	}
	return nil
}

func (s *InstructorServiceImpl) DeleteInstructorListsByCourseID(courseID uuid.UUID) error {
	if err := s.repo.RemoveInstructorListsByCourseID(courseID); err != nil {
		return err
	}
	return nil
}
