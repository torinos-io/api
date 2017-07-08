package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAuthorization authorize user
func GetAuthorization(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Authorize",
	})
}

// Authenticate user
func Authenticate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Authenticate",
	})
}
