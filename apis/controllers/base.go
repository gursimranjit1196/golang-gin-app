package controllers

import (
	"fmt"
	"gin-app/apis/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize() {
	DBInit(server)
	server.Router = gin.Default()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func DBInit(server *Server) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", "localhost", "5432", "gursimranjit", "gin-app")
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println("UNABLE TO CONNECT DB")
	} else {
		server.DB = db
		fmt.Println("CONNECTED WITH DATABASE...")
		DBMigrations(server)
	}
}

func DBMigrations(server *Server) {
	server.DB.Debug().AutoMigrate(
		&models.User{},
	)
	fmt.Println("DB MIGRATIONS DONE...")
}
