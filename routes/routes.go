package routes

import (
	"XNetVPN-Back/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login/:key", controllers.Login)
}
