package main

import (
	"log"
	"notification_go/config"
	"notification_go/controller"
	"notification_go/security"
	"notification_go/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	jwtUtil := &security.JwtUtil{
		Secret:       "supersecretkey12345",
		ExpirationMs: 3600000,
	}

	r.Use(security.JwtMiddleware(jwtUtil))

	// RabbitMQ
	config.SetupRabbitMQ("amqp://guest:guest@localhost:5672/")

	// Servicios
	notifService := service.NewNotificationService()

	// Rutas
	controller.RegisterHealthRoutes(r)
	controller.RegisterAuthRoutes(r, jwtUtil)
	controller.RegisterNotificationRoutes(r, notifService)

	log.Println("🚀 Running on :8080")
	r.Run(":8080")
}
