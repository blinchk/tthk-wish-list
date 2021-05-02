package auth

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
	var err error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	accessToken := c.GetHeader("Token")
	err, user := database.UserData(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true, "user": user}
	c.JSON(http.StatusOK, message)
	return
}
