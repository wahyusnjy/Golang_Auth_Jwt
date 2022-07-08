package main

import (
	"auth/controllers"
	"auth/database"
	"auth/middlewares"
	"github.com/gin-gonic/gin"
)
func main() {
	//inisialisasi Databasenya
	database.Connect("root@tcp(localhost:3306)/jwt-go?parseTime=true")
	database.Migrate()
	//inisialisasi Route
	router := initRouter()
	router.Run("8080")
}

func initRouter() *gin.Engine{
	router := gin.Default()
	api := router.Group("/Api")
	{
			api.POST("/token", controllers.GenerateToken)
			api.POST("/user/register", controllers.RegisterUser)
			secured := api.Group("/secured").Use(middlewares.Auth())
			{
				secured.GET("/ping", controllers.Ping)
			}
	}
	return router
}
