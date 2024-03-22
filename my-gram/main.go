package main

import (
	"my-gram/config"
	"my-gram/controller"
	"my-gram/repository"
	"my-gram/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	serverConfig := config.NewServerConfig()
	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(validate, userRepository)
	userController := controller.NewUserController(userService)

	r := gin.Default()

	// Users
	users := r.Group("/users")
	users.GET("/ping", userController.Ping)
	r.Run(serverConfig.URI)
}
