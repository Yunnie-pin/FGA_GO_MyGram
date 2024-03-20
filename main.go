package main

import (
	"log"
	"mygram/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := router.StartApp()

	log.Println("ENV: ", os.Getenv("ENVIRONMENT_VARIABEL"))
	log.Println("PORT: ", os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}
