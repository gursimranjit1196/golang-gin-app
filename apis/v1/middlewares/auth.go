package middlewares

import (
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/services/user_service"
	"gin-app/apis/v1/utils/loggers"
	"gin-app/apis/v1/utils/response_handler"
	"gin-app/apis/v1/utils/wrappers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		"/users/:id",
	}
	_, found := wrappers.Find(whitelistURLs, path)
	return found
}

func onAuthTokenMissing(c *gin.Context) {
	loggers.Log(constants.MissingAuthTokenLog)
	response_handler.Error(c, 401, constants.AuthTokenMissingMsg)
	c.Abort()
}

func onInvalidUserToken(c *gin.Context, err error) {
	loggers.Log(constants.InvalidAuthTokenLog, err.Error())
	response_handler.Error(c, 401, constants.InvalidAuthTokenLog)
	c.Abort()
}

func onValidUserToken(c *gin.Context, loggedInUser *models.User) {
	c.Set(constants.LoggedInUserKey, loggedInUser)

	loggers.Log(constants.LoggedInUserLog, loggedInUser)
}
