package auth

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	USERNAME = "admin"
	PASSWORD = "password"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authHeader = strings.TrimPrefix(authHeader, "Basic")

		decoded, err := base64.StdEncoding.DecodeString(authHeader)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(decoded), ":", 2)

		if len(credentials) != 2 || credentials[0] != USERNAME || credentials[1] != PASSWORD {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			c.Abort()
			return
		}

		c.Next()
	}
}
