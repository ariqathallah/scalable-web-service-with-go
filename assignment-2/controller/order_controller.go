package controller

import (
	"assignment-2/model"
	"assignment-2/repository"
	"assignment-2/util"
	"net/http"

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
	orderRepository repository.OrderRepository
}

func NewOrderController(orderRepository repository.OrderRepository) OrderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var order model.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	if err := c.orderRepository.CreateOrder(&order); err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusCreated, util.CreateResponse(true, order, ""))
}

func (c *orderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.orderRepository.GetAllOrders()
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, orders, ""))
}

func (c *orderController) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := c.orderRepository.GetOrderByID(id)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, order, ""))
}

func (c *orderController) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var order model.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	order, err = c.orderRepository.UpdateOrder(id, &order)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, order, ""))
}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.orderRepository.DeleteOrder(id)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, ""))
}
