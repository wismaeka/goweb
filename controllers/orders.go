package controllers

import (
	"goweb/config"
	"goweb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	c.ShouldBindJSON(&order)
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create data!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetAllOrder(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
func GetOrderById(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	if err := config.DB.Preload(("Items")).Where("order_id=?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}
func UpdateOrder(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	if err := config.DB.Where("order_id=?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.ShouldBindJSON(&order)
	if err := config.DB.Where("order_id=?", order_id).Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	var item models.Item
	// if err := config.DB.Where("order_id=?", order_id).Delete(&order).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
	// 	return
	// }
	if err := config.DB.Where("order_id=?", order_id).Delete(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	if err := config.DB.Where("order_id=?", order_id).Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_id" + order_id: "is deleted"})
}
