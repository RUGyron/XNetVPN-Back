package main

import (
	"XNetVPN-Back/services/utils"
	yookassapackage "XNetVPN-Back/services/yookassa"
)

func init() {
	utils.InitValidator()
	//repositories.ConnectToMongoDB()
}

func main() {
	err := yookassapackage.RequestBillingSave("pivosh098@gmail.com")
	if err != nil {
		panic(err)
	}
	//router := gin.Default()
	//router.Use(cors.New(services.GetCorsConfig()))
	//routes.SetupRoutes(router)
	//err := router.Run(":" + os.Getenv("PORT"))
	//if err != nil {
	//	fmt.Println("Error starting server:", err)
	//	return
	//}
}
