// service/notification_consumer.go
package service

import (
    "encoding/json"
    "fmt"
    "log"
    "notification_go/config"
    "notification_go/entity"
)

// StartNotificationConsumer inicia la escucha de la cola
func StartNotificationConsumer(notificationService *NotificationService) {
    msgs, err := config.RabbitChannel.Consume(
        "notification_queue", // Nombre de la cola
        "",                   // consumer
        true,                 // auto-ack
        false,                // exclusive
        false,                // no-local
        false,                // no-wait
        nil,                  // args
    )
    if err != nil {
        log.Fatalf("❌ Error al consumir mensajes: %v", err)
    }

    go func() {
        for msg := range msgs {
            var notif entity.Notification
            if err := json.Unmarshal(msg.Body, &notif); err != nil {
                log.Printf("⚠️ Error al parsear notificación: %v", err)
                continue
            }

            fmt.Printf("📥 Recibida notificación para %s vía %s\n", notif.Recipient, notif.Channel)
            notificationService.ProcessNotification(&notif)
        }
    }()
}
