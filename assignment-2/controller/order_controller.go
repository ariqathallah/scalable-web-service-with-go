package controller

import (
	"assignment-2/service"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	GetAllOrders(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (c *orderController) CreateOrder(ctx *gin.Context)

func (c *orderController) GetAllOrders(ctx *gin.Context)

func (c *orderController) GetOrderByID(ctx *gin.Context)

func (c *orderController) UpdateOrder(ctx *gin.Context)

func (c *orderController) DeleteOrder(ctx *gin.Context)
