package router

import (
	"fmt"
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
				return
			}

			if token != nil && token.Valid {
				endpoint(c)
				return
			}

		} else {
			message := gin.H{"success": false, "error": "You are not authorized."}
			c.JSON(http.StatusUnauthorized, message)
			return
		}

	})
}

func AddAuthEndpoints(router *gin.Engine) {
	authAPI := router.Group("/auth")
	authAPI.POST("/login", auth.Login)
	authAPI.POST("/register", auth.Register)
}
