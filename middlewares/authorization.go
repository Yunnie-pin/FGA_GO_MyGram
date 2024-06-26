package middlewares

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Select("id").First(&User, uint(userID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "NOT FOUND",
				"message": "user not found",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		productId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD REQUEST",
				"message": "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "NOT FOUND",
				"message": "photo not found",
			})
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "FORBIDDEN",
				"message": "You are not authorized",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD REQUEST",
				"message": "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "NOT FOUND",
				"message": "comment not found",
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "FORBIDDEN",
				"message": "You are not authorized",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD REQUEST",
				"message": "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "NOT FOUND",
				"message": "social media not found",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "FORBIDDEN",
				"message": "You are not authorized",
			})
			return
		}

		c.Next()
	}
}
