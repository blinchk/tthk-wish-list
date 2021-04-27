package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	var err error
	c.BindJSON(&user)
	user.AccessToken, err = CreateToken(user.Username)
	user.RegistrationTime = time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err = database.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	message := gin.H{"token": user.AccessToken}
	c.JSON(http.StatusOK, message)
}
func CreateToken(username string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = username
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
