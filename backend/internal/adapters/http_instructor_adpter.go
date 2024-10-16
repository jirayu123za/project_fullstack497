package adapters

import (
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Primary adapters
type HttpInstructorHandler struct {
	services     services.InstructorService
	userServices services.UserService
}

func NewHttpInstructorHandler(services services.InstructorService, userServices services.UserService) *HttpInstructorHandler {
	return &HttpInstructorHandler{
		services:     services,
		userServices: userServices,
	}
}

// CreateCourse godoc
// @Summary Create a new course
// @Description Create a course by providing course details
// @Tags Courses
// @Accept  json
// @Produce  json
// @Param course body models.Course true "Course Data"
// @Success 201 {object} models.Course
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router api/CreateCourse [post]
func (h *HttpInstructorHandler) CreateCourse(c *fiber.Ctx) error {
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	if err := h.services.CreateCourse(&course); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create course",
			"error":   err.Error(),
		})
	}

	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	/*
		userIDParam := c.Query("user_id")
		userID, err := uuid.Parse(userIDParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid instructor_id",
				"error":   err.Error(),
			})
		}
	*/

	instructorList := models.InstructorList{
		CourseID: course.CourseID,
		UserID:   userID,
	}

	if err := h.services.CreateInstructorList(course.CourseID, &instructorList); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create instructor list",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":         "Course is created",
		"course":          course,
		"instructor_list": instructorList,
	})
}

// GetCourses godoc
// @Summary Get all courses
// @Description Retrieve a list of all courses
// @Tags Courses
// @Produce  json
// @Success 200 {array} models.Course
// @Failure 500 {object} fiber.Map
// @Router api/QueryCourses [get]
func (h *HttpInstructorHandler) GetCourses(c *fiber.Ctx) error {
	courses, err := h.services.GetCourses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get courses",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Course found",
		"courses": courses,
	})
}

// GetCourseByID godoc
// @Summary Get a course by its ID
// @Description Retrieve a course by providing its course ID
// @Tags Courses
// @Produce  json
// @Param course_id query string true "Course ID"
// @Success 200 {object} models.Course
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /courses/{course_id} [get]
func (h *HttpInstructorHandler) GetCourseByID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	course, err := h.services.GetCourseByID(courseID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Course found from query",
		"course":  course,
	})
}

// UpdateCourse godoc
// @Summary Update a course by its ID
// @Description อัปเดตข้อมูล course
// @Tags Courses
// @Param course_id path string true "Course ID"
// @Param course body models.Course true "Updated Course Data"
// @Success 200 {object} models.Course
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /courses/{course_id} [put]
func (h *HttpInstructorHandler) UpdateCourse(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	course, err := h.services.GetCourseByID(courseID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newCourse := new(models.Course)
	if err := c.BodyParser(&newCourse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	course.CourseName = newCourse.CourseName
	//course.CourseDescription = newCourse.CourseDescription
	//course.Term = newCourse.Term

	if err := h.services.UpdateCourse(course); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update course",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Course is updated",
		"course":  course,
	})
}

// DeleteCourse godoc
// @Summary Delete a course
// @Description Delete a course by providing its course ID
// @Tags Courses
// @Param course_id query string true "Course ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /courses/{course_id} [delete]
func (h *HttpInstructorHandler) DeleteCourse(c *fiber.Ctx) error {
	// Parse course_id from query parameters
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	// Use service layer to delete the course and its related data
	if err := h.services.DeleteCourse(courseID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete course",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Course and related data deleted successfully",
	})
}

// using JWT
func (h *HttpInstructorHandler) GetCourseByUserID(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
			fmt.Println("Recovered from panic:", r)
		}
	}()
	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	courses, err := h.services.GetCourseByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get courses by user ID",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, course := range courses {
		response = append(response, map[string]interface{}{
			"course_id":    course.CourseID,
			"course_name":  course.CourseName,
			"course_code":  course.CourseCode,
			"course_color": course.Color,
			"course_image": course.ImageURL,
			"Assignment":   course.Assignments,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Courses found",
		"courses": response,
	})
}

func (h *HttpInstructorHandler) GetNameByUserID(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
			fmt.Println("Recovered from panic:", r)
		}
	}()
	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	name, err := h.services.GetNameByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get courses by user ID",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "name of instructor found",
		"name":    name,
	})
}

func (h *HttpInstructorHandler) GetPersonDataByUserID(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
			fmt.Println("Recovered from panic:", r)
		}
	}()

	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	user, err := h.services.GetPersonDataByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user by user ID",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	response = append(response, map[string]interface{}{
		"user_id":         user.UserID,
		"user_email":      user.Email,
		"user_first_name": user.FirstName,
		"user_last_name":  user.LastName,
		"user_image_url":  user.ProfileImageURL,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User found",
		"user":    response,
	})
}

func (h *HttpInstructorHandler) GetUserGroupByUserID(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
			fmt.Println("Recovered from panic:", r)
		}
	}()

	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	groupName, err := h.services.GetUserGroupByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user group by user ID",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":         "Group name of user found",
		"user_group_name": groupName,
	})
}

func (h *HttpInstructorHandler) GetAssignmentByUserID(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
			fmt.Println("Recovered from panic:", r)
		}
	}()

	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	assignments, err := h.services.GetAssignmentByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get assignments by user ID",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, assignment := range assignments {
		response = append(response, map[string]interface{}{
			"course_id":              assignment.CourseID,
			"assignment_id":          assignment.AssignmentID,
			"assignment_name":        assignment.AssignmentName,
			"assignment_description": assignment.AssignmentDescription,
			"due_date":               assignment.DueDate.Format("02-01-2006"),
			"Submissions":            assignment.Submissions,
			"AssignmentFiles":        assignment.AssignmentFiles,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Assignments found",
		"assignments": response,
	})
}

func (h *HttpInstructorHandler) GetAssignmentByUserIDSorted(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
			fmt.Println("Recovered from panic:", r)
		}
	}()

	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	assignments, err := h.services.GetAssignmentByUserIDSorted(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get assignments by user ID",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, assignment := range assignments {
		response = append(response, map[string]interface{}{
			"course_id":       assignment.CourseID,
			"assignment_id":   assignment.AssignmentID,
			"assignment_name": assignment.AssignmentName,
			"due_date":        assignment.DueDate.Format("02-01-2006"),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Assignments found",
		"assignments": response,
	})
}

// Under line here be HttpInstructorHandler of Instructor assignment
func (h *HttpInstructorHandler) CreateAssignment(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	var assignment models.Assignment
	if err := c.BodyParser(&assignment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	if err := h.services.CreateAssignment(courseID, &assignment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create assignment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Assignment is created",
		"assignment": assignment,
	})
}

func (h HttpInstructorHandler) GetAssignmentByAssignmentID(c *fiber.Ctx) error {
	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	assignment, err := h.services.GetAssignmentByAssignmentID(assignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get assignment by assignment ID",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Assignment found",
		"assignment": assignment,
	})
}

func (h HttpInstructorHandler) GetAssignments(c *fiber.Ctx) error {
	assignments, err := h.services.GetAssignments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get assignments",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Assignments found",
		"assignments": assignments,
	})
}

func (h *HttpInstructorHandler) GetAssignmentsByCourseID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	assignments, err := h.services.GetAssignmentsByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get assignments by course ID",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, assignment := range assignments {
		dueDate := assignment.DueDate.Format("02-01-2006")

		response = append(response, map[string]interface{}{
			"course_id":              assignment.CourseID,
			"assignment_id":          assignment.AssignmentID,
			"assignment_name":        assignment.AssignmentName,
			"assignment_description": assignment.AssignmentDescription,
			"assignment_due_date":    dueDate,
		})
	}

	// Modify the response to only return ...
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Assignments are retrieved",
		"assignments": response,
	})
}

func (h *HttpInstructorHandler) GetAssignmentByCourseIDAndAssignmentID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	assignment, err := h.services.GetAssignmentByCourseIDAndAssignmentID(courseID, assignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get assignment by course ID and assignment ID",
			"error":   err.Error(),
		})
	}

	response := map[string]interface{}{
		"course_id":              assignment.CourseID,
		"assignment_id":          assignment.AssignmentID,
		"assignment_name":        assignment.AssignmentName,
		"assignment_description": assignment.AssignmentDescription,
		"due_date":               assignment.DueDate,
		"AssignmentFiles":        assignment.AssignmentFiles,
		"Submissions":            assignment.Submissions,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Assignment found",
		"assignment": response,
	})
}

func (h *HttpInstructorHandler) UpdateAssignment(c *fiber.Ctx) error {
	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	assignment, err := h.services.GetAssignmentByAssignmentID(assignmentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newAssignment := new(models.Assignment)
	if err := c.BodyParser(&newAssignment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	assignment.AssignmentName = newAssignment.AssignmentName
	assignment.AssignmentDescription = newAssignment.AssignmentDescription
	assignment.DueDate = newAssignment.DueDate

	if err := h.services.UpdateAssignment(assignment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update assignment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "assignment is updated",
		"assignment": assignment,
	})
}

func (h *HttpInstructorHandler) UpdateAssignmentByCourseIDAndAssignmentID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	assignment, err := h.services.GetAssignmentByCourseIDAndAssignmentID(courseID, assignmentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newAssignment := new(models.Assignment)
	if err := c.BodyParser(&newAssignment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	assignment.AssignmentDescription = newAssignment.AssignmentDescription
	assignment.DueDate = newAssignment.DueDate

	if err := h.services.UpdateAssignmentByCourseIDAndAssignmentID(courseID, assignmentID, assignment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update assignment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "assignment is updated",
		"assignment": assignment,
	})
}

func (h *HttpInstructorHandler) DeleteAssignment(c *fiber.Ctx) error {
	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	assignment, err := h.services.GetAssignmentByAssignmentID(assignmentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.services.DeleteAssignment(assignment.AssignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete assignment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Assignment is deleted",
	})
}

func (h *HttpInstructorHandler) DeleteAssignmentByCourseIDAndAssignmentID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	err = h.services.DeleteAssignmentsByCourseIDAndAssignmentID(courseID, assignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete assignment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Assignment is deleted",
	})
}

func (h *HttpInstructorHandler) GetSubmissionsByCourseIDAndAssignmentID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	assignmentIDParam := c.Query("assignment_id")
	assignmentID, err := uuid.Parse(assignmentIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment_id",
			"error":   err.Error(),
		})
	}

	submissions, err := h.services.GetSubmissionsByCourseIDAndAssignmentID(courseID, assignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get submissions by course ID and assignment ID",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, submission := range submissions {
		response = append(response, map[string]interface{}{
			"user_id":        submission.UserID,
			"user_name":      submission.FirstName + " " + submission.LastName,
			"user_submitted": submission.Submitted,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Submissions found",
		"submissions": response,
	})
}

// Using minio
func (h *HttpInstructorHandler) UploadAssignmentFile(c *fiber.Ctx) error {

	/*
		userIDStr := c.FormValue("user_id")
		assignmentIDStr := c.FormValue("assignment_id")
		userGroupName := c.FormValue("user_group_name")
		userName := c.FormValue("user_name")

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid user ID",
				"error":   err.Error(),
			})
		}

		assignmentID, err := uuid.Parse(assignmentIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid assignment ID",
				"error":   err.Error(),
			})
		}

		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to get file from request",
				"error":   err.Error(),
			})
		}

		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to open file",
				"error":   err.Error(),
			})
		}
		defer file.Close()

		fileExtension := filepath.Ext(fileHeader.Filename)

		uploadID, err := h.services.UploadAssignmentFile(userID, assignmentID, userGroupName, userName, file, fileExtension)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload assignment file",
				"error":   err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "File uploaded successfully",
			"upload_id": uploadID,
		})
	*/

	userIDStr := c.FormValue("user_id")
	assignmentIDStr := c.FormValue("assignment_id")
	userGroupName := c.FormValue("user_group_name")
	userName := c.FormValue("user_name")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	assignmentID, err := uuid.Parse(assignmentIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assignment ID",
			"error":   err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get files from request",
			"error":   err.Error(),
		})
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No files uploaded",
		})
	}

	uploadIDs, err := h.services.UploadAssignmentFiles(userID, assignmentID, userGroupName, userName, files)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload files",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Files uploaded successfully",
		"upload_ids": uploadIDs,
	})
}

// Under line here be HttpInstructorHandler of Instructor list
func (h *HttpInstructorHandler) CreateInstructorList(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	var instructorList models.InstructorList
	if err := c.BodyParser(&instructorList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	if err := h.services.CreateInstructorList(courseID, &instructorList); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create instructor list",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":         "Instructor list is created",
		"instructor_list": instructorList,
	})
}

func (h *HttpInstructorHandler) GetInstructorsList(c *fiber.Ctx) error {
	instructorLists, err := h.services.GetInstructorsList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get instructor lists",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          "Instructor lists found",
		"instructor_lists": instructorLists,
	})
}

func (h *HttpInstructorHandler) GetInstructorsListByCourseID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	instructorLists, err := h.services.GetInstructorsListByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get instructor lists",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":          "Instructor lists found",
		"instructor_lists": instructorLists,
	})
}

func (h *HttpInstructorHandler) GetInstructorsListByListID(c *fiber.Ctx) error {
	listIDParam := c.Query("list_id")
	listID, err := uuid.Parse(listIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid list_id",
			"error":   err.Error(),
		})
	}

	instructorList, err := h.services.GetInstructorsListByListID(listID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get instructor list",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":         "Instructor list found",
		"instructor_list": instructorList,
	})
}

func (h *HttpInstructorHandler) DeleteInstructorList(c *fiber.Ctx) error {
	listIDParam := c.Query("list_id")
	listID, err := uuid.Parse(listIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid list_id",
			"error":   err.Error(),
		})
	}

	instructorList, err := h.services.GetInstructorsListByListID(listID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.services.DeleteInstructorList(instructorList)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete instructor list",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Instructor list is deleted",
	})
}

// under line for instructions enrollments
func (h *HttpInstructorHandler) CreateEnrollment(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	var payload struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	userID, err := h.userServices.GetUserIDByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user_id by email",
			"error":   err.Error(),
		})
	}

	enrollment := models.Enrollment{
		UserID:   userID,
		CourseID: courseID,
	}

	if err := h.services.CreateEnrollment(courseID, &enrollment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create enrollment",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Enrollment is created",
		"enrollment": enrollment,
	})
}

func (h *HttpInstructorHandler) GetEnrollments(c *fiber.Ctx) error {
	enrollments, err := h.services.GetEnrollments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get enrollments",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Enrollments found",
		"enrollments": enrollments,
	})
}

func (h *HttpInstructorHandler) GetEnrollmentsByCourseID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	enrollments, err := h.services.GetEnrollmentsByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get enrollments by course ID",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Enrollments found",
		"enrollments": enrollments,
	})
}

func (h *HttpInstructorHandler) DeleteEnrollment(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	enrollments, err := h.services.GetEnrollmentsByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get enrollments by course ID",
			"error":   err.Error(),
		})
	}

	for _, enrollment := range enrollments {
		if err := h.services.DeleteEnrollment(enrollment); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to delete enrollment",
				"error":   err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Enrollments are deleted",
	})
}

func (h *HttpInstructorHandler) GetUsersEnrollment(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	users, err := h.services.GetUsersEnrollment(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get users by course_id",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, user := range users {
		response = append(response, map[string]interface{}{
			"user_id":           user.UserID,
			"first_name":        user.FirstName,
			"last_name":         user.LastName,
			"profile_image_url": user.ProfileImageURL,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users found",
		"users":   response,
	})
}

func (h *HttpInstructorHandler) DeleteUserEnrollment(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	userIDParam := c.Query("user_id")

	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course_id",
		})
	}

	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user_id",
		})
	}

	err = h.services.DeleteUserEnrollment(courseID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to remove enrollment",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Enrollment removed successfully",
	})
}

// DeleteAssignmentsByCourseID godoc
// @Summary Delete assignments by course ID
// @Description Delete all assignments related to a specific course by its course ID
// @Tags Assignments
// @Param course_id query string true "Course ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /assignments/delete [delete]
func (h *HttpInstructorHandler) DeleteAssignmentsByCourseID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course_id",
		})
	}

	assignments, err := h.services.GetAssignmentsByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get assignments by course ID",
		})
	}

	for _, assignment := range assignments {
		if err := h.services.DeleteAssignment(assignment.AssignmentID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete assignment",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Assignments are deleted",
	})
}

// DeleteInstructorListsByCourseID godoc
// @Summary Delete instructor lists by course ID
// @Description Delete all instructor lists related to a specific course by its course ID
// @Tags InstructorLists
// @Param course_id query string true "Course ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /instructors/lists/delete [delete]
func (h *HttpInstructorHandler) DeleteInstructorListsByCourseID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course_id",
		})
	}

	instructorLists, err := h.services.GetInstructorsListByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get instructor lists by course ID",
		})
	}

	for _, instructorList := range instructorLists {
		if err := h.services.DeleteInstructorList(instructorList); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete instructor list",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Instructor lists are deleted",
	})
}

// DeleteEnrollmentsByCourseID godoc
// @Summary Delete enrollments by course ID
// @Description Delete all enrollments related to a specific course by its course ID
// @Tags Enrollments
// @Param course_id query string true "Course ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /enrollments/delete [delete]
func (h *HttpInstructorHandler) DeleteEnrollmentsByCourseID(c *fiber.Ctx) error {
	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid course_id",
		})
	}

	enrollments, err := h.services.GetEnrollmentsByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get enrollments by course ID",
		})
	}

	for _, enrollment := range enrollments {
		if err := h.services.DeleteEnrollment(enrollment); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete enrollment",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Enrollments are deleted",
	})
}
