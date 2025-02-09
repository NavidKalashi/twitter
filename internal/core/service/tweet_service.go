package service

import "gorm.io/gorm"

type TweetService struct {
	db *gorm.DB
}

func NewTweetService(db *gorm.DB) *TweetService {
	return &TweetService{db: db}
}