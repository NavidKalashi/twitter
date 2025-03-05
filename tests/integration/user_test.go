package integration

// import (
// 	"testing"
// 	"time"

// 	"github.com/NavidKalashi/twitter/internal/core/domain/models"
// 	"github.com/NavidKalashi/twitter/internal/core/service"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockOTPRepo struct {
// 	mock.Mock
// }

// func generateTestToken(email string, secretKey []byte) (string, error) {
//     claims := jwt.MapClaims{
//         "email": email,
//         "exp":   time.Now().Add(time.Hour * 1).Unix(),
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

//     return token.SignedString(secretKey)
// }

// func (mo *MockOTPRepo) Create(user *models.User, code uint) error {
//     args := mo.Called(user, code)
//     return args.Error(0)
// }

// func (mo *MockOTPRepo) FindByUserID(userID string) (*models.OTP, error) {
// 	args := mo.Called(userID)
// 	if args.Get(0) != nil {
// 		return args.Get(0).(*models.OTP), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }

// func (mo *MockOTPRepo) Verified(otp *models.OTP) error {
// 	args := mo.Called(otp)
// 	return args.Error(0)
// }

// func (mo *MockOTPRepo) SendOTP(to string, code uint) error {
// 	args := mo.Called(to, code)
// 	return args.Error(0)
// }

// type MockUserRepo struct {
// 	mock.Mock
// }

// func TestRegsiter(t *testing.T) {
	
// }

// func TestVerify(t *testing.T) {

// 	// arrange
// 	mockOTPRepo := new(MockOTPRepo)
//     us := service.UserService{OTPRepo: mockOTPRepo}

// 	// mock data
// 	userID := uuid.New()
// 	user := &models.User{ID: userID, Email: "test@gmail.com"}
// 	otp := &models.OTP{Code: 123456, CreatedAt: time.Now()}

// 	secretKey := []byte("your_secret_key")
//     tokenString, err := generateTestToken(user.Email, secretKey)
//     assert.NoError(t, err)

// 	// setting up mock
// 	mockOTPRepo.On("FindByUserID", user.ID.String()).Return(otp, nil)
// 	mockOTPRepo.On("Verified", otp).Return(nil)

// 	// act
// 	err = us.Verify(tokenString, user.ID.String(), 123456)

// 	// assert 
// 	assert.NoError(t, err)
// 	mockOTPRepo.AssertExpectations(t)
// }

// func TestGet(t *testing.T){

// }

// func TestUpdate(t *testing.T) {

// }

// func TestDelete(t *testing.T){

// }