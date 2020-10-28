package config

import (
	"fmt"
	"os"

	"github.com/IvanaKrizic/Dashboard/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to read .env file.")
	}

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSERNAME"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	database, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}
	DB = database
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Statistic{}, &models.AuthData{})
}
