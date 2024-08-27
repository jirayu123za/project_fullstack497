package adapters

import (
	"backend_fullstack/internal/models"

	"gorm.io/gorm"
)

// Secondary adapters
type GormAdminRepository struct {
	db *gorm.DB
}

func NewGormAdminRepository(db *gorm.DB) *GormAdminRepository {
	return &GormAdminRepository{
		db: db,
	}
}

func (r *GormAdminRepository) AddUserGroup(UserGroup *models.UserGroup) error {
	// Implement the logic to AddUserGroup to the database using GORM.
	if userGroup := r.db.Create(UserGroup); userGroup.Error != nil {
		return userGroup.Error
	}
	return nil
}

func (r *GormAdminRepository) FindUserGroupByID(UserGroupID uint) (*models.UserGroup, error) {
	var userGroup *models.UserGroup
	if result := r.db.Preload("Users").First(&userGroup, UserGroupID); result.Error != nil {
		return nil, result.Error
	}
	return userGroup, nil
}

func (r *GormAdminRepository) FindUserGroups() ([]*models.UserGroup, error) {
	var userGroups []*models.UserGroup
	if result := r.db.Preload("Users").Find(&userGroups); result.Error != nil {
		return nil, result.Error
	}
	return userGroups, nil
}

func (r *GormAdminRepository) ModifyUserGroup(UserGroup *models.UserGroup) error {
	var existingUserGroup *models.UserGroup
	if result := r.db.First(&existingUserGroup, "group_id = ?", UserGroup.GroupID); result.Error != nil {
		return result.Error
	}

	existingUserGroup.GroupName = UserGroup.GroupName

	if result := r.db.Save(&existingUserGroup); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormAdminRepository) RemoveUserGroup(UserGroup *models.UserGroup) error {
	var findUserGroup *models.UserGroup
	if result := r.db.First(&findUserGroup, UserGroup.GroupID); result.Error != nil {
		return result.Error
	}

	if result := r.db.Delete(&findUserGroup); result.Error != nil {
		return result.Error
	}
	return nil
}
