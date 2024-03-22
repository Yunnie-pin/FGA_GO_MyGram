package controllers

import "github.com/gin-gonic/gin"

func SocialMediaCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Social Media Created",
	})
}