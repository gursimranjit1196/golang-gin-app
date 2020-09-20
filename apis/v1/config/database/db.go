package database

import (
	"fmt"
	"gin-app/apis/v1/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	fmt.Println("INITIALIZING DATABASE...")
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", "localhost", "5432", "gursimranjit", "gin-app")
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println("UNABLE TO CONNECT DB")
	} else {
		DB = db
		fmt.Println("CONNECTED WITH DATABASE...")
		DBMigrations()
	}
}

func DBMigrations() {
	DB.Debug().AutoMigrate(
		&models.User{},
	)
	fmt.Println("DB MIGRATIONS DONE...")
}

func GetDB() *gorm.DB {
	return DB
}
