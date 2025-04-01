package services

import (
	"XNetVPN-Back/config"
	"github.com/gin-contrib/cors"
	"time"
)

func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     config.Config.CorsAllowedOrigins,
		AllowMethods:     config.Config.CorsAllowedMethods,
		AllowHeaders:     config.Config.CorsAllowedHeaders,
		ExposeHeaders:    config.Config.CorsExposedHeaders,
		AllowCredentials: config.Config.CorsAllowCredentials,
		MaxAge:           time.Duration(config.Config.CorsMaxAge) * time.Hour,
	}
}
