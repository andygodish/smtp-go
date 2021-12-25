package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type server struct {
	host string
	port string
}

func (s *server) address() string {
	return s.host + ":" + s.port
}

func (s *server) sendEmail(emailRecipients []string, message string) {
	from := os.Getenv("GMAIL_ACCOUNT")
	password := os.Getenv(("GMAIL_PASSWORD"))

	ByteMessage := []byte(message)
	auth := smtp.PlainAuth("", from, password, s.host)

	err := smtp.SendMail(s.address(), auth, from, emailRecipients, ByteMessage)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email successfully sent")
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	emailRecipients := []string{"8146597105@vtext.com"}
	emailMessage := "message!"

	smptServer := server{
		host: "smtp.gmail.com",
		port: "587",
	}

	smptServer.sendEmail(emailRecipients, emailMessage)
}
