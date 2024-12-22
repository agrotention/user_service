package main

import (
	"log"
	"os"

	"github.com/agrotention/user_service/db"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	conn, err := gorm.Open(postgres.Open(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err.Error())
	}
	conn.AutoMigrate(&db.User{})
}
