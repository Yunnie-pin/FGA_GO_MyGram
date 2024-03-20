package main

import (
	"fmt"
	"mygram/helpers"
	"mygram/router"
)

func main() {
	r := router.StartApp()
	env := helpers.GetEnv("ENV")
	port := helpers.GetEnv("PORT")

	// if env == "development" {
	// 	log.Println("Server running on port :8080")
	// 	r.Run(":8080")
	// } else if env == "production" {
	// 	log.Println("Server running on port :80")
	// 	r.Run(":80")
	// }

	fmt.Println("ENV: ", env)
	fmt.Println("PORT: ", port)
	r.Run(":" + port)
}
