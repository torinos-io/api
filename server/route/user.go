package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUser return the current user
func CurrentUser(c *gin.Context) {
	c.String(http.StatusOK, "CurrentUser")
}
