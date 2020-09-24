package database

import (
	"fmt"
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/utils/loggers"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	loggers.Log(constants.InitDB)
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", "localhost", "5432", "gursimranjit", "gin-app")
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		loggers.Log(constants.UnableToConnectDB)
	} else {
		DB = db
		loggers.Log(constants.DBConnected)
		DBMigrations()
	}
}

func DBMigrations() {
	DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	loggers.Log(constants.DBMigrated)
}

func GetDB() *gorm.DB {
	return DB
}
