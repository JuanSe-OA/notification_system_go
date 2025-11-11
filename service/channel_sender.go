// service/channel_sender.go
package service

import "notification-service/entity"

type ChannelSender interface {
    Send(notification *entity.Notification) error
}
