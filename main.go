package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/joho/godotenv"
	"os"

	"note/app/config"
	"note/app/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	port := os.Getenv("PORT")

	// Connect to the database
	config.ConnectDatabase()

	// Set up Gin server
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "Server is running!",
		})
	})
	api := server.Group("/api")
	routes.RegisterRoutes(api)
	server.Run(":" + port)
}