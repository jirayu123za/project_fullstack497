package repositories

import "backend_fullstack/internal/models"

// Secondary ports
type AdminRepository interface {
	AddUserGroup(UserGroup *models.UserGroup) error
	FindUserGroupByID(UserGroupID uint) (*models.UserGroup, error)
	FindUserGroups() ([]*models.UserGroup, error)
	ModifyUserGroup(UserGroup *models.UserGroup) error
	RemoveUserGroup(UserGroup *models.UserGroup) error
}
