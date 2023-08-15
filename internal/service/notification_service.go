package service

import (
	"github.com/joaoluizcadore/domain-seller/config"
	"github.com/joaoluizcadore/domain-seller/infra"
)

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) SendNotification(text string) error {
	return infra.SendEmail(config.GetConfig().SendNotificationTo, "Domain Seller: New Contact", text)
}
