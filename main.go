package main

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()
	// Use in production build: autotls.Run(r, "wish-api.bredbrains.tech")
	r.Run()
}
