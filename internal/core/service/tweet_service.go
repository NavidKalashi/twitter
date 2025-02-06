package service

import "database/sql"

type TweetService struct {
	db *sql.DB
}

func NewTweetService(db *sql.DB) *TweetService {
	return &TweetService{db: db}
}