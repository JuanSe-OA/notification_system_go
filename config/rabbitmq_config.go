// config/rabbitmq_config.go
package config

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

const (
	NotificationQueue      = "notification.queue"
	NotificationExchange   = "notification.exchange"
	NotificationRoutingKey = "notification.routingkey"
)

func SetupRabbitMQ(uri string) (*amqp091.Connection, *amqp091.Channel) {
	conn, err := amqp091.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	// Declare queue, exchange, and binding
	_, _ = ch.QueueDeclare(NotificationQueue, true, false, false, false, nil)
	_ = ch.ExchangeDeclare(NotificationExchange, "direct", true, false, false, false, nil)
	_ = ch.QueueBind(NotificationQueue, NotificationRoutingKey, NotificationExchange, false, nil)

	log.Println("✅ RabbitMQ initialized")
	return conn, ch
}
