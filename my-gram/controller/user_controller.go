package controller

import (
	"my-gram/model/web"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var request web.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response, err := c.userService.Register(request)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Code:    http.StatusCreated,
		Message: "ok",
		Data:    response,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var request web.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response, err := c.userService.Login(request)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    response,
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	// get params
	param := ctx.Param("id")
	intParam, _ := strconv.Atoi(param)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	var request web.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response, err := c.userService.UpdateUser(intParam, intUSerID, request)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    response,
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	// get params
	param := ctx.Param("id")
	intParam, _ := strconv.Atoi(param)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	err := c.userService.DeleteUser(intParam, intUSerID)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Your account has been successfully deleted",
		Data:    nil,
	})
}
