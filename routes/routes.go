package routes

import (
	"my-gin-app/project/controllers"
	"my-gin-app/project/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.AuthMiddleware(), controllers.Me)
}
