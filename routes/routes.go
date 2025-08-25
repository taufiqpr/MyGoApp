package routes

import (
	"my-gin-app/project/controllers"
	"my-gin-app/project/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/healths", controllers.Healths)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.AuthMiddleware(), controllers.Me)
	r.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)
	r.PUT("/profile", middleware.AuthMiddleware(), controllers.UpdateProfile)
	r.DELETE("/profile", middleware.AuthMiddleware(), controllers.DeleteProfile)
	r.GET("/recipes", middleware.AuthMiddleware(), controllers.GetRecipes)
	r.GET("/recipe-detail", middleware.AuthMiddleware(), controllers.GetRecipeDetail)
	r.POST("/upload", middleware.AuthMiddleware(), controllers.UploadImage)
	r.POST("/products", middleware.AuthMiddleware(), controllers.CreateProduct)
	r.PATCH("/products/:id", middleware.AuthMiddleware(), controllers.UpdateProduct)
	r.DELETE("/products/:id", middleware.AuthMiddleware(), controllers.DeleteProduct)
	r.POST("/products/:id/stock", middleware.AuthMiddleware(), controllers.UpdateStock)
	r.POST("/bank/account", middleware.AuthMiddleware(), controllers.CreateBankAccount)
	r.GET("/bank/account", middleware.AuthMiddleware(), controllers.ListBankAccounts)
	r.PATCH("/bank/account/:id", middleware.AuthMiddleware(), controllers.UpdateBankAccount)
	r.DELETE("/bank/account/:id", middleware.AuthMiddleware(), controllers.DeleteBankAccount)
	r.GET("/products", middleware.AuthMiddleware(), controllers.ListProductsWithFilter)
	r.GET("/products/:id", middleware.AuthMiddleware(),controllers.GetProductDetail)
	r.POST("/products/:id/buy", middleware.AuthMiddleware(), controllers.BuyProduct)
}
