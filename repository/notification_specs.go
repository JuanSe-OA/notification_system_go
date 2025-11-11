// repository/notification_specs.go
package repository

import (
	"notification_go/entity"
	"strings"
	"time"
)

type NotificationFilter struct {
	Recipient string
	Status    entity.NotificationStatus
	Channel   entity.Channel
	FromDate  *time.Time
	ToDate    *time.Time
	Query     string
}
