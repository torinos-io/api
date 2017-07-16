package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Subscribe subscription
func Subscribe(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "Subscribed",
		"uuid":    uuid,
	})
}

// Unsubscribe subscription
func Unsubscribe(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "Unsubscribed",
		"uuid":    uuid,
	})
}
