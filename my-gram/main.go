package main

import (
	"my-gram/config"
	"my-gram/controller"
	"my-gram/middleware"
	"my-gram/repository"
	"my-gram/service"

	_ "github.com/joho/godotenv/autoload"

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
	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	socialMediaRepository := repository.NewSocialMediaRepository(db)

	userService := service.NewUserService(validate, userRepository)
	photoService := service.NewPhotoService(validate, photoRepository, userRepository)
	commentService := service.NewCommentService(validate, commentRepository, photoRepository, userRepository)
	socialMediaService := service.NewSocialMediaService(validate, socialMediaRepository, userRepository)

	userController := controller.NewUserController(userService)
	photoController := controller.NewPhotoController(photoService)
	commentController := controller.NewCommentController(commentService)
	socialMediaController := controller.NewSocialMediaController(socialMediaService)

	r := gin.Default()

	// user
	users := r.Group("/users")
	users.POST("/register", userController.Register)
	users.POST("/login", userController.Login)
	users.PUT("/:id", middleware.IsAuthenticated, userController.UpdateUser)
	users.DELETE("/:id", middleware.IsAuthenticated, userController.DeleteUser)

	// photo
	photos := r.Group("/photos")
	photos.Use(middleware.IsAuthenticated)
	photos.POST("/", photoController.CreatePhoto)
	photos.GET("/", photoController.GetAllPhotos)
	photos.PUT("/:id", photoController.UpdatePhoto)
	photos.DELETE("/:id", photoController.DeletePhoto)

	// comment
	comments := r.Group("/comments")
	comments.Use(middleware.IsAuthenticated)
	comments.POST("/", commentController.CreateComment)
	comments.GET("/", commentController.GetAllComments)
	comments.PUT("/:id", commentController.UpdateComment)
	comments.DELETE("/:id", commentController.DeleteComment)

	// social medias
	socialMedias := r.Group("/socialmedias")
	socialMedias.Use(middleware.IsAuthenticated)
	socialMedias.POST("/", socialMediaController.CreateSocialMedia)
	socialMedias.GET("/", socialMediaController.GetAllSocialMedias)
	socialMedias.PUT("/:id", socialMediaController.UpdateSocialMedia)
	socialMedias.DELETE("/:id", socialMediaController.DeleteSocialMedia)

	r.Run(serverConfig.URI)
}
