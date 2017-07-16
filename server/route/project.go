package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProject creates project
func CreateProject(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		return
	}

	files := form.File["upload[]"]

	c.JSON(http.StatusOK, gin.H{
		"message": "Project created",
	})
}

// GetProject returns the project
func GetProject(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetProject",
		"uuid":    uuid,
		"result":  "",
	})
}
