package main

import (
	"fmt"
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Start the main function")
	// Initialize Database
	database.Connect("username:password@tcp(host:port)/db?parseTime=true")
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":9017")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/token", controllers.GenerateToken)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
