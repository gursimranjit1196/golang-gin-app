package v1

import (
	"fmt"
	"gin-app/apis/v1/config/database"
	"gin-app/apis/v1/config/server"
	"gin-app/apis/v1/controllers"
	"gin-app/apis/v1/routes"
)

func init() {
	fmt.Println("INITIALIZING V1 API SERVER...")
}

func Run() {
	database.InitDB()
	controllers.SetDB()

	routes := routes.InitRoutes()

	apiPort := fmt.Sprintf(":%s", "9000")
	fmt.Printf("LISTENING TO PORT %s...", apiPort)

	server.Run(apiPort, routes)
}
