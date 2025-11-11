// entity/notification.go
package entity

import "time"

type Channel string
type NotificationStatus string

const (
	EMAIL    Channel = "EMAIL"
	SMS      Channel = "SMS"
	WHATSAPP Channel = "WHATSAPP"
	PUSH     Channel = "PUSH"
)

const (
	SCHEDULED NotificationStatus = "SCHEDULED"
	PENDING   NotificationStatus = "PENDING"
	SENT      NotificationStatus = "SENT"
	FAILED    NotificationStatus = "FAILED"
)

type Notification struct {
	ID            int64              `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string             `json:"title"`
	Message       string             `json:"message"`
	Recipient     string             `json:"recipient" gorm:"index"`
	Channel       Channel            `json:"channel"`
	Status        NotificationStatus `json:"status"`
	ScheduledAt   *time.Time         `json:"scheduledAt"`
	CreatedAt     time.Time          `json:"createdAt"`
	SentAt        *time.Time         `json:"sentAt"`
	FailureReason string             `json:"failureReason"`
}
