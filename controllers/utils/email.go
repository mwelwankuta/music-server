package utils

import (
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendEmail(email []string, message string) error {
	godotenv.Load("../../.env")
	EMAIL_FROM := os.Getenv("EMAIL_FROM")
	EMAIL_PASSWORD := os.Getenv("EMAIL_PASSWORD")

	host := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("", EMAIL_FROM, EMAIL_PASSWORD, host)
	var err = smtp.SendMail(host+":"+port, auth, EMAIL_FROM, email, []byte(message))

	return err
}
