package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	JWTSecret []byte
	BaseURL string
	ApiKey string
	MinioClient *minio.Client
	S3Bucket    string
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

func InitMinio() {
	endpoint := os.Getenv("S3_ENDPOINT")      // misal http://localhost:9000
	accessKeyID := os.Getenv("S3_ACCESS_KEY") // S3_ID
	secretAccessKey := os.Getenv("S3_SECRET_KEY")
	S3Bucket = os.Getenv("S3_BUCKET")         // misal "mybucket"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // kalau local http
	})
	if err != nil {
		log.Fatal(err)
	}
	MinioClient = minioClient

	// create bucket kalau belum ada
	exists, err := MinioClient.BucketExists(context.Background(), S3Bucket)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		if err := MinioClient.MakeBucket(context.Background(), S3Bucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Minio ready")
}