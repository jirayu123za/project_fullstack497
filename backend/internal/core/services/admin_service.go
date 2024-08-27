package services

import (
	"backend_fullstack/internal/core/repositories"
	"backend_fullstack/internal/models"
)

// Primary port
type AdminService interface {
	CreateUserGroup(UserGroup *models.UserGroup) error
	GetUserGroupByID(UserGroupID uint) (*models.UserGroup, error)
	GetUserGroups() ([]*models.UserGroup, error)
	UpdateUserGroup(UserGroup *models.UserGroup) error
	DeleteUserGroup(UserGroup *models.UserGroup) error
}

type AdminServiceImpl struct {
	repo repositories.AdminRepository
}

// func instance business logic call
func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &AdminServiceImpl{
		repo: repo,
	}
}

func (s *AdminServiceImpl) CreateUserGroup(UserGroup *models.UserGroup) error {
	if err := s.repo.AddUserGroup(UserGroup); err != nil {
		return err
	}
	return nil
}

func (s *AdminServiceImpl) GetUserGroupByID(UserGroupID uint) (*models.UserGroup, error) {
	UserGroup, err := s.repo.FindUserGroupByID(UserGroupID)
	if err != nil {
		return nil, err
	}
	return UserGroup, nil
}

func (s *AdminServiceImpl) GetUserGroups() ([]*models.UserGroup, error) {
	UserGroups, err := s.repo.FindUserGroups()
	if err != nil {
		return nil, err
	}
	return UserGroups, nil
}

func (s *AdminServiceImpl) UpdateUserGroup(UserGroup *models.UserGroup) error {
	existingUserGroups, err := s.repo.FindUserGroupByID(UserGroup.GroupID)
	if err != nil {
		return err
	}

	existingUserGroups.GroupName = UserGroup.GroupName

	if err := s.repo.ModifyUserGroup(existingUserGroups); err != nil {
		return err
	}
	return nil
}

func (s *AdminServiceImpl) DeleteUserGroup(UserGroup *models.UserGroup) error {
	deleteUserGroup, err := s.repo.FindUserGroupByID(UserGroup.GroupID)
	if err != nil {
		return err
	}

	if err := s.repo.RemoveUserGroup(deleteUserGroup); err != nil {
		return err
	}
	return nil
}
