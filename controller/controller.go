package controller

import (
	"assigmentdua/config"
	"assigmentdua/model"

	"github.com/gin-gonic/gin"
)

// Handlers
func CreateOrder(c *gin.Context) {

	db := config.GetDB()

	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&order)
	c.JSON(200, order)
}

func GetAllData(c *gin.Context) {
	db := config.GetDB()
	var orders []model.Order
	db.Preload("Items").Find(&orders)
	c.JSON(200, orders)
}

func UpdateData(c *gin.Context) {
	db := config.GetDB()
	orderID := c.Param("id")
	var order model.Order

	if err := db.Preload("Items").First(&order, orderID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Save(&order)
	c.JSON(200, order)
}

func DeleteData(c *gin.Context) {
	db := config.GetDB()
	orderID := c.Param("id")
	var order model.Order

	if err := db.Preload("Items").First(&order, orderID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	db.Delete(&order)
	c.JSON(200, gin.H{"message": "Order deleted successfully"})
}
