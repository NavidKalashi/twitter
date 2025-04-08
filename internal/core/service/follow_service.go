package service

import (
	"fmt"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
)

type FollowService struct {
	followRepo ports.Follow
	userRepo ports.User
}

func NewFollowService(followRepo ports.Follow, userRepo ports.User) *FollowService {
	return &FollowService{followRepo: followRepo, userRepo: userRepo}
}

func (fs *FollowService) FollowUser(followerName string, followingName string) error {
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