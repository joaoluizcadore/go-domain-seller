package service

import (
	"fmt"

	"github.com/joaoluizcadore/domain-seller/config"
	"github.com/joaoluizcadore/domain-seller/infra"
	"github.com/joaoluizcadore/domain-seller/internal/shared"
)

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) SendNotification(msg shared.Message) error {
	return infra.SendEmail(config.GetConfig().SendNotificationTo, "Domain Seller: New Contact", buildMessage(msg))
}

func buildMessage(msg shared.Message) string {
	return fmt.Sprintf(
		"New contact from Domain Seller\n\n"+
			"Domain:  %s\n"+
			"Name:  %s\n"+
			"Email: %s\n"+
			"Phone: %s\n"+
			"Message: %s\n", msg.Domain, msg.Name, msg.Email, msg.Phone, msg.Message)
}
