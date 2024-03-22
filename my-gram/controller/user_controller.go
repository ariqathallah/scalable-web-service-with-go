package controller

import (
	"my-gram/model/web"
	"my-gram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	})
}
