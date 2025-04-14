package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/minio/minio-go/v7"
)

type MinioStorage struct {
	client *minio.Client
}

func NewMinioStorage(client *minio.Client) ports.Storage {
	return &MinioStorage{client: client}
}

func (ms *MinioStorage) UploadMedia(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	objectName := "uploads/" + filePath
	_, err = ms.client.PutObject(context.TODO(), "testbucket", objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to minio: %v", err)
	}

	fileURL := fmt.Sprintf("http://%s/%s/%s", "localhost:9000", "testbucket", objectName)
	return fileURL, nil
}
