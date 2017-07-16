package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/torinos-io/api/server/middleware"
)

// GetCurrentUser return the current user
func GetCurrentUser(c *gin.Context) {
	if user := middleware.GetCurrentUser(c); user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusUnauthorized, "Can not find authorized current user")
	}
}
