package adapters

import (
	"backend_fullstack/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpAssignmentHandler struct {
	service services.AssignmentService
}

func NewHttpAssignmentHandler(service services.AssignmentService) *HttpAssignmentHandler {
	return &HttpAssignmentHandler{
		service: service,
	}
}

func (h *HttpAssignmentHandler) GetAssignmentDashboard(c *fiber.Ctx) error {
	studentIDStr := c.Locals("user_id").(string)
	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	dashboardItems, err := h.service.GetAssignmentDashboard(studentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch assignment dashboard",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Assignment dashboard fetched successfully",
		"data":    dashboardItems,
	})
}

func (h *HttpAssignmentHandler) GetAssignment(c *fiber.Ctx) error {
	assignmentID := c.Params("id")
	assignment, err := h.service.GetAssignmentByID(assignmentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Assignment not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Assignment fetched successfully",
		"data":    assignment,
	})
}

// Add other handler methods as needed (e.g., CreateAssignment, UpdateAssignment, DeleteAssignment)
