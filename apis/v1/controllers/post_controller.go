package controllers

import (
	"gin-app/apis/v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (uc *PostController) CreatePost(c *gin.Context) {
	u := models.User{}
	loggedInUser, err := u.GetUser(DB, 1)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "N0 logged in user found.",
		})
		return
	}

	var post models.Post
	post.User = *loggedInUser
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
		return
	}

	createdPost, err := post.CreatePost(DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": createdPost,
	})
}

func (uc *PostController) GetPosts(c *gin.Context) {
	post := models.Post{}
	posts, err := post.GetAllPosts(DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": posts,
	})
}
