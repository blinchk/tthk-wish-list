package auth

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var verified bool
	var accessToken string
	var err error
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	err, verified, accessToken = database.VerifyUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	if verified == true {
		message := gin.H{"success": true, "accessToken": accessToken}
		c.JSON(http.StatusOK, message)
	} else {
		message := gin.H{"success": false}
		c.JSON(http.StatusUnauthorized, message)
	}
}
