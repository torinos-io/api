package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create noitication of project
func Create(c *gin.Context) {
	uuid := c.Param("uuid")
	c.String(http.StatusOK, "Create notification:"+uuid)
}

// Delete notification of project
func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	c.String(http.StatusOK, "Delete notification:"+uuid)
}
