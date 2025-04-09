package routes

import (
	"XNetVPN-Back/controllers"
	"XNetVPN-Back/services/jwt"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.GET("/subscriptions", controllers.Subscriptions)
	router.POST("/update-token", jwt.RefreshBearerRequired(), controllers.UpdateToken)
	router.POST("/profile", jwt.AccessBearerRequired(), controllers.Profile)
	router.POST("/device/add", jwt.AccessBearerRequired(), controllers.AddDevice)
}
