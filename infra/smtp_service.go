package infra

import (
	"fmt"
	"net/smtp"

	"github.com/joaoluizcadore/domain-seller/config"
)

func SendNotification(msg string) error {
	config := config.NewConfig()
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)

	host := fmt.Sprintf("%v:%v", config.SMTPHost, config.SMTPPort)

	body := []byte(fmt.Sprintf("To: %v\r\nSubject: Domain Seller - New Notification\r\n\r\n%v\r\n", config.SendNotificationTo, msg))

	return smtp.SendMail(host, auth, config.SMTPUsername, []string{config.SendNotificationTo}, body)
}
