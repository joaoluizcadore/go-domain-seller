package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	SMTPHost           string
	SMTPPort           int
	SMTPUsername       string
	SMTPPassword       string
	SMTPTLS            bool
	SendNotificationTo string //Who will receive emails (email address)
	ServerPort         int
}

func NewConfig() *Config {
	smtPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMT_PORT configuration: %v", err)
		panic(err)
	}

	serverPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SERVER_PORT configuration: %v", err)
		panic(err)
	}

	return &Config{
		SMTPHost:           os.Getenv("SMTP_HOST"),
		SMTPPort:           smtPort,
		SMTPUsername:       os.Getenv("SMTP_USERNAME"),
		SMTPPassword:       os.Getenv("SMTP_PASSWORD"),
		SMTPTLS:            os.Getenv("SMTP_TLS") == "true",
		SendNotificationTo: os.Getenv("SEND_NOTIFICATION_TO"),
		ServerPort:         serverPort,
	}
}
