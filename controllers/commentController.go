package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentGetAll(c *gin.Context) {

	db := database.GetDB()
	Comments := []models.Comment{}

	err := db.Debug().Model(&models.Comment{}).Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "comment not found",
		})
		return
	}

	c.JSON(http.StatusOK, Comments)
}

func CommentGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comments := []models.Comment{}
	userID := uint(userData["id"].(float64))

	//return all comments by user
	err := db.Debug().Model(&models.Comment{}).Where("user_id = ?", userID).Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "comment not found",
		})
		return
	}

	c.JSON(http.StatusOK, Comments)

}

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.BindJSON(&Comment)
	} else {
		c.Bind(&Comment)
	}

	errPhoto := db.First(&Photo, Comment.PhotoID).Error
	if errPhoto != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "photo not found",
		})
		return
	}

	Comment.UserID = userID

	errComment := db.Create(&Comment).Error
	if errComment != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": errComment.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)

}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.BindJSON(&Comment)
	} else {
		c.Bind(&Comment)
	}

	err := db.Model(&Comment).Where("id = ?", userID).Updates(
		models.Comment{
			Message: Comment.Message,
		},
	).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "successfully updated comment",
		},
	)

}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	commentID := c.Param("commentId")
	userID := uint(userData["id"].(float64))

	err := db.Where("id = ? AND user_id = ?", commentID, userID).Delete(&models.Comment{}).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "Your comment has been successfully deleted",
		},
	)

}
