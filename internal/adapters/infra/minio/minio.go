package minio

import (
	"context"
	"fmt"
	"log"

	"github.com/NavidKalashi/twitter/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	mc     *minio.Client
	bucket string
}

func InitMinio(cfg *config.Config) (*Minio, error) {
	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("connected to minio successfully")

	ctx := context.Background()
	bucketName := cfg.Minio.BucketName

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket: %w", err)
		}
		fmt.Println("Bucket created:", bucketName)
	} else {
		fmt.Println("Bucket already exists:", bucketName)
	}

	return &Minio{mc: minioClient, bucket: bucketName}, nil
}

func (m *Minio) GetMinio() *minio.Client {
	return m.mc
}