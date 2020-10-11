package database

import (
	"fmt"
	"gin-app/apis/v1/config/channels"
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/utils/loggers"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	loggers.Log(constants.InitDBLog)
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", "localhost", "5432", "gursimranjit", "gin-app")
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		loggers.Log(constants.UnableToConnectDBLog)
		channels.CriticalErrorChannel <- constants.UnableToConnectDBLog
	} else {
		DB = db
		loggers.Log(constants.DBConnectedLog)
		DBMigrations()
	}
}

func DBMigrations() {
	DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	loggers.Log(constants.DBMigratedLog)
}

func GetDB() *gorm.DB {
	return DB
}
