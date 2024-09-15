package adapters

import (
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Primary adapter
type HttpAuthHandler struct {
	services services.AuthService
}

func NewHttpAuthHandler(services services.AuthService) *HttpAuthHandler {
	return &HttpAuthHandler{
		services: services,
	}
}

// Add fuc verify jwt
func (h *HttpAuthHandler) VerifyToken(c *fiber.Ctx) error {
	token := c.Cookies("jwt-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing JWT token",
		})
	}

	err := h.services.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}

func (h *HttpAuthHandler) Login(c *fiber.Ctx) error {
	var loginUser models.User
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	token, err := h.services.Login(loginUser.UserName, loginUser.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to login",
			"error":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *HttpAuthHandler) Logout(c *fiber.Ctx) error {
	token := c.Cookies("jwt-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	err := h.services.Logout(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to logout",
			"error":   err.Error(),
		})
	}

	c.ClearCookie("jwt-token")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
