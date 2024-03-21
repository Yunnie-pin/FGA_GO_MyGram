package middlewares

import (
	"mygram/database"
	"mygram/models"

	"mygram/helpers"
	"net/http"

	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userId, err := strconv.Atoi(c.Param("userId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "BAD REQUEST",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))
		User := models.User{}

		err = db.Select("id").First(&User, uint(userId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "NOT FOUND",
				"message": "user not found",
			})
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "you don't have permission",
			})
			return
		}

		c.Next()
	}
}
