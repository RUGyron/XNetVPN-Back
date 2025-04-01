package routes

import (
	"github.com/gin-gonic/gin"
	"xnet-vpn-back/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login/:key", controllers.Login)
}
