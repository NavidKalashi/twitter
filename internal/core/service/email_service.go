package service

import (
	"errors"
	"fmt"

	"github.com/NavidKalashi/twitter/internal/config"
	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/resend/resend-go/v2"
)

type EmailService struct {
	emailServ ports.EmailService
	cfg       *config.Config
}

func NewEmailService(emailServ ports.EmailService) *EmailService {
	return &EmailService{emailServ: emailServ}
}

func (es *EmailService) SendOTP(to string, code uint) error {
	apiKey := es.cfg.Apikey

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"kalashinavid@gmail.com"},
		Html:    "<strong>your verification code: </strong>" + fmt.Sprintf("%d", code),
		Subject: "hello from twitter",
	}

	claims, err := client.Emails.Send(params)
	if err != nil {
		return err
	}
	if claims != nil {
		return errors.New("email not sent")
	}

	return nil
}
