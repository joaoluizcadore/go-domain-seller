package infra

import (
	"fmt"
	"net/smtp"

	"github.com/joaoluizcadore/domain-seller/config"
)

func SendEmail(to string, title string, text string) error {
	cfg := config.GetConfig()

	auth := smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPHost)

	host := fmt.Sprintf("%v:%v", cfg.SMTPHost, cfg.SMTPPort)

	body := []byte(fmt.Sprintf("To: %v\r\nSubject: %v\r\n\r\n%v\r\n", to, title, text))

	return smtp.SendMail(host, auth, cfg.SMTPUsername, []string{to}, body)
}
