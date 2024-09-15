package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthService interface {
	Login(userName string, password string) (string, error)
	Logout(token string) error
	VerifyToken(token string) error
}

type AuthServiceImpl struct {
	repo      repositories.AuthRepository
	userRepo  repositories.UserRepository
	jwtSecret string
}

func NewAuthService(repo repositories.AuthRepository, userRepo repositories.UserRepository, jwtSecret string) AuthService {
	return &AuthServiceImpl{
		repo:      repo,
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthServiceImpl) Login(userName string, password string) (string, error) {
	userID, err := s.repo.AuthenticateUser(userName, password)
	if err != nil {
		return "", err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.FindUserByID(userUUID)
	if err != nil {
		return "", err
	}

	token, err := s.generateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthServiceImpl) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"userID":      user.UserID,
		"userGroupID": user.GroupID,
		"exp":         time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthServiceImpl) VerifyToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := parsedToken.Claims.(*jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println(claims)
		return nil
	} else {
		return fmt.Errorf("invalid token")
	}
}

func (s *AuthServiceImpl) Logout(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return err
	}

	return s.repo.DeleteJWTToken(token)
}
