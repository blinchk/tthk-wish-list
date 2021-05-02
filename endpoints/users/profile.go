package users

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserProfile(c *gin.Context) {
	var wishes []models.Wish
	var user models.User
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, user = database.UserDataById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Requested user isn't exists."})
		return
	}
	err, wishes = database.GetWishes(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "wishes": wishes})
	return
}
