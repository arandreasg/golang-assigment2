package main

import (
	"assigmentdua/config"
	"assigmentdua/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	// Setup Echo
	g := gin.Default()

	// Endpoints
	g.POST("/createOrder", controller.CreateOrder)
	g.GET("/getAll", controller.GetAllData)
	g.PUT("/updateData/:id", controller.UpdateData)
	g.PATCH("/updateData/:id", controller.UpdateData)
	g.DELETE("/deleteData/:id", controller.DeleteData)

	// Run server
	g.Run(":8070")

}
