package controllers

import (
	"go-hl/config"
	"go-hl/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user models.User

	// user = models.User{
	// 	Name:  "Henglong",
	// 	Email: "Henglong@gmail.com",
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var exUsers []models.User
	config.DB.Find(&exUsers)

	if len(user.Name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created Successfully",
		"user":    user,
	})
}

func GetUsers(c *gin.Context) {

	var db = *config.DB

	var users []models.User
	db.Debug().Where("name = ?", "Henglong").Preload("Products").Find(&users)

	c.JSON(http.StatusOK, users)
}
