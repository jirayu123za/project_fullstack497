package adapters

import (
	"backend_fullstack/internal/core/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpAssignmentHandler struct {
	service *services.AssignmentStudentService
}

func NewHttpAssignmentHandler(service *services.AssignmentStudentService) *HttpAssignmentHandler {
	return &HttpAssignmentHandler{service: service}
}

func (h *HttpAssignmentHandler) GetUserAssignments(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Locals("userID").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	assignments, err := h.service.GetUserAssignments(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(assignments)
}

func (h *HttpAssignmentHandler) DownloadAssignment(c *fiber.Ctx) error {
	assignmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid assignment ID"})
	}
	fileURL, err := h.service.GetAssignmentFile(assignmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Redirect(fileURL, fiber.StatusTemporaryRedirect)
}

func (h *HttpAssignmentHandler) SubmitAssignment(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Locals("userID").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	assignmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid assignment ID"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
	}

	// Create a unique filename
	filename := uuid.New().String() + "_" + file.Filename

	// Define the path where the file will be saved
	uploadPath := "./uploads/" + filename

	// Save the file
	if err := c.SaveFile(file, uploadPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	err = h.service.SubmitAssignment(assignmentID, userID, filename)
	if err != nil {
		// If submission fails, we should delete the uploaded file
		os.Remove(uploadPath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Assignment submitted successfully",
	})
}

func (h *HttpAssignmentHandler) DeleteSubmission(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Locals("userID").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	assignmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid assignment ID"})
	}

	err = h.service.DeleteSubmission(assignmentID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Submission deleted successfully"})
}
