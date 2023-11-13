package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/template-service/pkg/utils"
)

var apiKey string

// Authenticate is a middleware that checks if the request has a valid API key (X-API-KEY header).
// This API key should be coming from another microservice that is trying to access this service.
// If the API key is not valid, the request is aborted with a 401 status code.
func Authenticate() gin.HandlerFunc {
	if apiKey == "" {
		utils.LoadEnv()
		apiKey = os.Getenv("API_KEY")
	}

	return func(c *gin.Context) {
		key := c.Request.Header.Get("X-API-KEY")
		if key != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "unauthorized"})
			return
		}
	}
}
