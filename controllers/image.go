package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"my-gin-app/project/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	defer file.Close()

	if header.Size > 2*1024*1024 || header.Size < 10*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file size must be 10KB-2MB"})
		return
	}

	if !strings.HasSuffix(strings.ToLower(header.Filename), ".jpg") &&
		!strings.HasSuffix(strings.ToLower(header.Filename), ".jpeg") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file must be .jpg or .jpeg"})
		return
	}

	fileName := uuid.New().String() + ".jpeg"

	_, err = config.MinioClient.PutObject(
		context.Background(),
		config.S3Bucket,
		fileName,
		file,
		header.Size,
		minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload"})
		return
	}

	url := fmt.Sprintf("%s/%s/%s", os.Getenv("S3_BASE_URL"), config.S3Bucket, fileName)
	c.JSON(http.StatusOK, gin.H{"imageUrl": url})
}
