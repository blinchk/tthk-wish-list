package main

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/bredbrains/tthk-wish-list/endpoints/wishes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	authAPI := router.Group("/auth")
	authAPI.POST("/register", auth.Register)
	authAPI.POST("/login", auth.Login)
	wishAPI := router.Group("/wishes")
	wishAPI.GET("/suggestion", wishes.Suggestion)
	wishAPI.GET("/recieve", wishes.Receive)
	wishAPI.POST("/hide", wishes.Hide)
	wishAPI.POST("/add", wishes.Add)
	wishAPI.POST("/delete", wishes.Delete)
	wishAPI.POST("/edit", wishes.Edit)
	database.Connect()
	// Use in production build
	// autotls.Run(r, "wish-api.bredbrains.tech")
	router.Run()
}
