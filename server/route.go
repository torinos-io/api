package server

import (
	"github.com/creasty/gin-contrib/app_error"
	"github.com/creasty/gin-contrib/recovery"
	"github.com/gin-gonic/gin"

	"github.com/torinos-io/api/server/middleware"
	"github.com/torinos-io/api/server/route"
	system_route "github.com/torinos-io/api/server/route/system"
	hb_service "github.com/torinos-io/api/service/hb_service"
	"github.com/torinos-io/api/type/system"
)

func drawRoutes(r *gin.Engine, appContext *system.AppContext) {
	r.Use(recovery.WrapWithCallback(func(c *gin.Context, body []byte, err interface{}) {
		hb_service.NotifyGinError(err, body, c)
	}))
	r.Use(app_error.WrapWithCallback(func(c *gin.Context, body []byte, err error) {
		hb_service.NotifyGinError(err, body, c)
	}))
	r.Use(middleware.SetAppContextWrapper(appContext))

	{
		r.GET("/", route.Ping)
		r.GET("/ping", route.Ping)
		r.GET("/user", nil)
		r.GET("/oauth/github/authorization", nil)
		r.POST("/oauth/github/authentication", nil)
		r.POST("/projects", nil)
		r.GET("/projects/:uuid", nil)
		r.POST("/projects/:uuid/notification", nil)
		r.DELETE("/projects/:uuid/notification", nil)
	}

	{
		r := r.Group("/system")
		r.Use(gin.BasicAuth(gin.Accounts{
			appContext.Config.BasicAuthUsername: appContext.Config.BasicAuthPassword,
		}))

		r.GET("/appinfo", system_route.GetAppInfo)
	}

}
