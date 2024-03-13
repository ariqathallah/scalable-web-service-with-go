package main

import (
	"assignment-2/config"
	"assignment-2/controller"
	"assignment-2/repository"
	"assignment-2/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	itemRepository := repository.NewItemRepository(db)
	orderService := service.NewOrderService(orderRepository, itemRepository)
	orderController := controller.NewOrderController(orderService)

	ginEngine := gin.Default()
	ginEngine.GET("/order/:id", orderController.GetOrderByID)
	ginEngine.GET("/order", orderController.GetAllOrders)
	ginEngine.POST("/order", orderController.CreateOrder)
	ginEngine.PUT("/order", orderController.UpdateOrder)
	ginEngine.DELETE("/order/:id", orderController.DeleteOrder)

	err = ginEngine.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
