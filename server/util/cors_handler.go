package util

import (
	"net/http"

	"github.com/rs/cors"
)

// NewCorsHandler creates a new http.Handler to support CORS
func NewCorsHandler(handler http.Handler, allowOrigins []string) http.Handler {
	mw := cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type", "Accept"},
		Debug:            false,
	})

	return mw.Handler(handler)
}
