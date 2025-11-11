// entity/notification_status.go
package entity

type NotificationStatus string

const (
    Scheduled NotificationStatus = "SCHEDULED"
    Pending   NotificationStatus = "PENDING"
    Sent      NotificationStatus = "SENT"
    Failed    NotificationStatus = "FAILED"
)
