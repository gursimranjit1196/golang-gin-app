package routes

import (
	"fmt"
	"gin-app/apis/v1/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	fmt.Println("INITIALIZING ROUTES...")
	r := gin.Default()

	// Base Controller Routes
	baseController := controllers.BaseController{}
	r.GET("/ping", baseController.Ping)

	// User Controller Routes
	userController := controllers.UserController{}
	r.POST("/user", userController.CreateUser)
	r.GET("/users", userController.GetUsers)
	r.PUT("/users/:id", userController.UpdateUser)

	return r
}
