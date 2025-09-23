package main

import (
	"productApp/cmd/injection"
	"productApp/internal/auth"
	"productApp/internal/config"
	defotanim "productApp/internal/defo_tanim"
	order "productApp/internal/order"
	"productApp/internal/uretim"
	"productApp/pkg/logger"

	_ "productApp/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Production App api
// @version 1.0
// @description Production App Api documentation
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// gin settings
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// initialize logger

	logger.InitLogger()

	// initialize config

	config.InitConfig()

	// initialize db

	db := config.InitDB()

	// api
	api := r.Group("/api")

	// initialize  controller

	authController := injection.InitializeAuthController(&db)
	orderListController := injection.InitializeOrderController(&db)
	uretimController := injection.InitializeUretimController(&db)
	defoTanimController := injection.InitializeDefoTanimController(&db)

	// initalize  routes

	auth.RegisterAuthRoutes(api, authController)
	order.RegisterOrderRoutes(api, orderListController)
	uretim.RegisterUretimRoutes(api, uretimController)
	defotanim.RegisterDefoTanimRoutes(api, defoTanimController)

	// r.GET("/uploads/*filepath", func(c *gin.Context) {
	// 	filepath := c.Param("filepath")
	// 	fullPath := "./uploads" + filepath
	// 	c.File(fullPath)
	// })
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run server
	logger.Logger.Info("server start :3000")

	panic(r.Run(":3000"))

}
