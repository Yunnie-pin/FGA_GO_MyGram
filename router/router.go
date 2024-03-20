package router

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("views/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	userRouter := r.Group("/users")
	{
		userRouter.GET("/register", controllers.UserRegister)
		userRouter.GET("/login", controllers.UserLogin)
	}

	return r
}
