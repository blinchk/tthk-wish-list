package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// Use in production build: autotls.Run(r, "wish-api.bredbrains.tech")
	r.Run()
}
