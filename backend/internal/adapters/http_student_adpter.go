package adapters

import (
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/core/services"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Primary adapters
type HttpStudentHandler struct {
	services services.StudentService
}

func NewHttpStudentHandler(services services.StudentService) *HttpStudentHandler {
	return &HttpStudentHandler{
		services: services,
	}
}

// using JWT
func (h *HttpStudentHandler) GetCourseByUserIDStd(c *fiber.Ctx) (err error) {
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

	courses, err := h.services.GetCourseByUserIDStd(userID)
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

func (h *HttpStudentHandler) GetAssignmentByUserIDStd(c *fiber.Ctx) (err error) {
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

	assignments, err := h.services.GetAssignmentByUserIDStd(userID)
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

func (h *HttpStudentHandler) GetAssignmentByUserIDSortedStd(c *fiber.Ctx) (err error) {
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

	assignments, err := h.services.GetAssignmentByUserIDSortedStd(userID)
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

func (h *HttpStudentHandler) GetUpcomingAssignments(c *fiber.Ctx) (err error) {
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

	courseIDParam := c.Query("course_id")
	courseID, err := uuid.Parse(courseIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course_id",
			"error":   err.Error(),
		})
	}

	assignments, err := h.services.GetUpcomingAssignments(userID, courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get upcoming assignments",
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
		"message":     "Upcoming assignments found",
		"assignments": response,
	})
}

func (h *HttpStudentHandler) UploadAssignmentFile(c *fiber.Ctx) error {
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

	uploadIDs, err := h.services.CreateAssignmentFiles(userID, assignmentID, userGroupName, userName, files)
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

func (h *HttpStudentHandler) GetSubmissionsStatus(c *fiber.Ctx) error {
	courseIDStr := c.Query("course_id")
	assignmentIDStr := c.Query("assignment_id")

	courseID, err := uuid.Parse(courseIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course ID",
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

	userToken := c.Cookies("jwt-token")
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, _ := jwt.ParseWithClaims(userToken, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(jwtSecret)
		return []byte(jwtSecret), nil
	})

	claims, ok := parsedToken.Claims.(*jwt.MapClaims)
	if !ok || claims == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JWT claims",
		})
	}

	userID, err := uuid.Parse((*claims)["userID"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id in JWT",
			"error":   err.Error(),
		})
	}

	submissions, err := h.services.GetSubmissionsStatus(courseID, assignmentID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get submissions status",
			"error":   err.Error(),
		})
	}

	var response []map[string]interface{}
	for _, submission := range submissions {
		response = append(response, map[string]interface{}{
			"user_id":         submission.UserID,
			"user_name":       submission.FirstName + " " + submission.LastName,
			"user_submission": submission.Submission,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Submission status found",
		"submissions": response,
	})
}
