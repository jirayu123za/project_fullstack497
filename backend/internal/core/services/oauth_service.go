package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"

	"golang.org/x/oauth2"
)

// Primary port
type OAuthService interface {
	GetGoogleLoginURL(state string) string
	GetGoogleToken(code string) (*oauth2.Token, error)
	GetGoogleUserInfo(accessToken string) (*models.GoogleUserInfo, error)
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
