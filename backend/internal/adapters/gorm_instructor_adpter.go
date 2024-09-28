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
	//existingCourse.CourseDescription = Course.CourseDescription

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

// Using jwt
func (r *GormInstructorRepository) FindCourseByUserID(UserID uuid.UUID) ([]*models.Course, error) {
	var courses []*models.Course
	if err := r.db.
		Joins("JOIN instructor_lists ON instructor_lists.course_id = courses.course_id").
		Where("instructor_lists.user_id = ?", UserID).
		Preload("Assignments").
		Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *GormInstructorRepository) FindNameByUserID(userID uuid.UUID) (string, error) {
	var user models.User
	if result := r.db.First(&user, "user_id = ?", userID); result.Error != nil {
		return "", result.Error
	}
	return user.FirstName + " " + user.LastName, nil
}

func (r *GormInstructorRepository) FindPersonDataByUserID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if result := r.db.First(&user, "user_id = ?", userID); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormInstructorRepository) FindUserGroupByUserID(userID uuid.UUID) (string, error) {
	var groupName string
	if result := r.db.
		Table("users").
		Select("user_groups.group_name").
		Joins("JOIN user_groups ON user_groups.group_id = users.group_id").
		Where("users.user_id = ?", userID).
		Scan(&groupName); result.Error != nil {
		return "", result.Error
	}
	return groupName, nil
}

func (r *GormInstructorRepository) FindAssignmentByUserID(userID uuid.UUID) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if err := r.db.
		Joins("JOIN courses ON courses.course_id = assignments.course_id").
		Joins("JOIN instructor_lists ON instructor_lists.course_id = courses.course_id").
		Where("instructor_lists.user_id = ?", userID).
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

// Under line here be GormInstructorRepository of Instructor assignments
func (r *GormInstructorRepository) AddAssignment(CourseID uuid.UUID, Assignment *models.Assignment) error {
	// Implement the logic to AddAssignment to the database using GORM.
	var findCourse *models.Course
	if result := r.db.First(&findCourse, "course_id = ?", CourseID); result.Error != nil {
		return result.Error
	}
	Assignment.CourseID = CourseID
	if assignment := r.db.Create(Assignment); assignment.Error != nil {
		return assignment.Error
	}
	return nil
}

func (r *GormInstructorRepository) FindAssignmentByAssignmentID(AssignmentID uuid.UUID) (*models.Assignment, error) {
	var assignment *models.Assignment
	if result := r.db.First(&assignment, "assignment_id = ?", AssignmentID); result.Error != nil {
		return nil, result.Error
	}
	return assignment, nil
}

func (r *GormInstructorRepository) FindAssignments() ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if result := r.db.Find(&assignments); result.Error != nil {
		return nil, result.Error
	}
	return assignments, nil
}

func (r *GormInstructorRepository) FindAssignmentsByCourseID(CourseID uuid.UUID) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if result := r.db.Find(&assignments, "course_id = ?", CourseID); result.Error != nil {
		return nil, result.Error
	}
	return assignments, nil
}

func (r *GormInstructorRepository) ModifyAssignment(assignment *models.Assignment) error {
	var existingAssignment *models.Assignment
	if result := r.db.Find(&existingAssignment, "assignment_id = ?", assignment.AssignmentID); result.Error != nil {
		return result.Error
	}
	existingAssignment.AssignmentName = assignment.AssignmentName
	existingAssignment.AssignmentDescription = assignment.AssignmentDescription
	existingAssignment.DueDate = assignment.DueDate

	if result := r.db.Save(&existingAssignment); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormInstructorRepository) RemoveAssignment(AssignmentID uuid.UUID) error {
	var findAssignment *models.Assignment
	if result := r.db.First(&findAssignment, "assignment_id = ?", AssignmentID); result.Error != nil {
		return result.Error
	}

	if result := r.db.Delete(&findAssignment); result.Error != nil {
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
