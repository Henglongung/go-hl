package main

import (
	"go-hl/config"
	"go-hl/models"
	"go-hl/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{}, &models.Product{})

	routes.SetUpRoutes(route)

	route.Run(":8080")

}
