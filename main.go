package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"xnet-vpn-back/routes"
	"xnet-vpn-back/services"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(services.GetCorsConfig()))
	routes.SetupRoutes(router)
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
