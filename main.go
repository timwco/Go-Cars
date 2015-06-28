package main

import (
	"github.com/twhitacre/go-cars/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up Gin Routes
	router := gin.Default()
	router.GET("/cars", controllers.CarList)
	router.POST("/cars", controllers.CarNew)
	router.GET("/cars/:id", controllers.CarSingle)
	router.Run(":8000")
}
