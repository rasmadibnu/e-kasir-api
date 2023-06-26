package routes

import (
	"kasir-cepat-api/app/controller"
	"kasir-cepat-api/app/repository"
	"kasir-cepat-api/app/service"
	"kasir-cepat-api/config"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebRouter(db config.Database) {
	// Repository Asset
	userRepo := repository.NewUserRepository(db)

	// Service Asset
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	//Controller Asset
	authController := controller.NewAuthController(userService, authService)
	userController := controller.NewUserConstroller(userService)

	// Route
	httpRouter := gin.Default()

	// Register routing
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Testing  connection
	httpRouter.GET("/status-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Service ✅ API Up and Running"})
	})

	httpRouter.POST("/api/v1/login", authController.Login)

	v1 := httpRouter.Group("/api/v1") // Grouping routes
	v1.GET("/users", userController.Index)
	v1.POST("/users", userController.Store)
	v1.GET("/users/:id", userController.Show)
	v1.DELETE("/users/:id", userController.Delete)

	httpRouter.Run(":" + os.Getenv("APP_PORT")) // Run Routes with PORT
}
