package controllers

import (
	"goweb/config"
	"goweb/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder godoc
// @Summary Create order data
// @Description Add by JSON Order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.Order true "Add Order"
// @success 200 {object} models.Order
// @Router /order [post]
func CreateOrder(c *gin.Context) {
	var order models.Order
	c.ShouldBindJSON(&order)
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create data!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// GetOrders godoc
// @Summary Show a list of orders including the items
// @Tags orders
// @Accept  json
// @Produce  json
// @success 200 {object} models.Order
// @Router /orders [get]
func GetAllOrder(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrdersDetailById godoc
// @Summary Show order detail including items based on Order Id
// @Tags orders
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param order_id path int true "Order ID"
// @success 200 {object} models.Order
// @Router /order/{order_id} [get]
func GetOrderById(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	if err := config.DB.Preload(("Items")).Where("order_id=?", order_id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// UpdateOrderById godoc
// @Summary Update order or items by OrderId
// @Description Update by JSON Order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path int true "Order Id"
// @Param order body models.Order true "Update Order"
// @success 200 {object} models.Order
// @Router /order/{order_id} [put]
func UpdateOrder(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	if err := config.DB.Preload(("Items")).First(&order, order_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	c.ShouldBindJSON(&order)
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Cannot Updated!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// DeleteOrderById godoc
// @Summary Delete order including items by order id
// @Tags order
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param order_id path int true "Order ID"
// @success 200 {object} models.Order
// @Router /order/{order_id} [delete]
func DeleteOrder(c *gin.Context) {
	order_id := c.Params.ByName("order_id")
	var order models.Order
	var item models.Item
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
