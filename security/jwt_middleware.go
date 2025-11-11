package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware(jwtUtil *JwtUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.Next()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		username, err := jwtUtil.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
