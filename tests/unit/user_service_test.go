package unit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) Send(to string, message string) error {
	args := m.Called(to, message)
	return args.Error(0)
}

func (m *MockEmailService) SendOTP(to string, code uint) error {
	args := m.Called(to, code)
	return args.Error(0)
}

func TestSend_Success(t *testing.T) {
	mockEmailService := new(MockEmailService)

	mockEmailService.On("Send", "kalashinavid@gmail.com", "Hello!").Return(nil)

	err := mockEmailService.Send("kalashinavid@gmail.com", "Hello!")

	assert.NoError(t, err)

	mockEmailService.AssertCalled(t, "Send", "kalashinavid@gmail.com", "Hello!")
}

func TestSend_Error(t *testing.T) {
	mockEmailService := new(MockEmailService)

	mockEmailService.On("Send", "kalashinavid@gmail.com", "Hello!").Return(errors.New("failed to send email"))

	err := mockEmailService.Send("kalashinavid@gmail.com", "Hello!")

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to send email")

	mockEmailService.AssertCalled(t, "Send", "kalashinavid@gmail.com", "Hello!")
}

func TestSendOTP_Success(t *testing.T) {
	mockEmailService := new(MockEmailService)

	mockEmailService.On("SendOTP", "kalashinavid@gmail.com", uint(123456)).Return(nil)

	err := mockEmailService.SendOTP("kalashinavid@gmail.com", 123456)

	assert.NoError(t, err)

	mockEmailService.AssertCalled(t, "SendOTP", "kalashinavid@gmail.com", uint(123456))
}

func TestSendOTP_Error(t *testing.T) {
	mockEmailService := new(MockEmailService)

	mockEmailService.On("SendOTP", "kalashinavid@gmail.com", uint(123456)).Return(errors.New("failed to send OTP"))

	err := mockEmailService.SendOTP("kalashinavid@gmail.com", 123456)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to send OTP")

	mockEmailService.AssertCalled(t, "SendOTP", "kalashinavid@gmail.com", uint(123456))
}
