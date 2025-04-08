package ports

import (
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/google/uuid"
)

type User interface {
	Register(username, name, email, hashPass, bio string, birthday time.Time) error
	GetByEmail(email string) (*models.User, error)
	GetByID(userID string) (*models.User, error)
	GetByName(username string) (*models.User, error)
	Edit(user *models.User) error
	EmailExist(email string) (*models.User, error)
	UsernameExist(username string) (*models.User, error)
	Verified(user *models.User, sit bool) error
	Search(username string) (*models.User, error)
}

type EmailService interface {
	Send(to string, message string) error
	SendOTP(to string, code uint) error
}

type OTP interface {
	Set(email string, code uint) error
	Get(email string) (uint, error)
}

type RefreshToken interface {
	Create(userID uuid.UUID, refreshToken string) error
	Get(userID uuid.UUID) error
	Delete(userID uuid.UUID) error
}

type Tweet interface {
	Create(tweet *models.Tweet) error
	DeleteAll(username string) error
	Delete(username, tweetID string) error
	GetTweets() ([]models.Tweet, error)
	GetByID(tweetID string) (*models.Tweet, error)
	Update(tweet *models.Tweet) error
}

type Media interface {
	SaveMedia(media *models.Media) error
}

type Storage interface {
	UploadMedia(filePath string) (string, error)
}

type Gesture interface {
	Save(gesture *models.Gesture) error
	Count(tweetID, username string) (int, error)
	GetByUsername(tweetID, username string) (*models.Gesture, error)
}

type Follow interface {
	Save(follow *models.Follow) error
	GetFollowers(username string) ([]models.Follow, error)
	Delete(followerName, followingName string) error
	// GetFollowing(username string) ([]models.Follow, error)
	// GetByFollowerAndFollowing(followerName, followingName string) (*models.Follow, error)
	// Delete(followerName, followingName string) error
	// GetFollowerCount(username string) (int, error)
	// GetFollowingCount(username string) (int, error)
}