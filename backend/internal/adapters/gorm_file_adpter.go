package adapters

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormFileRepository struct {
	db *gorm.DB
}

func NewGormFileRepository(db *gorm.DB) *GormFileRepository {
	return &GormFileRepository{db: db}
}

func (r *GormFileRepository) SaveAssignmentFile(file models.AssignmentFile) error {
	return r.db.Create(&file).Error
}

func (r *GormFileRepository) GetAssignmentFile(fileID uuid.UUID) (models.AssignmentFile, error) {
	var file models.AssignmentFile
	err := r.db.First(&file, fileID).Error
	return file, err
}

func (r *GormFileRepository) DeleteAssignmentFile(fileID uuid.UUID) error {
	return r.db.Delete(&models.AssignmentFile{}, fileID).Error
}
