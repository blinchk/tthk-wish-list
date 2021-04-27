package main

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/bredbrains/tthk-wish-list/endpoints/wishes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	authAPI := router.Group("/auth")
	authAPI.POST("/login", auth.Login)
	authAPI.POST("/register", auth.Register)
	wishAPI := router.Group("/wishes")
	wishAPI.GET("/suggestion", wishes.Suggestion)
	wishAPI.GET("/recieve", wishes.Receive)
	wishAPI.POST("/hide", wishes.Hide)
	wishAPI.POST("/add", wishes.Add)
	wishAPI.POST("/delete", wishes.Delete)
	wishAPI.POST("/edit", wishes.Edit)
	database.Connect()
	// Use in production build: autotls.Run(r, "wish-api.bredbrains.tech")
	router.Run()
}
