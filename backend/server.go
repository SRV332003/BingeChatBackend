package main

import (
	"HangAroundBackend/controllers"
	"HangAroundBackend/controllers/socket"
	"HangAroundBackend/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Strive Backend API
// @version 1.0
// @description This is the documentation for backend APIs of Strive. You can explore the API endpoints here.

func initRouter() *gin.Engine {
	app := gin.Default()

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := app.Group("/api/v1")
	{
		router.POST("/user", controllers.RegisterUser)
		router.POST("/user/token", controllers.CreateToken)
		router.GET("/user/token", controllers.ReCreateToken)
		router.GET("/googlelogin", controllers.GetGoogleLoginUri)
		router.POST("/googlelogin", controllers.VerifyAuthCode)
	}
	router.Use(middlewares.AuthMiddlware)
	{
		router.PUT("/user", controllers.UpdateUser)
		router.HEAD("/googlelogin", controllers.VerifyToken)
		router.GET("/user", controllers.GetUser)
		router.GET("/chat", socket.SocketController)
	}

	return app
}
