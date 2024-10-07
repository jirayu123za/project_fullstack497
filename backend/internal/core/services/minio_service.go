package services

import (
	"backend_fullstack/internal/core/repositories"
	"mime/multipart"
)

// Primary port
type MinIOService interface {
	CreateFileToMinIO(file multipart.File, userGroupName, userName string, fileName string) error
	//CreateFileToMinIO(file multipart.File, userGroupName, userName string, fileExtension string) error
}

type MinIOServiceServiceImpl struct {
	repo repositories.MinIORepository
}

// func instance business logic call
func NewMinIOServiceService(repo repositories.MinIORepository) MinIOService {
	return &MinIOServiceServiceImpl{
		repo: repo,
	}
}

/*
func (s *MinIOServiceServiceImpl) CreateFileToMinIO(file multipart.File, userGroupName, userName, fileExtension string) error {
	if err := s.repo.SaveFileToMinIO(file, userGroupName, userName, fileExtension); err != nil {
		return err
	}
	return nil
}
*/

func (s *MinIOServiceServiceImpl) CreateFileToMinIO(file multipart.File, userGroupName, userName, fileName string) error {
	if err := s.repo.SaveFileToMinIO(file, userGroupName, userName, fileName); err != nil {
		return err
	}
	return nil
}
