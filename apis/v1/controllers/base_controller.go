package controllers

import (
	"gin-app/apis/v1/config/database"
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/utils/response_handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type BaseController struct{}

func SetDB() {
	DB = database.GetDB()
}

func (bc *BaseController) Ping(c *gin.Context) {
	response_handler.Success(c, 200, constants.V1ApisWorkingMsg, nil)
}
