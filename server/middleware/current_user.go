package middleware

import (
	"strings"
	
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
		
		splited := strings.Split(h, "Bearer")
		
		if len(splited) < 2 {
			c.Next()
			return
		}
		
		
		token := strings.TrimSpace(splited[1])
		
		userStore := user_store.New(appContext.MainDB)
		service := oauth_service.New(oauth_service.Context{
			Config:    appContext.Config,
			UserStore: userStore,
		})

		user, err := service.FindByAccessToken(token)

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
	if v, exists := c.Get(currentUserContextName); exists {
		if user, ok := v.(*model.User); ok {
			return user
		}
	}

	return nil
}
