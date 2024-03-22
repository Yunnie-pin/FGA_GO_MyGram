package database

import (
	"fmt"
	"log"
	"mygram/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
	dburl    string
	config   string
	err      error
	db       *gorm.DB
)

func StartDB() {

	host = os.Getenv("PG_HOST")
	port = os.Getenv("PG_PORT")
	user = os.Getenv("PG_USER")
	password = os.Getenv("PG_PASSWORD")
	dbname = os.Getenv("PG_DBNAME")
	dburl = os.Getenv("DATABASE_URL")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		config = dburl
	} else {
		config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}

	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	fmt.Println("Successfully connected to the database!")
	db.Debug().AutoMigrate(&models.User{}, models.Photo{}, models.Socialmedia{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
