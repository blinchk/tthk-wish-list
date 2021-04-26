package main

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*s"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	authAPI := router.Group("/auth")
	authAPI.POST("/register", auth.Register)
	database.Connect()
	// Use in production build
	// autotls.Run(r, "wish-api.bredbrains.tech")
	router.Run()
}
