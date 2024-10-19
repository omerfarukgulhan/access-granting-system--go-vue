package service

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
	BaseURL  string
}

var emailConfig *EmailConfig

func init() {
	var err error
	emailConfig, err = loadEnv()
	if err != nil {
		log.Fatalf("Failed to load email configuration: %v", err)
	}
}

func loadEnv() (*EmailConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	emailConfig := &EmailConfig{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		From:     os.Getenv("EMAIL_FROM"),
		BaseURL:  os.Getenv("CLIENT_URL"),
	}

	if emailConfig.Host == "" || emailConfig.Port == "" || emailConfig.Username == "" ||
		emailConfig.Password == "" || emailConfig.From == "" || emailConfig.BaseURL == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return emailConfig, nil
}

func SendActivationEmail(recipientEmail, activationToken string) error {
	htmlBody := fmt.Sprintf(`
		<html>
			<body>
				<h1>Activate your account</h1>
				<p>Click the <a href="%s/activate-user/%s">link</a> to activate your account: </p>
			</body>
		</html>
	`, emailConfig.BaseURL, activationToken)

	return sendEmail(recipientEmail, "Activation Email", htmlBody)
}

func SendPasswordResetEmail(recipientEmail, resetToken string) error {
	htmlBody := fmt.Sprintf(`
		<html>
			<body>
				<h1>Reset your password</h1>
				<p>Click the <a href="%s/reset-password/%s">link</a> to reset your password.</p>
				<p>If you did not request a password reset, please ignore this email.</p>
			</body>
		</html>
	`, emailConfig.BaseURL, resetToken)

	return sendEmail(recipientEmail, "Password Reset Request", htmlBody)
}

func createEmailMessage(to, subject, body string) []byte {
	toHeader := fmt.Sprintf("To: %s\n", to)
	subjectHeader := fmt.Sprintf("Subject: %s\n", subject)

	return []byte(
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n" +
			toHeader + subjectHeader + "\n" + body,
	)
}

func sendEmail(recipientEmail, subject, body string) error {
	auth := smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.Host)
	to := []string{recipientEmail}
	msg := createEmailMessage(recipientEmail, subject, body)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", emailConfig.Host, emailConfig.Port),
		auth,
		emailConfig.From,
		to,
		msg,
	)

	if err != nil {
		log.Printf("Error sending email to %s: %v", recipientEmail, err)
		return fmt.Errorf("failed to send email to %s: %v", recipientEmail, err)
	}

	return nil
}
