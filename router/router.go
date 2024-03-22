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
	guestRouter := r.Group("/users")
	{
		guestRouter.POST("/register", controllers.UserRegister)
		guestRouter.POST("/login", controllers.UserLogin)
	}

	userRouter := r.Use(middlewares.Authentication())
	{
		userRouter.PUT("/users", middlewares.UserAuthorization(), controllers.UserUpdate)
		userRouter.DELETE("/users", middlewares.UserAuthorization(), controllers.UserDelete)

		photoRouter := r.Group("/photos")
		{
			photoRouter.GET("/all", controllers.PhotoGetAll)
			photoRouter.GET("/", controllers.PhotoGet)
			photoRouter.POST("/", controllers.PhotoCreate)
			photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoUpdate)
			photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoDelete)
		}

		commentRouter := r.Group("/comments")
		{
			commentRouter.GET("/all", controllers.CommentGetAll)
			commentRouter.GET("/", controllers.CommentGet)
			commentRouter.POST("/", controllers.CommentCreate)
			commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.CommentUpdate)
			commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.CommentDelete)
		}

		socialMediaRouter := r.Group("/socialmedias")
		{
			socialMediaRouter.POST("/", controllers.SocialMediaCreate)
		}
	}
	return r
}
