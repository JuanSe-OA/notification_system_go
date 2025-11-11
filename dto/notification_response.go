// dto/notification_response.go
package dto

import (
	"notification_go/entity"
	"time"
)

type NotificationResponse struct {
	ID            int64               `json:"id"`
	Title         string              `json:"title"`
	Message       string              `json:"message"`
	Recipient     string              `json:"recipient"`
	Channel       entity.Channel      `json:"channel"`
	Status        entity.NotificationStatus `json:"status"`
	CreatedAt     time.Time           `json:"createdAt"`
	ScheduledAt   *time.Time          `json:"scheduledAt"`
	SentAt        *time.Time          `json:"sentAt"`
	FailureReason string              `json:"failureReason"`
}
