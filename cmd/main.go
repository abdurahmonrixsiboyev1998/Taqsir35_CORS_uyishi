package main

import (
	"cors/config"
	"cors/handlers"
	"cors/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitRedis()

	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/origins", handlers.GetOrigins)
		authorized.POST("/origins", handlers.AddOrigin)
		authorized.DELETE("/origins/:origin", handlers.DeleteOrigin)
	}

	r.Run(":8080")
}
