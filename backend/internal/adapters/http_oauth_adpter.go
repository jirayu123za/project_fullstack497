package adapters

import (
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/core/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Primary adapter
type HttpOAuthHandler struct {
	services services.OAuthService
}

func NewHttpOAuthHandler(services services.OAuthService) *HttpOAuthHandler {
	return &HttpOAuthHandler{
		services: services,
	}
}

func (h *HttpOAuthHandler) GetGoogleLoginURL(c *fiber.Ctx) error {
	state, err := utils.GenerateRandomState(32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate random state",
			"error":   err,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
	})

	loginURL := h.services.GetGoogleLoginURL(state)

	return c.Redirect(loginURL)
}

func (h *HttpOAuthHandler) GetGoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	if !utils.ValidateState(c, state) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid state parameter",
		})
	}

	token, err := h.services.GetGoogleToken(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to exchange code for token",
			"error":   err,
		})
	}

	userInfo, err := h.services.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user info",
			"error":   err,
		})
	}
	/*
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User is logged in by Google OAuth2",
			"token":   token,
			"user":    userInfo,
		})
	*/

	jwtToken, err := h.services.GenerateGoogleJWT(userInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate JWT token",
			"error":   err.Error(),
		})
	}

	redirectURL := "http://localhost:5173/signup?token=" + jwtToken
	return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)

}
