package services

import (
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

// Primary port
type OAuthService interface {
	GetGoogleLoginURL(state string) string
	GetGoogleToken(code string) (*oauth2.Token, error)
	GetGoogleUserInfo(accessToken string) (*models.GoogleUserInfo, error)
	GenerateGoogleJWT(googleUser *models.GoogleUserInfo) (string, error)
}

type OAuthServiceImpl struct {
	repo repositories.OAuthRepository
}

func NewOAuthService(repo repositories.OAuthRepository) OAuthService {
	return &OAuthServiceImpl{
		repo: repo,
	}
}

func (s *OAuthServiceImpl) GetGoogleLoginURL(state string) string {
	return s.repo.GetGoogleLoginURL(state)
}

func (s *OAuthServiceImpl) GetGoogleToken(code string) (*oauth2.Token, error) {
	return s.repo.GetGoogleToken(code)
}

func (s *OAuthServiceImpl) GetGoogleUserInfo(accessToken string) (*models.GoogleUserInfo, error) {
	userInfo, err := s.repo.GetGoogleUserInfo(accessToken)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (s *OAuthServiceImpl) GenerateGoogleJWT(googleUser *models.GoogleUserInfo) (string, error) {
	config.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"email":     googleUser.Email,
		"fullName":  googleUser.Name,
		"firstName": googleUser.GivenName,
		"lastName":  googleUser.FamilyName,
		"picture":   googleUser.Picture,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
