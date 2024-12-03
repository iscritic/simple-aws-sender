package service

import (
	"errors"
	"github.com/iscritic/simple-aws-sender/internal/repository"
)

type SMTPService interface {
	SendEmail(emailRequest EmailRequest) error
}

type smtpService struct {
	smtpRepo repository.SMTPRepository
}

func NewSMTPService(r repository.SMTPRepository) SMTPService {
	return &smtpService{smtpRepo: r}
}

type EmailRequest struct {
	To      string `json:"to" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func (s *smtpService) SendEmail(request EmailRequest) error {
	if request.To == "" || request.Subject == "" || request.Body == "" {
		return errors.New("invalid email request")
	}
	return s.smtpRepo.SendEmail(request.To, request.Subject, request.Body)
}
