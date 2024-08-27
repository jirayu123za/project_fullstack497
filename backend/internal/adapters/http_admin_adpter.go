package adapters

import (
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Primary adapters
type HttpAdminHandler struct {
	services services.AdminService
}

func NewHttpAdminHandler(services services.AdminService) *HttpAdminHandler {
	return &HttpAdminHandler{
		services: services,
	}
}

func (h *HttpAdminHandler) CreateUserGroup(c *fiber.Ctx) error {
	var userGroup models.UserGroup
	if err := c.BodyParser(&userGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	if err := h.services.CreateUserGroup(&userGroup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user group",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "User group is created",
		"userGroup": userGroup,
	})
}

func (h *HttpAdminHandler) GetUserGroupByID(c *fiber.Ctx) error {
	userGroupID := c.Query("group_id")
	groupID, err := strconv.ParseUint(userGroupID, 0, 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid group_id",
			"error":   err.Error(),
		})
	}

	userGroup, err := h.services.GetUserGroupByID(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "User group found from query",
		"userGroup": userGroup,
	})
}

func (h *HttpAdminHandler) GetUserGroups(c *fiber.Ctx) error {
	userGroups, err := h.services.GetUserGroups()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "User groups found",
		"userGroups": userGroups,
	})
}

func (h *HttpAdminHandler) UpdateUserGroup(c *fiber.Ctx) error {
	groupID, err := strconv.ParseUint(c.Query("group_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid group_id",
			"error":   err.Error(),
		})
	}

	userGroup, err := h.services.GetUserGroupByID(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	newUserGroup := new(models.UserGroup)
	if err := c.BodyParser(&newUserGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	userGroup.GroupName = newUserGroup.GroupName

	if err := h.services.UpdateUserGroup(userGroup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user group",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "User group is updated",
		"userGroup": userGroup,
	})
}

func (h *HttpAdminHandler) DeleteUserGroup(c *fiber.Ctx) error {
	userGroupID := c.Query("group_id")
	groupID, err := strconv.ParseUint(userGroupID, 0, 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid group_id",
			"error":   err.Error(),
		})
	}

	userGroup, err := h.services.GetUserGroupByID(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	err = h.services.DeleteUserGroup(userGroup)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user group",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User group is deleted",
	})
}
