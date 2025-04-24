package service

import (
	"log"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (e *EmailService) SendEmail(email string) error {
	// Simulate sending email
	log.Printf("Email sent to %s\n", email)
	return nil
}
