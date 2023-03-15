package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lappet69/resep_makanan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading")
	}
}

func ConnectDB() {

	var err error
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbname := os.Getenv("DB_NAME")
	uname := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	tz := "ASIA/JAKARTA"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, uname, pass, dbname, port, tz)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Category{}, &models.FoodIngredient{}, &models.FoodRecipe{}, &models.RecipeDetail{})
}
