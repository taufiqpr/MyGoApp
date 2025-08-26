package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	JWTSecret   []byte
	MinioClient *minio.Client
	S3Bucket    string
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port,
	)

	secret := os.Getenv("JWT_SECRET")
	if dsn == "" || secret == "" {
		log.Fatal("Database config or JWT_SECRET is empty")
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
	endpoint := os.Getenv("S3_ENDPOINT")
	accessKeyID := os.Getenv("S3_ACCESS_KEY")
	secretAccessKey := os.Getenv("S3_SECRET_KEY")
	S3Bucket = os.Getenv("S3_BUCKET")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	MinioClient = minioClient

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
