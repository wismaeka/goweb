package main

import (
	"goweb/config"
	"goweb/controllers"
	_ "goweb/docs"

	"github.com/gin-gonic/gin"                 // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Orders API
// @version 1.0
// @description This is a basic API To Make Order of Items using Gin and Gorm.

// @contact.name Orders.Support
// @contact.email wisma@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	config.ConnectDB()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/order", controllers.CreateOrder)
		v1.GET("/orders", controllers.GetAllOrder)
		v1.GET("/order/:order_id", controllers.GetOrderById)
		v1.PUT("/order/:order_id", controllers.UpdateOrder)
		v1.DELETE("/order/:order_id", controllers.DeleteOrder)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
