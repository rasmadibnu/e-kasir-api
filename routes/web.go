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
	supplierRepo := repository.NewSupplierRepository(db)
	produkRepo := repository.NewProdukRepository(db)
	stokRepo := repository.NewStokRepository(db)
	transaksiRepo := repository.NewTransaksiRepository(db)
	kategoriRepo := repository.NewKategoriRepository(db)
	cartRepo := repository.NewCartRepository(db)

	// Service Asset
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	supplierService := service.NewSupplierService(supplierRepo)
	produkService := service.NewProdukService(produkRepo)
	stokService := service.NewStokService(stokRepo)
	transaksiService := service.NewTransaksiService(transaksiRepo, stokRepo, cartRepo)
	kategoriService := service.NewKategoriService(kategoriRepo)
	cartService := service.NewCartService(cartRepo, stokRepo, produkRepo)

	//Controller Asset
	authController := controller.NewAuthController(userService, authService)
	userController := controller.NewUserConstroller(userService)
	supplierController := controller.NewSupplierController(supplierService)
	produkController := controller.NewProdukController(produkService)
	stokController := controller.NewStokController(stokService)
	transaksiController := controller.NewTransaksiController(transaksiService)
	kategoriController := controller.NewKategoriController(kategoriService)
	cartsController := controller.NewCartController(cartService)

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
	v1.Static("/uploads", "./uploads")

	v1.GET("/users", userController.Index)
	v1.POST("/users", userController.Store)
	v1.GET("/users/:id", userController.Show)
	v1.PUT("/users/:id", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)

	v1.GET("/suppliers", supplierController.Index)
	v1.POST("/suppliers", supplierController.Store)
	v1.GET("/suppliers/:id", supplierController.Show)
	v1.PUT("/suppliers/:id", supplierController.Update)
	v1.DELETE("/suppliers/:id", supplierController.Delete)

	v1.GET("/produk", produkController.Index)
	v1.POST("/produk", produkController.Store)
	v1.GET("/produk/:id", produkController.Show)
	v1.PUT("/produk/:id", produkController.Update)
	v1.DELETE("/produk/:id", produkController.Delete)

	v1.GET("/kategori", kategoriController.Index)
	v1.POST("/kategori", kategoriController.Store)
	v1.GET("/kategori/:id", kategoriController.Show)
	v1.PUT("/kategori/:id", kategoriController.Update)
	v1.DELETE("/kategori/:id", kategoriController.Delete)

	v1.GET("/stoks", stokController.Index)
	v1.POST("/stoks", stokController.Store)
	v1.GET("/stoks/:id", stokController.Show)
	v1.DELETE("/stoks/:id", stokController.Delete)

	v1.GET("/transaksi", transaksiController.Index)
	v1.POST("/transaksi", transaksiController.Store)
	v1.GET("/transaksi/:id", transaksiController.Show)
	v1.DELETE("/transaksi/:id", transaksiController.Delete)

	v1.GET("/carts", cartsController.Index)
	v1.POST("/carts/:type", cartsController.Store)
	v1.GET("/carts/:id", cartsController.Show)
	v1.PUT("/carts/:id", cartsController.Update)
	v1.DELETE("/carts/:id", cartsController.Delete)

	httpRouter.Run(":" + os.Getenv("APP_PORT")) // Run Routes with PORT
}
