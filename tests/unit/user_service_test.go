package unit

import (
	"testing"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/domain/models"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOTPRepo struct {
	mock.Mock
}

func (mo *MockOTPRepo) Create(user *models.User, code uint) error {
	args := mo.Called(user, code)
	return args.Error(0)
}

func (mo *MockOTPRepo) FindByUserID(userID string) (*models.OTP, error) {
	args := mo.Called(userID)
	if args.Get(0) != nil {
		return args.Get(0).(*models.OTP), args.Error(1)
	}
	return nil, args.Error(1)
}

func (mo *MockOTPRepo) Verified(otp *models.OTP) error {
	args := mo.Called(otp)
	return args.Error(0)
}

func (mo *MockOTPRepo) SendOTP(to string, code uint) error {
	args := mo.Called(to, code)
	return args.Error(0)
}

type MockRefreshTokenRepo struct {
	mock.Mock
}

func (mo *MockRefreshTokenRepo) Create(userID uuid.UUID, refreshToken string) error {
	args := mo.Called(userID, refreshToken)
	return args.Error(0)
}

type MockAccessTokenRepo struct {
	mock.Mock
}

func (mo *MockAccessTokenRepo) Set(userID string, accessToken string) error {
	args := mo.Called(userID, accessToken)
	return args.Error(0)
}

type UserService struct {
	OTPRepo          ports.OTP
	RefreshTokenRepo ports.RefreshToken
	AccessTokenRepo  ports.AccessToken
	SecretKey        []byte 
}

func generateTestToken(email string, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func generateTestAccessAndRefresh(userID string, secretKey []byte) (string, string, error) {
	// access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return refreshTokenString, accessTokenString, nil
}

func (us *UserService) Verify(tokenString, userIDString string, code uint) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return us.SecretKey, nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return "", "", err
	}

	otp, err := us.OTPRepo.FindByUserID(userID.String())
	if err != nil {
		return "", "", err
	}

	if otp.Code != code {
		return "", "", err
	}

	err = us.OTPRepo.Verified(otp)
	if err != nil {
		return "", "", err
	}

	refreshToken, accessToken, err := generateTestAccessAndRefresh(userID.String(), us.SecretKey)
	if err != nil {
		return "", "", err
	}

	err = us.RefreshTokenRepo.Create(userID, refreshToken)
	if err != nil {
		return "", "", err
	}

	err = us.AccessTokenRepo.Set(userID.String(), accessToken)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func TestVerify(t *testing.T) {

	// arrange
	mockOTPRepo := new(MockOTPRepo)
	mockRefreshTokenRepo := new(MockRefreshTokenRepo)
	mockAccessTokenRepo := new(MockAccessTokenRepo)
	secretKey := []byte("your_secret_key") 

	us := UserService{
		OTPRepo:          mockOTPRepo,
		RefreshTokenRepo: mockRefreshTokenRepo,
		AccessTokenRepo:  mockAccessTokenRepo,
		SecretKey:        secretKey,
	}

	// mock data
	userID := uuid.New()
	user := &models.User{ID: userID, Email: "test@gmail.com"}
	otp := &models.OTP{Code: 123456, CreatedAt: time.Now()}

	tokenString, err := generateTestToken(user.Email, secretKey)
	assert.NoError(t, err)

	refreshToken, accessToken, err := generateTestAccessAndRefresh(user.ID.String(), secretKey)
	assert.NoError(t, err)

	// setting up mock
	mockOTPRepo.On("FindByUserID", user.ID.String()).Return(otp, nil)
	mockOTPRepo.On("Verified", otp).Return(nil)
	mockRefreshTokenRepo.On("Create", user.ID, refreshToken).Return(nil)
	mockAccessTokenRepo.On("Set", user.ID.String(), accessToken).Return(nil)

	// act
	returnedRefreshToken, returnedAccessToken, err := us.Verify(tokenString, user.ID.String(), 123456)
	assert.NoError(t, err)

	// assert
	assert.Equal(t, refreshToken, returnedRefreshToken)
	assert.Equal(t, accessToken, returnedAccessToken)

	mockOTPRepo.AssertExpectations(t)
	mockRefreshTokenRepo.AssertExpectations(t)
	mockAccessTokenRepo.AssertExpectations(t)
}
