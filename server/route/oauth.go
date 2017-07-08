package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorization user
func Authorization(c *gin.Context) {
	c.String(http.StatusOK, "Authorization")
}

// Authentication user
func Authentication(c *gin.Context) {
	c.String(http.StatusOK, "Authentication")
}
