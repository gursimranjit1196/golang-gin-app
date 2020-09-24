package v1

import (
	"fmt"
	"gin-app/apis/v1/config/database"
	"gin-app/apis/v1/config/server"
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/controllers"
	"gin-app/apis/v1/routes"
	"gin-app/apis/v1/utils/loggers"
)

func init() {
	loggers.Log(constants.InitV1APIServer)
}

func Run() {
	database.InitDB()
	controllers.SetDB()

	routes := routes.InitRoutes()

	apiPort := fmt.Sprintf(":%s", "9000")
	loggers.Log(constants.ListingToPort, apiPort)

	server.Run(apiPort, routes)
}
