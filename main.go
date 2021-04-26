package main

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	authAPI := router.Group("/auth")
	authAPI.POST("/register", auth.Register)
	authAPI.POST("/login", auth.Login)
	database.Connect()
	// Use in production build: autotls.Run(r, "wish-api.bredbrains.tech")
	router.Run()
}
