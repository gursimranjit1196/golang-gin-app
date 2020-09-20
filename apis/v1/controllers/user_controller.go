package controllers

import (
	"gin-app/apis/v1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
		return
	}

	createdUser, err := user.CreateUser(DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": createdUser,
	})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	user := models.User{}
	users, err := user.GetAllUsers(DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": users,
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	UID := c.Param("id")
	userID, err := strconv.ParseUint(UID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid User Id",
		})
		return
	}

	userStr := models.User{}
	user, err := userStr.GetUser(DB, userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No User found with given id.",
		})
		return
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
		return
	}

	updatedUser, err := user.UpdateUser(DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":   200,
		"response": updatedUser,
	})
}
