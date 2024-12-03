package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iscritic/simple-aws-sender/internal/delivery"
	"github.com/iscritic/simple-aws-sender/internal/repository"
	"github.com/iscritic/simple-aws-sender/internal/service"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	smtpRepo, err := repository.NewSMTPRepository()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	smtpService := service.NewSMTPService(smtpRepo)

	smtpHandler := delivery.NewSMTPHandler(smtpService)

	r.POST("/send-email", smtpHandler.SendEmail)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
