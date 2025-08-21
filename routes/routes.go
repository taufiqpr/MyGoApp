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
	r.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)
	r.PUT("/profile", middleware.AuthMiddleware(), controllers.UpdateProfile)
	r.DELETE("/profile", middleware.AuthMiddleware(), controllers.DeleteProfile)
	r.GET("/recipes", middleware.AuthMiddleware(), controllers.GetRecipes)
	r.GET("/recipe-detail", middleware.AuthMiddleware(), controllers.GetRecipeDetail)
}
