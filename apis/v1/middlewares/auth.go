package middlewares

import (
	"fmt"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/services/user_service"
	"gin-app/apis/v1/utils/wrappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AUTHENTICATING...")

		if skipUrlAuth(c.FullPath()) {
			c.Next()
			return
		}

		authToken := c.GetHeader("Authorization")
		if len(authToken) == 0 {
			onAuthTokenMissing(c)
			return
		}

		loggedInUser, err := user_service.ValidateUserToken(authToken)

		if err != nil {
			onInvalidUserToken(c, err)
			return
		}

		onValidUserToken(c, loggedInUser)

		c.Next()
	}
}

func skipUrlAuth(path string) bool {
	whitelistURLs := []string{
		"/signup",
		"/signin",
	}
	_, found := wrappers.Find(whitelistURLs, path)
	return found
}

func onAuthTokenMissing(c *gin.Context) {
	fmt.Println("AUTH TOKEN MISSING...")
	c.JSON(http.StatusUnauthorized, gin.H{
		"status": http.StatusUnauthorized,
		"error":  "Auth token missing",
	})
	c.Abort()
}

func onInvalidUserToken(c *gin.Context, err error) {
	fmt.Println("INVALIDATE USER TOKEN...", err.Error())
	c.JSON(http.StatusUnauthorized, gin.H{
		"status": http.StatusUnauthorized,
		"error":  err.Error(),
	})
	c.Abort()
}

func onValidUserToken(c *gin.Context, loggedInUser *models.User) {
	c.Set("loggedInUser", loggedInUser)

	fmt.Println("LOGGED IN USER IS...", loggedInUser)
}
