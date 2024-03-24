package middleware

import (
	"my-gram/helper"
	"my-gram/model/web"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, web.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
		ctx.Abort()
		return
	}

	token := strings.Split(authorization, " ")[1]
	claims, err := helper.ParseJWT(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, web.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
		ctx.Abort()
		return
	}

	ctx.Set("claims", claims)
	ctx.Next()
}
