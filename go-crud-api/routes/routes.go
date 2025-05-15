package routes

import (
	"go-crud-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)

		api.GET("/users/:id", controllers.GetUserByID)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}
}
