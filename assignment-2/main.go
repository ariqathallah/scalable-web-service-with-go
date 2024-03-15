package main

import (
	"assignment-2/config"
	"assignment-2/controller"
	"assignment-2/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	ginEngine := gin.Default()
	ginEngine.GET("/orders/:id", orderController.GetOrderByID)
	ginEngine.GET("/orders", orderController.GetAllOrders)
	ginEngine.POST("/orders", orderController.CreateOrder)
	ginEngine.PUT("/orders/:id", orderController.UpdateOrder)
	ginEngine.DELETE("/orders/:id", orderController.DeleteOrder)

	err = ginEngine.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
