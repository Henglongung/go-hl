package routes

import (
	"go-hl/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	route := router.Group("api/v1")

	userRoute := route.Group("/user")
	{
		userRoute.GET("/", controllers.GetUsers)
		userRoute.POST("/", controllers.CreateUser)
	}

	productRoute := route.Group("/product")
	{
		productRoute.GET("/", controllers.GetProducts)
		productRoute.POST("/", controllers.CreateProduct)
	}
}
