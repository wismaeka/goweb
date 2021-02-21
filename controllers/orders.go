package controllers

import (
	"goweb/config"
	"goweb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create data!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetAllOrder(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
func GetOrderById(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	if err := config.DB.Where("order_id=?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}
func UpdateOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order models.Order
	var newOrder models.Order
	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.BindJSON(&newOrder)
	if err := config.DB.Model(&order).Updates(newOrder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order models.Order
	if err := config.DB.Where("id = ?", id).Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
