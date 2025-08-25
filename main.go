package main

import (
	"fmt"
	"my-gin-app/project/config"
	"my-gin-app/project/models"
	"my-gin-app/project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.InitMinio()

	if err := config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.BankAccount{}, &models.Payment{}); err != nil {
    panic(fmt.Sprintf("Failed to migrate: %v", err))}

	r := gin.Default()
	routes.Setup(r)

	r.Run(":8080")
}
	