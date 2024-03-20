package controllers

import "github.com/gin-gonic/gin"

func UserRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ini adalah user register",
	})
}

func UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ini adalah user login",
	})
}
