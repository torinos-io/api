package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNotification craete noitication of project
func CreateNotification(c *gin.Context) {
	uuid := c.Param("uuid")
	c.String(http.StatusOK, "Create notification:"+uuid)
}

// DeleteNotification delete notification of project
func DeleteNotification(c *gin.Context) {
	uuid := c.Param("uuid")
	c.String(http.StatusOK, "Delete notification:"+uuid)
}
