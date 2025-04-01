package middleware

import (
	"github.com/gin-gonic/gin"
	"xnet-vpn-back/config"
	"xnet-vpn-back/responses"
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
