package adapters

import (
	"backend_fullstack/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type HttpAssignmentHandler struct {
	service *services.AssignmentStudentService
}

func NewHttpAssignmentHandler(service *services.AssignmentStudentService) *HttpAssignmentHandler {
	return &HttpAssignmentHandler{service: service}
}

func (h *HttpAssignmentHandler) GetUserAssignments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	assignments, err := h.service.GetUserAssignments(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(assignments)
}

func (h *HttpAssignmentHandler) DownloadAssignment(c *fiber.Ctx) error {
	assignmentID := c.Params("id")
	fileURL, err := h.service.GetAssignmentFile(assignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Redirect(fileURL, fiber.StatusTemporaryRedirect)
}
