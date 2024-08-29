package adapters

import (
	"backend_fullstack/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) AuthenticateUser(username string, password string) (string, error) {
	// Implement the logic to authenticate the user
	var selectedUser models.User
	if result := r.db.Where("user_name = ?", username).First(&selectedUser); result.Error != nil {
		return "", result.Error
	}

	// then compare password of before user
	err := bcrypt.CompareHashAndPassword(
		[]byte(selectedUser.Password),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return selectedUser.UserID.String(), nil
}

func (r *AuthRepository) DeleteJWTToken(token string) error {
	return nil
}
