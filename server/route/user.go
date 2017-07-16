package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser return the current user
func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if exists {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}
