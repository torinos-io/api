package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser return the current user
func GetCurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetCurrentUser",
		"user":    "yamada",
	})
}
