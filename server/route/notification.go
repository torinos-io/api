package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNotification create notification of project
func CreateNotification(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification created",
		"uuid":    uuid,
	})
}

// DeleteNotification delete notification of project
func DeleteNotification(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification deleted",
		"uuid":    uuid,
	})
}
