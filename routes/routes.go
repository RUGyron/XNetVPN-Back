package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/location", controllers.Location)
	router.POST("/locations", controllers.Locations)
}
