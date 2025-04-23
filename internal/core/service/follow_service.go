package service

import (
	"fmt"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type FollowService struct {
	followRepo ports.Follow
	userRepo   ports.User
	tweetRepo  ports.Tweet
	feedRepo   ports.Feed
}

func NewFollowService(followRepo ports.Follow, userRepo ports.User, tweetRepo ports.Tweet, feedRepo ports.Feed) *FollowService {
	return &FollowService{
		followRepo: followRepo, 
		userRepo: userRepo, 
		tweetRepo: tweetRepo, 
		feedRepo: feedRepo,
	}
}

func (fs *FollowService) FollowUser(followerName, followingName string) error {
	if followerName == followingName {
		return fmt.Errorf("you cannot follow yourself")
	}

	user, err := fs.userRepo.GetByName(followingName)
	if err != nil {
		return fmt.Errorf("user does not exist")
	}

	exists, err := fs.followRepo.Exists(followerName, followingName)
	if err != nil {
		return err
	}
	if exists {
		return nil // یا return fmt.Errorf("already following this user")
	}

	follow := &models.Follow{
		FollowerName:  followerName,
		FollowingName: user.Username,
	}

	return fs.followRepo.Save(follow)
}

func (fs *FollowService) UnfollowUser(followerName, followingName string) error {
	user, err := fs.userRepo.GetByName(followingName)
	if err != nil {
		return fmt.Errorf("user does not exist")
	}

	return fs.followRepo.Delete(followerName, user.Username)
}

func (fs *FollowService) GetFollowers(username string) ([]string, error) {
	follows, err := fs.followRepo.GetFollowers(username)
	if err != nil {
		return nil, fmt.Errorf("you don't have any followers")
	}

	var followers []string
	for _, follow := range follows {
		followers = append(followers, follow.FollowerName)
	}

	return followers, nil
}

func (fs *FollowService) GetFollowing(username string) ([]string, error) {
	follows, err := fs.followRepo.GetFollowing(username)
	if err != nil {
		return nil, fmt.Errorf("you don't have any following")
	}

	var followings []string
	for _, follow := range follows {
		followings = append(followings, follow.FollowingName)
	}

	return followings, nil
}