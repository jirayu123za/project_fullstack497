package repositories

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

// Secondary ports
type UserRepository interface {
	Register(user *models.User) error
	FindUserByID(userID uuid.UUID) (*models.User, error)
	FindUserByUserName(userName string) (*models.User, error)
	FindUserIDByEmail(email string) (uuid.UUID, error)
	FindUsers() ([]*models.User, error)
	ModifyUser(user *models.User) error
	RemoveUser(user *models.User) error
}
