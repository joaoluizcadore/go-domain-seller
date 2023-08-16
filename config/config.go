package config

import (
	"log"
	"os"
	"strconv"
	"sync"
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

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {

	if instance == nil {
		once.Do(func() {
			smtPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
			if err != nil {
				log.Fatalf("Invalid SMT_PORT configuration: %v", err)
				panic(err)
			}

			serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
			if err != nil {
				log.Fatalf("Invalid SERVER_PORT configuration: %v", err)
				panic(err)
			}

			instance = &Config{
				SMTPHost:           os.Getenv("SMTP_HOST"),
				SMTPPort:           smtPort,
				SMTPUsername:       os.Getenv("SMTP_USERNAME"),
				SMTPPassword:       os.Getenv("SMTP_PASSWORD"),
				SMTPTLS:            os.Getenv("SMTP_TLS") == "true",
				SendNotificationTo: os.Getenv("SEND_NOTIFICATION_TO"),
				ServerPort:         serverPort,
			}
		})
	}
	return instance
}
