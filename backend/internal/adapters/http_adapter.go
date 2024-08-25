package adapters

import (
	"backend_fullstack/internal/adapters/auth"
	"backend_fullstack/internal/core/utils" // Import the package that contains GenerateRandomState

	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {
	state, err := utils.GenerateRandomState(32) // Use the GenerateRandomState function from the imported package
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
		"message": "User is logged in",
		"token":   token,
		"user":    userInfo,
	})
}
