// repository/notification_repository.go
package repository

import (
	"notification-service/entity"
	"strings"
	"sync"
	"time"
)

type NotificationRepository struct {
	data map[int64]*entity.Notification
	mu   sync.RWMutex
	id   int64
}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{
		data: make(map[int64]*entity.Notification),
	}
}

func (r *NotificationRepository) Save(n *entity.Notification) *entity.Notification {
	r.mu.Lock()
	defer r.mu.Unlock()
	if n.ID == 0 {
		r.id++
		n.ID = r.id
	}
	r.data[n.ID] = n
	return n
}

func (r *NotificationRepository) FindByID(id int64) *entity.Notification {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.data[id]
}

func (r *NotificationRepository) FindAll(filter NotificationFilter) []*entity.Notification {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Notification
	for _, n := range r.data {
		if filter.Recipient != "" && !strings.EqualFold(n.Recipient, filter.Recipient) {
			continue
		}
		if filter.Status != "" && n.Status != filter.Status {
			continue
		}
		if filter.Channel != "" && n.Channel != filter.Channel {
			continue
		}
		if filter.FromDate != nil && n.CreatedAt.Before(*filter.FromDate) {
			continue
		}
		if filter.ToDate != nil && n.CreatedAt.After(*filter.ToDate) {
			continue
		}
		if filter.Query != "" && !strings.Contains(strings.ToLower(n.Message+n.Title), strings.ToLower(filter.Query)) {
			continue
		}
		result = append(result, n)
	}
	return result
}

func (r *NotificationRepository) FindByRecipient(recipient string, channel entity.Channel) []*entity.Notification {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Notification
	for _, n := range r.data {
		if strings.EqualFold(n.Recipient, recipient) && (channel == "" || n.Channel == channel) {
			result = append(result, n)
		}
	}
	return result
}
