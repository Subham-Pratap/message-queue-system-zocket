// handlers.go
package api

import (
	"fmt"
	"net/http"

	"message-queue-system-zocket/database"
	"message-queue-system-zocket/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	// Check if the user exists
	var user models.User
	if err := db.First(&user, product.UserID).Error; err != nil {
		// User doesn't exist, create the user
		newUser := models.User{
			ID:   product.UserID, // Assign the user ID
			Name: fmt.Sprintf("User%d", product.UserID),
			// Assign other user fields as needed
		}
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		user = newUser
	}

	// Associate the product with the user
	product.UserID = user.ID

	// Create the product
	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
