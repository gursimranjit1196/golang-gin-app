package controllers

import (
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/utils/response_handler"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (uc *PostController) CreatePost(c *gin.Context) {
	lUser, exists := c.Get(constants.LoggedInUserKey)
	if !exists {
		response_handler.Error(c, 401, constants.NoLoggedInUserMsg)
		return
	}

	loggedInUser := lUser.(*models.User)

	var post models.Post
	post.UserID = int(loggedInUser.ID)
	if err := c.ShouldBind(&post); err != nil {
		response_handler.Error(c, 412, err.Error())
		return
	}

	createdPost, err := post.CreatePost(DB)
	if err != nil {
		response_handler.Error(c, 412, err.Error())
		return
	}

	response_handler.Success(c, 201, constants.PostCreatedSuccessfullyMsg, createdPost)
}

func (uc *PostController) GetPosts(c *gin.Context) {
	post := models.Post{}
	posts, err := post.GetAllPosts(DB)
	if err != nil {
		response_handler.Error(c, 500, err.Error())
		return
	}

	response_handler.Success(c, 200, constants.PostsFetchedSuccessfullyMsg, posts)
}
