package middleware

import (
	"github.com/gin-gonic/gin"

	oauth_service "github.com/torinos-io/api/service/oauth_service"
	user_store "github.com/torinos-io/api/store/user_store"
	"github.com/torinos-io/api/type/model"
	"github.com/torinos-io/api/type/system"
)

const (
	currentUserContextName = "CurrentUser"
)

// SetCurrentUser sets current authenticated user from authorization header
func SetCurrentUser(appContext *system.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")

		if h == "" {
			c.Next()
			return
		}

		userStore := user_store.New(appContext.MainDB)
		service := oauth_service.New(oauth_service.Context{
			Config:    appContext.Config,
			UserStore: userStore,
		})

		user, err := service.FindByAuthToken(h)

		if err != nil {
			c.Next()
			return
		}

		c.Set(currentUserContextName, user)
		c.Next()
	}
}

// GetCurrentUser returns current authenticated user from context
func GetCurrentUser(c *gin.Context) *model.User {
	v, exists := c.Get(currentUserContextName)

	if !exists {
		return nil
	}

	user, ok := v.(*model.User)

	if !ok {
		return nil
	}

	return user
}
