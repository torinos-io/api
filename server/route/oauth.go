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

// GetAuthentication authenticate user
func GetAuthentication(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Authenticate",
	})
}
