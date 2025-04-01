package middleware

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/responses"
	"github.com/gin-gonic/gin"
)

func MaintenanceAndKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("api-key") != config.Config.ApiKey {
			c.AbortWithStatusJSON(responses.Forbidden("invalid key"))
			return
		}
		c.Next()
	}
}
