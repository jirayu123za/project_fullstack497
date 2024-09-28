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
	services services.InstructorService
}

func NewHttpInstructorHandler(services services.InstructorService) *HttpInstructorHandler {
	return &HttpInstructorHandler{
		services: services,
	}
}

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

func (h *HttpInstructorHandler) DeleteCourse(c *fiber.Ctx) error {
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

	instructorLists, err := h.services.GetInstructorsListByCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get instructor lists by course ID",
			"error":   err.Error(),
		})
	}

	for _, instructorList := range instructorLists {
		if err := h.services.DeleteInstructorList(instructorList); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to delete instructor list",
				"error":   err.Error(),
			})
		}
	}

	err = h.services.DeleteCourse(course)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete course",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Course is deleted",
	})
}

// using JWT
func (h *HttpInstructorHandler) GetCourseByUserID(c *fiber.Ctx) error {
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

func (h *HttpInstructorHandler) GetNameByUserID(c *fiber.Ctx) error {
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

func (h *HttpInstructorHandler) GetPersonDataByUserID(c *fiber.Ctx) error {
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

func (h *HttpInstructorHandler) GetUserGroupByUserID(c *fiber.Ctx) error {
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

func (h *HttpInstructorHandler) GetAssignmentByUserID(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Assignments found",
		"assignments": assignments,
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Assignments found",
		"assignments": assignments,
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
