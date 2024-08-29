package adapters

import (
	"backend_fullstack/internal/adapters/auth"
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/core/utils"
	"backend_fullstack/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GoogleLogin(c *fiber.Ctx) error {
	state, err := utils.GenerateRandomState(32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate random state",
			"error":   err,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "oauth_state",
		Value: state,
	})

	loginURL := auth.GetGoogleLoginURL(state)

	return c.Redirect(loginURL)
}

func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	if !utils.ValidateState(c, state) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid state parameter",
		})
	}

	token, err := auth.GetGoogleToken(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to exchange code for token",
			"error":   err,
		})
	}

	userInfo, err := auth.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user info",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User is logged in by google OAuth2",
		"token":   token,
		"user":    userInfo,
	})
}

// Primary adapter
type HttpUserHandler struct {
	services services.UserService
}

func NewHttpUserHandler(services services.UserService) *HttpUserHandler {
	return &HttpUserHandler{
		services: services,
	}
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	if err := h.services.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User is registered",
		"user":    user,
	})
}

func (h *HttpUserHandler) GetUserByID(c *fiber.Ctx) error {
	// Query user_id from the URL
	userIDParam := c.Query("user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id",
			"error":   err.Error(),
		})
	}

	user, err := h.services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User found from query",
		"user":    user,
	})
}

func (h *HttpUserHandler) GetUserByUserName(c *fiber.Ctx) error {
	userName := c.Query("user_name")
	user, err := h.services.GetUserByUserName(userName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User found from query",
		"user":    user,
	})
}

func (h *HttpUserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.services.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users found",
		"users":   users,
	})
}

func (h *HttpUserHandler) UpdateUser(c *fiber.Ctx) error {
	// Query user_id from the URL
	userIDParam := c.Query("user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id",
			"error":   err.Error(),
		})
	}

	user, err := h.services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	newUser := new(models.User)
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.GroupID = newUser.GroupID
	user.UserName = newUser.UserName

	if err := h.services.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User is updated",
		"user":    user,
	})
}

func (h *HttpUserHandler) DeleteUser(c *fiber.Ctx) error {
	userIDParam := c.Query("user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id",
			"error":   err.Error(),
		})
	}

	user, err := h.services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	err = h.services.DeleteUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User successfully deleted",
	})
}
