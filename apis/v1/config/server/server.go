package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Run(addr string, routes *gin.Engine) {
	log.Fatal(http.ListenAndServe(addr, routes))
}
