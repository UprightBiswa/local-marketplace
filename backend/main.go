package main

import (
	"github.com/UprightBiswa/local-marketplace/config"
	"github.com/UprightBiswa/local-marketplace/models"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	// Connect to the database
    config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})


    // Health Check Route
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "API is running!"})
    })

    router.Run(":8080")
}
