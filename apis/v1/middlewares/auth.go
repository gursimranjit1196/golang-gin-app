package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AUTHENTICATING...")

		if false {
			fmt.Println("UNAUTHORIZED USER...")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "User Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
