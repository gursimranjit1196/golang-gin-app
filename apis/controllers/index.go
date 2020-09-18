package controllers

import "github.com/gin-gonic/gin"

func (s *Server) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GIN APP IS WORKING...",
	})
}
