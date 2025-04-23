package ports

import (
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/events"
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
	Get(userID uuid.UUID) (*models.RefreshToken, error)
	Delete(userID uuid.UUID) error
}

type Tweet interface {
	Create(tweet *models.Tweet) error
	DeleteAll(username string) error
	Delete(username, tweetID string) error
	GetTweets() ([]models.Tweet, error)
	GetByID(tweetID string) (*models.Tweet, error)
	Update(tweet *models.Tweet) error
	GetByUsername(username string) (*models.Tweet, error)
}

type Media interface {
	SaveMedia(media *models.Media) error
}

type Storage interface {
	UploadMedia(filePath string) (string, error)
}

type Gesture interface {
	Save(gesture *models.Gesture) error
	Set(tweetID, gestureType, username string) error
	Count(tweetID, username string) (int, error)
	GetByUsername(tweetID, username string) (*models.Gesture, error)
	Exists(tweetID, username, typeStr string) (bool, error)
}

type Consume interface {
	ConsumeFeedEvents(handler func(events.Feed)) error
	ConsumeGestureEvents(handler func(events.Gesture)) error
}

type Producer interface {
	ProducerFeedEvents(feed events.Feed) error
	ProducerGestureEvents(createdGesture events.Gesture) error
}

type Follow interface {
	Save(follow *models.Follow) error
	Delete(followerName, followingName string) error
	GetFollowers(username string) ([]models.Follow, error)
	GetFollowing(username string) ([]models.Follow, error)
}

type Feed interface {
	Set(username, tweet string) error
	Get(username string) ([]string, error)
}