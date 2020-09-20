package controllers

import (
	"fmt"
	"gin-app/apis/v1/config/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type BaseController struct{}

func SetDB() {
	fmt.Println("SETTING DB INSTANCE IN CONTROLLER...", database.GetDB())
	DB = database.GetDB()
}

func (bc *BaseController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GIN APP IS WORKING...",
	})
}
