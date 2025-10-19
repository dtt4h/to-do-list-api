package middleware

import (
	"net/http"
	"strings"
	"to-do-list-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth header format must be Bearer {token}"})
			c.Abort()
			return
		}
		tokenString := parts[1]

		// validation token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token:" + err.Error()})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
