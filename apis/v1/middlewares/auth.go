package middlewares

import (
	"fmt"
	"gin-app/apis/v1/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AUTHENTICATING...")
		authToken := c.GetHeader("Authorization")
		if len(authToken) == 0 {
			fmt.Println("AUTH TOKEN MISSING...")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Auth token missing",
			})
			c.Abort()
			return
		}

		loggedInUser, err := user_service.ValidateUserToken(authToken)

		if err != nil {
			fmt.Println("INVALIDATE USER TOKEN...", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("loggedInUser", loggedInUser)

		fmt.Println("LOGGED IN USER IS...", loggedInUser)

		c.Next()
	}
}
