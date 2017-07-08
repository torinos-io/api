package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Analyze project
func Analyze(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Analyze",
	})
}

// GetAnalyzedResult return the analyzed result
func GetAnalyzedResult(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetAnalyzedResult",
		"uuid":    uuid,
		"result":  "",
	})
}
