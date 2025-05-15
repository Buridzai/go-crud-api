package main

import (
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/routes"

	_ "go-crud-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go CRUD API
// @version 1.0
// @host localhost:8080
// @BasePath /api
func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.Static("/web", "./frontend")

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r)
	r.Run(":8080")
}
