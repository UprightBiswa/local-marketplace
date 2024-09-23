package main

import (
	"log"

	"github.com/UprightBiswa/local-marketplace/config"
	"github.com/UprightBiswa/local-marketplace/models"
	"github.com/UprightBiswa/local-marketplace/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    router := gin.Default()
	// Connect to the database
    config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})


    // Health Check Route
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "API is running!"})
    })

    //register user routes
    routes.UserRouters(router)

    router.Run(":8080")
}
