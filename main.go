// main.go
package main

import (
	"message-queue-system-zocket/api"
	"message-queue-system-zocket/database"
	"message-queue-system-zocket/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// Initialize the database
	db, err := gorm.Open(sqlite.Open("my-database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate models
	db.AutoMigrate(&models.User{}, &models.Product{})

	// Inject the database instance into the API handlers
	database.SetDB(db)

	r.POST("/api/products", api.CreateProduct)

	r.Run(":8080")
}
