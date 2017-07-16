package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	oauth_service "github.com/torinos-io/api/service/oauth_service"
	user_store "github.com/torinos-io/api/store/user_store"
	"github.com/torinos-io/api/server/middleware"
	"github.com/go-errors/errors"
)

// GetAuthorization authorize user
func GetAuthorization(c *gin.Context) {
	ac := middleware.GetAppContext(c)
	userStore := user_store.New(ac.MainDB)
	service := oauth_service.New(oauth_service.Context{
		Config:    ac.Config,
		UserStore: userStore,
	})

	c.JSON(http.StatusOK, service.GetAuthorization())
}

// Authenticate user
func Authenticate(c *gin.Context) {
	ac := middleware.GetAppContext(c)
	userStore := user_store.New(ac.MainDB)
	service := oauth_service.New(oauth_service.Context{
		Config:    ac.Config,
		UserStore: userStore,
	})

	req := &oauth_service.SaveRequest{}

	if err := middleware.BindJSON(c, req); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.Wrap(err, 0))
		return
	}

	resp, err := service.Save(req)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, resp)

}
