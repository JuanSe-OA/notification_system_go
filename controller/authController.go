package controller

import (
	"net/http"
	"notification_go/dto"
	"notification_go/security"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, jwtUtil *security.JwtUtil) {
	auth := r.Group("/api/auth")
	auth.POST("/login", func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Dummy validation (puedes reemplazar con DB o LDAP)
		if req.Username != "admin" || req.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
			return
		}

		token, _ := jwtUtil.GenerateToken(req.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
