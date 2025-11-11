// service/notification_service.go
package service

import (
	"fmt"
	"log"
	"notification_go/entity"
	"notification_go/repository"
	"time"
)

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		repo: repository.NewNotificationRepository(),
	}
}

func (s *NotificationService) CreateNotification(n *entity.Notification) *entity.Notification {
	n.CreatedAt = time.Now()
	n.Status = entity.PENDING
	s.repo.Save(n)
	log.Printf("✅ Notification created: %+v", n)
	return n
}

func (s *NotificationService) GetByID(id int64) *entity.Notification {
	return s.repo.FindByID(id)
}

func (s *NotificationService) List(filter repository.NotificationFilter) []*entity.Notification {
	return s.repo.FindAll(filter)
}

func (s *NotificationService) MyNotifications(username string, channel string) []*entity.Notification {
	var ch entity.Channel
	if channel != "" {
		ch = entity.Channel(channel)
	}
	return s.repo.FindByRecipient(username, ch)
}

func (s *NotificationService) ProcessNotification(n *entity.Notification) {
	defer s.repo.Save(n)
	switch n.Channel {
	case entity.EMAIL:
		// Aquí luego llamas a EmailSender (ejemplo)
		fmt.Printf("📧 Sending email to %s: %s\n", n.Recipient, n.Title)
	case entity.SMS:
		fmt.Printf("📱 Sending SMS to %s: %s\n", n.Recipient, n.Message)
	default:
		fmt.Println("⚠️ Channel not implemented:", n.Channel)
	}

	n.Status = entity.SENT
	now := time.Now()
	n.SentAt = &now
}
