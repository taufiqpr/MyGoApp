package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	JWTSecret []byte
	BaseURL string
	ApiKey string
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_DSN")
	secret := os.Getenv("JWT_SECRET")
	BaseURL = os.Getenv("SPOONACULAR_BASE_URL")
	ApiKey = os.Getenv("SPOONACULAR_API_KEY")

	if dsn == "" || secret == "" {
		log.Fatal("DATABASE_DSN or JWT_SECRET is empty")
	}
	if BaseURL == "" || ApiKey == "" {
		log.Fatal("SPOONACULAR_BASE_URL or SPOONACULAR_API_KEY is empty")
	}

	JWTSecret = []byte(secret)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = dbConn

	fmt.Println("Database connected & config loaded")
}
