package adapters

import (
	"backend_fullstack/internal/config"
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
)

// Secondary adapters
type MinIOFileRepository struct {
	client *minio.Client
}

func NewMinIOFileRepository(client *minio.Client) *MinIOFileRepository {
	return &MinIOFileRepository{
		client: client,
	}
}

func (r *MinIOFileRepository) SaveFileToMinIO(file multipart.File, userGroupName, userName, fileName string) error {
	config.LoadEnv()
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	ctx := context.Background()
	exists, err := r.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		if err := r.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "ap-southeast-1"}); err != nil {
			return err
		}
	}

	objectName := filepath.Join(userGroupName, userName, fileName)
	objectName = strings.ReplaceAll(objectName, "\\", "/")

	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, fileName)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return err
	}

	contentType := "application/octet-stream"
	if strings.HasSuffix(fileName, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".jpeg") {
		contentType = "image/jpeg"
	} else if strings.HasSuffix(fileName, ".pdf") {
		contentType = "application/pdf"
	}

	_, err = r.client.FPutObject(ctx, bucketName, objectName, tempFilePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}

	defer os.Remove(tempFilePath)

	return nil
}

/*
func (r *MinIOFileRepository) SaveFileToMinIO(file multipart.File, userGroupName, userName, fileExtension string) error {
	config.LoadEnv()
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	ctx := context.Background()
	exists, err := r.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		if err := r.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "ap-southeast-1"}); err != nil {
			return err
		}
	}

	fileName := uuid.New().String() + fileExtension
	objectName := filepath.Join(userGroupName, userName, fileName)
	objectName = strings.ReplaceAll(objectName, "\\", "/")

	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, fileName)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return err
	}

	contentType := "application/octet-stream"
	if strings.HasSuffix(fileExtension, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(fileExtension, ".jpg") || strings.HasSuffix(fileExtension, ".jpeg") {
		contentType = "image/jpeg"
	} else if strings.HasSuffix(fileExtension, ".pdf") {
		contentType = "application/pdf"
	}

	_, err = r.client.FPutObject(ctx, bucketName, objectName, tempFilePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}

	defer os.Remove(tempFilePath)

	return nil
}
*/
