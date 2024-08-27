package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Primary port
type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

// func instance business logic call
func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) CreateUser(user *models.User) error {
	// hash password from user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := s.repo.Register(user); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) GetUserByID(userID uuid.UUID) (*models.User, error) {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUsers() ([]*models.User, error) {
	users, err := s.repo.FindUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) UpdateUser(user *models.User) error {
	existingUser, err := s.repo.FindUserByID(user.UserID)
	if err != nil {
		return err
	}
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.GroupID = user.GroupID
	existingUser.UserName = user.UserName

	if err := s.repo.ModifyUser(existingUser); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(user *models.User) error {
	deleteUser, err := s.repo.FindUserByID(user.UserID)
	if err != nil {
		return err
	}

	if err := s.repo.RemoveUser(deleteUser); err != nil {
		return err
	}
	return nil
}
