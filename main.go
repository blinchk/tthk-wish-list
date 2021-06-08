package main

import (
	"github.com/bredbrains/tthk-wish-list/router"
	"log"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.AddAuthEndpoints(engine)
	router.AddUserEndpoints(engine)
	router.AddWishEndpoints(engine)

	database.Connect()
	log.Fatal(engine.Run())
}
