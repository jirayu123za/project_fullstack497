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
	course.CourseDescription = newCourse.CourseDescription
	course.Term = newCourse.Term

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
