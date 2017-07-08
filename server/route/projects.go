package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Analyze project
func Analyze(c *gin.Context) {
	c.String(http.StatusOK, "Analyze")
}

// Result return the analyzed result
func Result(c *gin.Context) {
	uuid := c.Param("uuid")
	c.String(http.StatusOK, uuid)
}
