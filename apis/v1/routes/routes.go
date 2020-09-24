package routes

import (
	"gin-app/apis/v1/controllers"
	"gin-app/apis/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Authenticate())

	// Base Controller Routes
	baseController := controllers.BaseController{}
	r.GET("/ping", baseController.Ping)

	// User Controller Routes
	userController := controllers.UserController{}
	r.POST("/signup", userController.CreateUser)
	r.GET("/users", userController.GetUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	r.GET("/users/:id", userController.GetUser)

	// Post Controller Routes
	postController := controllers.PostController{}
	r.POST("/post", postController.CreatePost)
	r.GET("/posts", postController.GetPosts)

	return r
}
