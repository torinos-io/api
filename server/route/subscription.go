package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/torinos-io/api/server/middleware"
	subscription_service "github.com/torinos-io/api/service/subscription_service"
	subscription_store "github.com/torinos-io/api/store/subscription_store"
)

// Subscribe subscription
func Subscribe(c *gin.Context) {
	ac := middleware.GetAppContext(c)
	subscriptionStore := subscription_store.New(ac.MainDB)
	service := subscription_service.New(subscription_service.Context{
		Config:            ac.Config,
		SubscriptionStore: subscriptionStore,
	})

	projectUUID := c.Param("project_uuid")
	email := c.Param("email")

	req := &subscription_service.SubscribeRequest{
		ProjectUUID: projectUUID,
		Email:       email,
	}

	user := middleware.GetCurrentUser(c)

	subscription, err := service.Subscribe(req, user)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// UnSubscribe subscription
func UnSubscribe(c *gin.Context) {
	ac := middleware.GetAppContext(c)
	subscriptionStore := subscription_store.New(ac.MainDB)
	service := subscription_service.New(subscription_service.Context{
		Config:            ac.Config,
		SubscriptionStore: subscriptionStore,
	})

	projectUUID := c.Param("project_uuid")

	req := &subscription_service.UnSubscribeRequest{
		ProjectUUID: projectUUID,
	}
	user := middleware.GetCurrentUser(c)

	if err := service.UnSubscribe(req, user); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
