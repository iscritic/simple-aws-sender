package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
)

type SMTPRepository interface {
	SendEmail(to, subject, body string) error
}

type smtpRepository struct {
	sesClient   *ses.SES
	senderEmail string
}

func NewSMTPRepository() (SMTPRepository, error) {

	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	senderEmail := os.Getenv("AWS_SENDER_EMAIL")
	region := os.Getenv("AWS_SENDER_REGION")

	if awsSecretAccessKey == "" || awsAccessKeyID == "" || senderEmail == "" || region == "" {
		return nil, fmt.Errorf("one or more AWS environment variables are not set")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	sesClient := ses.New(sess)

	return &smtpRepository{
		sesClient:   sesClient,
		senderEmail: senderEmail,
	}, nil
}

func (r *smtpRepository) SendEmail(to, subject, body string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(to)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(r.senderEmail),
	}

	_, err := r.sesClient.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}
