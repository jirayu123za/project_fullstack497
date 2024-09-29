package repositories

import (
	"backend_fullstack/internal/models"

	"github.com/google/uuid"
)

type FileRepository interface {
	SaveAssignmentFile(file models.AssignmentFile) error
	GetAssignmentFile(fileID uuid.UUID) (models.AssignmentFile, error)
	DeleteAssignmentFile(fileID uuid.UUID) error
}
