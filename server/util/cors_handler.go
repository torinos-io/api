package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// NewCorsHandler creates a new http.Handler to support CORS
func NewCorsHandler(engine *gin.Engine, allowOrigins []string) http.Handler {
	mw := cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin"},
		Debug:            false,
	})

	return mw.Handler(engine)
}
