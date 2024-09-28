package adapters

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Secondary adapters
type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) Register(user *models.User) error {
	// Implement the logic to register to the database using GORM.
	if register := r.db.Create(user); register.Error != nil {
		return register.Error
	}
	return nil
}

func (r *GormUserRepository) FindUserByID(userID uuid.UUID) (*models.User, error) {
	var user *models.User
	if result := r.db.Preload(
		"Enrollments").Preload(
		"InstructorLists").Preload(
		"Submissions").Preload(
		"Uploads").First(
		&user, userID); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *GormUserRepository) FindUserByUserName(userName string) (*models.User, error) {
	var user *models.User
	if result := r.db.Where("user_name = ?", userName).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *GormUserRepository) FindUserIDByEmail(email string) (uuid.UUID, error) {
	var user *models.User
	if result := r.db.Where("email = ?", email).First(&user); result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.UserID, nil
}

func (r *GormUserRepository) FindUsers() ([]*models.User, error) {
	var users []*models.User
	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *GormUserRepository) ModifyUser(user *models.User) error {
	var existingUser *models.User
	if result := r.db.First(&existingUser, user.UserID); result.Error != nil {
		return result.Error
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.GroupID = user.GroupID
	existingUser.UserName = user.UserName

	if result := r.db.Save(&existingUser); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormUserRepository) RemoveUser(user *models.User) error {
	var findUser *models.User
	if result := r.db.First(&findUser, user.UserID); result.Error != nil {
		return result.Error
	}

	if result := r.db.Delete(&findUser); result.Error != nil {
		return result.Error
	}
	return nil
}
