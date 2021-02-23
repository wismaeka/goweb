package main

import (
	"goweb/config"
	"goweb/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	r.POST("/order", controllers.CreateOrder)
	r.GET("/orders", controllers.GetAllOrder)
	r.GET("/order/:order_id", controllers.GetOrderById)
	r.PUT("/order/:order_id", controllers.UpdateOrder)
	r.DELETE("/order/:order_id", controllers.DeleteOrder)

	r.Run()
}
