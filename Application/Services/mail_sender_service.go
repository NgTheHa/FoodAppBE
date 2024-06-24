package Services

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
)

type MailSettings struct {
	Mail        string `json:"mail"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
}

type MailContent struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// EmailSenderService defines methods for sending emails
type EmailSenderService interface {
	SendMail(mailContent *MailContent) error
}

// EmailSenderServiceImpl implements EmailSenderService
type EmailSenderServiceImpl struct {
	mailSettings *MailSettings
}

// NewEmailSenderService creates a new instance of EmailSenderService
func NewEmailSenderService(mailSettings *MailSettings) EmailSenderService {
	return &EmailSenderServiceImpl{mailSettings: mailSettings}
}

// SendMail sends an email
func (e *EmailSenderServiceImpl) SendMail(mailContent *MailContent) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.mailSettings.Mail)
	m.SetHeader("To", mailContent.To)
	m.SetHeader("Subject", mailContent.Subject)
	m.SetBody("text/html", mailContent.Body)

	d := gomail.NewDialer(e.mailSettings.Host, e.mailSettings.Port, e.mailSettings.Mail, e.mailSettings.Password)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		// Save email to file in case of failure
		saveEmailToFile(mailContent)
		return err
	}

	return nil
}

// Helper function to save email to file in case of failure
func saveEmailToFile(mailContent *MailContent) {
	dir := "mailsaves"
	_ = os.MkdirAll(dir, os.ModePerm)

	// Generate unique filename
	filename := uuid.NewV4().String() + ".eml"
	filePath := dir + "/" + filename

	// Save email content to file
	f, err := os.Create(filePath)
	if err != nil {
		logrus.WithError(err).Error("Failed to save email to file")
		return
	}
	defer f.Close()

	_, err = f.WriteString("To: " + mailContent.To + "\n")
	if err != nil {
		logrus.WithError(err).Error("Failed to write email content to file")
		return
	}
	_, err = f.WriteString("Subject: " + mailContent.Subject + "\n\n")
	if err != nil {
		logrus.WithError(err).Error("Failed to write email content to file")
		return
	}
	_, err = f.WriteString(mailContent.Body)
	if err != nil {
		logrus.WithError(err).Error("Failed to write email content to file")
		return
	}

	logrus.Info("Email saved to file:", filePath)
}

// HandleEmail sends an email using JSON request body
func HandleEmail(c *gin.Context) {
	var mailContent MailContent
	if err := c.ShouldBindBodyWith(&mailContent, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize Email Sender Service
	emailService := NewEmailSenderService(&MailSettings{
		Mail:     "your-email@example.com",
		Host:     "smtp.example.com",
		Port:     587,
		Password: "your-email-password",
	})

	// Send email
	if err := emailService.SendMail(&mailContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
