package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/iscritic/simple-aws-sender/internal/service"
	"net/http"
)

type SMTPHandler struct {
	smtpService service.SMTPService
}

func NewSMTPHandler(s service.SMTPService) *SMTPHandler {
	return &SMTPHandler{smtpService: s}
}

func (h *SMTPHandler) SendEmail(c *gin.Context) {
	var request service.EmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.smtpService.SendEmail(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
