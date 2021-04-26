package auth

import (
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	message := gin.H{"username": user.Username, "password": user.Password}
	c.JSON(http.StatusOK, message)
}

func CreateToken(id uint64) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = id
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
