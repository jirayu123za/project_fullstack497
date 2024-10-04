package adapters

import (
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		groupID := claims["userGroupID"].(float64)
		fmt.Println(groupID)
		if groupID == 1 {
			return c.Redirect("http://localhost:5173/stddash")
		} else if groupID == 2 {
			return c.Redirect("http://localhost:5173/insdash")
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid group ID",
				"groupID": groupID,
			})
		}
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to parse JWT token",
			"error":   err.Error(),
		})
	}
	/*
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
		})
	*/
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

	c.Cookie(&fiber.Cookie{
		Name:    "jwt-token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
