package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/torinos-io/api/server/middleware"
)

// GetCurrentUser return the current user
func GetCurrentUser(c *gin.Context) {
	user := middleware.GetCurrentUser(c)
	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, user)
}
