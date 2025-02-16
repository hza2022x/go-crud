package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/go-crud/config"
	"github.com/yodfhafx/go-crud/controllers"
	"github.com/yodfhafx/go-crud/middleware"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")

	authGroup := v1.Group("auth")
	authController := controllers.Auth{DB: db}
	{
		authGroup.POST("/sign-up", authController.Signup)
		authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
	}

	articlesGroup := v1.Group("articles")
	articleController := controllers.Articles{DB: db}
	{
		articlesGroup.GET("", articleController.FindAll)
		articlesGroup.GET("/:id", articleController.FindOne)
		articlesGroup.PATCH("/:id", articleController.Update)
		articlesGroup.DELETE("/:id", articleController.Delete)
		articlesGroup.POST("", articleController.Create)
	}

	categoriesGroup := v1.Group("categories")
	categoryController := controllers.Categories{DB: db}
	{
		categoriesGroup.GET("", categoryController.FindAll)
		categoriesGroup.GET("/:id", categoryController.FindOne)
		categoriesGroup.PATCH("/:id", categoryController.Update)
		categoriesGroup.DELETE("/:id", categoryController.Delete)
		categoriesGroup.POST("", categoryController.Create)
	}
}
