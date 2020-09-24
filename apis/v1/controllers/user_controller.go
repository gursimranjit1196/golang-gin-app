package controllers

import (
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/services/user_service"
	"strconv"

	"gin-app/apis/v1/utils/response_handler"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		response_handler.Error(c, 412, err.Error())
		return
	}

	createdUser, err := user.CreateUser(DB)
	if err != nil {
		response_handler.Error(c, 412, err.Error())
		return
	}

	authToken, err := user_service.GenerateUserAuthToken(createdUser)

	response_handler.Success(c, 201, constants.UserCreatedSuccessfullyMsg, map[string]interface{}{
		"user":      createdUser,
		"authToken": authToken,
	})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	user := models.User{}
	users, err := user.GetAllUsers(DB)
	if err != nil {
		response_handler.Error(c, 500, err.Error())
		return
	}

	response_handler.Success(c, 200, constants.UsersFetchedSuccessfullyMsg, users)
}

func (uc *UserController) GetUser(c *gin.Context) {
	UID := c.Param("id")
	userID, err := strconv.Atoi(UID)
	if err != nil {
		response_handler.Error(c, 400, constants.InvalidUserIDMsg)
		return
	}

	userStruct := models.User{}
	user, err := userStruct.GetUser(DB, userID)
	if err != nil {
		response_handler.Error(c, 400, constants.UserNotFoundMsg)
		return
	}

	authToken, err := user_service.GenerateUserAuthToken(user)

	response_handler.Success(c, 200, constants.UserFetchedSuccessfullyMsg, map[string]interface{}{
		"user":      user,
		"authToken": authToken,
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	UID := c.Param("id")
	userID, err := strconv.Atoi(UID)
	if err != nil {
		response_handler.Error(c, 400, constants.InvalidUserIDMsg)
		return
	}

	userStr := models.User{}
	user, err := userStr.GetUser(DB, userID)
	if err != nil {
		response_handler.Error(c, 400, constants.UserNotFoundMsg)
		return
	}

	if err := c.ShouldBind(&user); err != nil {
		response_handler.Error(c, 412, err.Error())
		return
	}

	updatedUser, err := user.UpdateUser(DB)
	if err != nil {
		response_handler.Error(c, 412, err.Error())
		return
	}

	response_handler.Success(c, 200, constants.UserUpdatedSuccessfullyMsg, updatedUser)
}
