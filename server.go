package main

import (
	"HangAroundBackend/controllers"
	"HangAroundBackend/controllers/socket"
	"HangAroundBackend/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title HangAround Backend API
// @version 1.0
// @description This is the documentation for backend APIs of Strive. You can explore the API endpoints here.
func initRouter() *gin.Engine {
	app := gin.Default()

	app.GET("/metrics", gin.WrapH(promhttp.Handler()))

	app.Use(middlewares.CORSMiddleware())
	router := app.Group("/api/v1")
	{
		router.POST("/user", controllers.RegisterUser)
		router.POST("/user/token", controllers.CreateToken)
		router.GET("/user/google", controllers.GetGoogleLoginUri)
		router.POST("/user/google", controllers.VerifyAuthCode)
		router.GET("/college", controllers.GetCollege)
		router.POST("/user/verify", controllers.VerifyUser)
		router.GET("/user/token", controllers.ReCreateToken)
		router.POST("/college", controllers.AddCollege)
	}
	router.Use(middlewares.AuthMiddlware)
	{
		router.PUT("/user", controllers.UpdateUser)
		router.GET("/user", controllers.GetUser)
		router.HEAD("/user/token", controllers.VerifyToken)
		router.GET("/chat", socket.SocketController)
		router.GET("/user/verify", controllers.IsUserVerified)
		router.POST("/report", controllers.CreateReport)
	}

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
