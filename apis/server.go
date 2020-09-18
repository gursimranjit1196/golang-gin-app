package apis

import (
	"fmt"
	"gin-app/apis/controllers"
)

func init() {
	fmt.Println("API SERVER INITIALIZING...")
}

func Run() {
	server := controllers.Server{}
	server.Initialize()

	apiPort := fmt.Sprintf(":%s", "9000")
	fmt.Printf("Listening to port %s", apiPort)

	server.Run(apiPort)
}
