package repository

import (
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type MediaRepository struct {
	cli *minio.Client
	db  *gorm.DB
}

func NewMediaRepository(cli *minio.Client, db *gorm.DB) ports.Media {
	return &MediaRepository{
		cli: cli,
		db:  db,
	}
}

