package main

import (
	"github.com/joho/godotenv"
	"go/web-api/internal/link"
	"go/web-api/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	err = db.AutoMigrate(&link.Link{})
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal(err)
	}
}
