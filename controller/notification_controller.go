// controller/notification_controller.go
package controller

import (
	"net/http"
	"notification_go/dto"
	"notification_go/entity"
	"notification_go/service"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(r *gin.Engine, svc *service.NotificationService) {
	api := r.Group("/api/notifications")

	api.GET("/channels", func(c *gin.Context) {
		channels := []string{"EMAIL", "SMS", "WHATSAPP"}
		c.JSON(http.StatusOK, channels)
	})

	api.POST("", func(c *gin.Context) {
		var n entity.Notification
		if err := c.ShouldBindJSON(&n); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		created := svc.CreateNotification(&n)
		c.JSON(http.StatusOK, created)
	})

	api.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		n := svc.GetByID(id)
		if n == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		}
		c.JSON(http.StatusOK, n)
	})

	api.GET("", func(c *gin.Context) {
		filter := service.NotificationFilter{
			Recipient: c.Query("recipient"),
			Status:    entity.NotificationStatus(c.Query("status")),
			Channel:   entity.Channel(c.Query("channel")),
			Query:     c.Query("q"),
		}
		if from := c.Query("from"); from != "" {
			if t, err := time.Parse(time.RFC3339, from); err == nil {
				filter.FromDate = &t
			}
		}
		if to := c.Query("to"); to != "" {
			if t, err := time.Parse(time.RFC3339, to); err == nil {
				filter.ToDate = &t
			}
		}

		result := svc.List(filter)
		c.JSON(http.StatusOK, result)
	})

	api.GET("/me", func(c *gin.Context) {
		username, ok := c.Get("username")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		channel := c.Query("channel")
		list := svc.MyNotifications(username.(string), strings.ToUpper(channel))
		c.JSON(http.StatusOK, list)
	})
}
