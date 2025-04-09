package service

import (
	"fmt"
	"log"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type FollowService struct {
	followRepo ports.Follow
	userRepo   ports.User
	tweetRepo  ports.Tweet
}

func NewFollowService(followRepo ports.Follow, userRepo ports.User, tweetRepo ports.Tweet) *FollowService {
	return &FollowService{followRepo: followRepo, userRepo: userRepo, tweetRepo: tweetRepo}
}

func (fs *FollowService) FollowUser(followerName, followingName string) error {
	user, err := fs.userRepo.GetByName(followingName)
	if err != nil {
		return fmt.Errorf("user does not exist")
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

func (fs *FollowService) Feed(username string) ([]models.Tweet, error) {
	followings, err := fs.GetFollowing(username)
	if err != nil {
		return nil, err
	}
	log.Println("followers:", followings, username)

	var feeds []models.Tweet
	for _, following := range followings {
		tweets, err := fs.tweetRepo.GetByUsername(following)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, tweets...)
	}

	return feeds, nil
}
