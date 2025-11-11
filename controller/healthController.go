// controller/health_controller.go
package controller

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

const Version = "notification-service-1.0.0"

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, buildHealthResponse("UP", "Service is running normally"))
	})
	r.GET("/health/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, buildHealthResponse("UP", "Service is alive"))
	})
	r.GET("/health/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, buildHealthResponse("UP", "Service is ready to handle requests"))
	})
}

func buildHealthResponse(status, description string) gin.H {
	uptime := time.Since(startTime)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return gin.H{
		"status":      status,
		"description": description,
		"version":     Version,
		"uptime":      uptime.String(),
		"timestamp":   time.Now().Format(time.RFC3339),
		"memory_mb":   m.Alloc / 1024 / 1024,
	}
}
