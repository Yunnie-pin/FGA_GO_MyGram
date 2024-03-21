package router

import (
	"mygram/controllers"
	"mygram/middlewares"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("views/index.html")

	//web
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html",
			gin.H{
				"status": strings.ToUpper(os.Getenv("ENVIRONMENT_VARIABEL")),
			})
	})

	//api
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.AuthMiddleware(), controllers.UserUpdate)
	}

	return r
}
