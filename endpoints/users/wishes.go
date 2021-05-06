package users

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Wishes(c *gin.Context) {
	var wishes []models.Wish
	var user models.User
	var err error
	user.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, wishes = database.GetWishes(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true, "wishes": wishes}
	c.JSON(http.StatusOK, message)
	return
}
