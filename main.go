package main

import (
	"log"
	"mygram/router"
	"os"
)

func main() {
	r := router.StartApp()
	env := os.Getenv("ENV")

	// if env == "development" {
	// 	log.Println("Server running on port :8080")
	// 	r.Run(":8080")
	// } else if env == "production" {
	// 	log.Println("Server running on port :80")
	// 	r.Run(":80")
	// }

	log.Print("ENV: ", env)
	log.Println("Server running on port :8080")
	r.Run(":8080")
}
