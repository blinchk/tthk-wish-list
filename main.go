package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bredbrains/tthk-wish-list/endpoints/follows"
	"github.com/bredbrains/tthk-wish-list/endpoints/users"

	"time"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/bredbrains/tthk-wish-list/endpoints/wishes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func isAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if c.Request.Header["Token"] != nil {
			token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Valid token": false})
			}

			if token != nil && token.Valid {
				endpoint(c)
			}

		} else {
			message := gin.H{"Authorization": "not authorized"}
			c.JSON(http.StatusUnauthorized, message)
		}

	})
}

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
	authAPI.POST("/login", auth.Login)
	authAPI.POST("/register", auth.Register)
	wishAPI := router.Group("/wish")
	wishAPI.GET("/suggestion", isAuthorized(wishes.Suggestion))
	wishAPI.PATCH("/:id/hide", isAuthorized(wishes.Hide))
	wishAPI.POST("/add", isAuthorized(wishes.Add))
	wishAPI.DELETE("/:id", isAuthorized(wishes.Delete))
	wishAPI.PATCH("/:id", isAuthorized(wishes.Update))
	userAPI := router.Group("/user")
	userAPI.GET("/", isAuthorized(auth.User))
	userAPI.POST("/", isAuthorized(users.EditUserProfile))
	userAPI.GET("/:id", users.GetUserProfile)
	userAPI.GET("/:id/wishes", isAuthorized(users.Wishes))
	userAPI.POST("/:id/follow", isAuthorized(follows.ToggleFollowing))
	database.Connect()
	// Use in production build
	// autotls.Run(r, "wish-api.bredbrains.tech")
	router.Run()
}
