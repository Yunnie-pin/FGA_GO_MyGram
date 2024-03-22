package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SocialMediaGetAll(c *gin.Context) {

	db := database.GetDB()
	SocialMedia := []models.SocialMedia{}

	err := db.Debug().Model(&models.SocialMedia{}).Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "social media not found",
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func SocialMediaGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := []models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	//return all social media by user
	err := db.Debug().Model(&models.SocialMedia{}).Where("user_id = ?", userID).Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "social media not found",
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)

}

func SocialMediaCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func SocialMediaUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMediaID := c.Param("socialMediaId")

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaURL: SocialMedia.SocialMediaURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMediaID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func SocialMediaDelete(c *gin.Context) {
	db := database.GetDB()
	socialMediaID := c.Param("socialMediaId")

	SocialMedia := models.SocialMedia{}

	err := db.Where("id = ?", socialMediaID).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "social media not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
