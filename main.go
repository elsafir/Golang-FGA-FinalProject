package main

import (
	"Golang-FGA-FinalProject/controllers"
	"Golang-FGA-FinalProject/database"
	"Golang-FGA-FinalProject/middleware"
	"Golang-FGA-FinalProject/repositories"
	"Golang-FGA-FinalProject/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db, err := database.ConnectDB(database.GoDotEnvVariable("DB_DRIVER"))
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	photoRepo := repositories.NewPhotoRepo(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService, photoService)

	socialMediaRepo := repositories.NewSocialMediaRepo(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	route := gin.Default()
	// User
	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)

		userRoute.Use(middleware.Auth())
		userRoute.PUT("/:userId", userController.UpdateUser)
		userRoute.DELETE("/:userId", userController.DeleteUser)
	}

	// Photo
	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(middleware.Auth())
		photoRoute.POST("", photoController.CreatePhoto)
		photoRoute.GET("", photoController.GetPhotos)
		photoRoute.PUT("/:photoId", photoController.UpdatePhoto)
		photoRoute.DELETE("/:photoId", photoController.DeletePhoto)
	}

	// Comment
	commentRoute := route.Group("/comments")
	{
		commentRoute.Use(middleware.Auth())
		commentRoute.POST("", commentController.CreateComment)
		commentRoute.GET("", commentController.GetComments)
		commentRoute.PUT("/:commentId", commentController.UpdateComment)
		commentRoute.DELETE("/:commentId", commentController.DeleteComment)
	}

	// Social Media
	socialMediaRoute := route.Group("/socialmedias")
	{
		socialMediaRoute.Use(middleware.Auth())
		socialMediaRoute.POST("", socialMediaController.CreateSocialMedia)
		socialMediaRoute.GET("", socialMediaController.GetSocialMedias)
		socialMediaRoute.PUT("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaRoute.DELETE("/:socialMediaId", socialMediaController.DeleteSocialMedia)
	}

	route.Run(database.GoDotEnvVariable("APP_PORT"))
}
