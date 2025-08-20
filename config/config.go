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
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_DSN")
	secret := os.Getenv("JWT_SECRET")

	if dsn == "" || secret == "" {
		log.Fatal("DATABASE_DSN or JWT_SECRET is empty")
	}

	JWTSecret = []byte(secret)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = dbConn

	fmt.Println("Database connected & config loaded")
}
