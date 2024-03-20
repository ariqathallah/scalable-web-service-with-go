package main

import (
	"assignment-3/controller"
	"assignment-3/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")

	go helper.UpdateData()

	r.GET("/", controller.GetData)
	r.Run("localhost:8080")
}
