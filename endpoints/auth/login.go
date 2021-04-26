package auth

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var err error
	c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err = database.VerifyUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
}
