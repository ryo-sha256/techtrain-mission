package Utils

import (
	"encoding/hex"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken() string {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// Extract token from the header
func ExtractAuthToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Authorization token is empty",
		})
		c.Abort()
		return ""
	}

	return authHeader
}
