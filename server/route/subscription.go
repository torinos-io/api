package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"

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

	req := &subscription_service.SubscribeRequest{}

	if err := middleware.BindJSON(c, req); err != nil {
		c.Error(errors.Wrap(err, 0))
		return
	}

	req.User = middleware.GetCurrentUser(c)

	subscription, err := service.Subscribe(req)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// Unsubscribe subscription
func Unsubscribe(c *gin.Context) {
	ac := middleware.GetAppContext(c)
	subscriptionStore := subscription_store.New(ac.MainDB)
	service := subscription_service.New(subscription_service.Context{
		Config:            ac.Config,
		SubscriptionStore: subscriptionStore,
	})

	req := &subscription_service.UnSubscribeRequest{}
	if err := middleware.BindJSON(c, req); err != nil {
		c.Error(errors.Wrap(err, 0))
		return
	}
	req.User = middleware.GetCurrentUser(c)

	if err := service.UnSubscribe(req); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
