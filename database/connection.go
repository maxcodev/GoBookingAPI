package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	HOST := os.Getenv("HOST")
	USER := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	PORT := os.Getenv("PORT")

	var DSN string = "host=" + HOST + " user=" + USER + " password=" + PASSWORD + " dbname=" + DBNAME + " port=" + PORT
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected!")
	}
}
