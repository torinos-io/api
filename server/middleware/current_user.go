package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	authenticatedUser = "AuthenticatedUser"
)

// SetCurrentUser sets current authenticated user from authorization header
func SetCurrentUser(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
